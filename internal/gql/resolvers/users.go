package resolvers

import (
	"context"
	"oe/internal/gql/models"
	"oe/internal/gql/tf"
	"oe/internal/handlers/middleware"
	"time"

	ocontext "oe/pkg/context"
	"oe/pkg/logger"

	dbm "oe/internal/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.Authuser, error) {
	return userCreateUpdate(r, input, false)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UserInput) (*models.Authuser, error) {
	return userCreateUpdate(r, input, true, id)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateRole(ctx context.Context, input models.RoleInput) (*models.Role, error) {
	return roleCreateUpdate(r, input, false)

}

func (r *mutationResolver) TokenCreate(ctx context.Context, input models.TokenCreateInput) (*models.Authuser, error) {

	return createToken(r, input, ctx)
}

func (r *queryResolver) Users(ctx context.Context, id *string) (*models.Users, error) {
	return userList(r, id)
}

func (r *queryResolver) Roles(ctx context.Context, id *string) (*models.Roles, error) {
	return roleList(r, id)
}

func createToken(r *mutationResolver, input models.TokenCreateInput, ctx context.Context) (*models.Authuser, error) {
	user, err := r.ORM.FindUserByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if user.CheckPassword(input.Password) {

		seller, err := r.ORM.FindSellerUser(user)
		var s bool
		s = false
		if len(seller.ID) > 0 {
			s = true
		}
		token := middleware.CreateJWTToken(*user, s)
		t := time.Now()

		gc, err := middleware.GinContextFromContext(ctx)

		db := r.ORM.DB.New()

		login := &dbm.LoginHistory{
			User:      *user,
			UserAgent: gc.Request.UserAgent(),
			IP:        gc.ClientIP(),
		}
		db = r.ORM.DB.Create(login)
		user.LastLogin = &t
		db = r.ORM.DB.Save(user)

		gql, err := tf.DBUserToGQLAuthUser(user, &token)
		db.Commit()
		return gql, err
	}

	return nil, err
}

func userCreateUpdate(r *mutationResolver, input models.UserInput, update bool, ids ...string) (*models.Authuser, error) {
	dbo, err := tf.GQLInputUserToDBUser(&input, update, ids...)

	if err != nil {
		return nil, err
	}

	db := r.ORM.DB.New().Begin()

	if !update {
		role, _ := r.ORM.DefaultRole()
		dbo.Roles = []dbm.Role{*role}
		dbo.Password = middleware.GeneratePassword(*input.Password)
		db = db.Create(dbo).First(dbo) // Create the user
	} else {
		db = db.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}

	token := middleware.CreateJWTToken(*dbo, false)

	gql, err := tf.DBUserToGQLAuthUser(dbo, &token)
	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}
	db = db.Commit()

	return gql, db.Error
}

func roleCreateUpdate(r *mutationResolver, input models.RoleInput, update bool, ids ...string) (*models.Role, error) {
	dbo, err := tf.GQLRoleToDBRole(&input, update, ids...)

	if err != nil {
		return nil, err
	}

	db := r.ORM.DB.New().Begin()

	db = db.Create(dbo).First(dbo)

	gql, err := tf.DBRoleToGQLRole(dbo)
	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}
	db = db.Commit()

	return gql, db.Error
}
func roleList(r *queryResolver, id *string) (*models.Roles, error) {
	// entity := "users"
	whereID := "id = ?"
	record := &models.Roles{}
	dbRecords := []*dbm.Role{}
	db := r.ORM.DB.New()
	if id != nil {
		db = db.Where(whereID, *id)
	}
	db = db.Preload("Permissions").Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		if rec, err := tf.DBRoleToGQLRole(dbRec); err != nil {
			// logger.Errorfn(entity, err)
		} else {
			record.List = append(record.List, rec)
		}
	}
	return record, db.Error
}

func userList(r *queryResolver, id *string) (*models.Users, error) {
	// entity := "users"
	whereID := "id = ?"
	record := &models.Users{}
	dbRecords := []*dbm.User{}
	db := r.ORM.DB.New()
	if id != nil {
		db = db.Where(whereID, *id)
	}
	db = db.Preload("Roles")
	db = db.Preload("Roles.Permissions")
	db = db.Find(&dbRecords).Count(&record.Count)

	for _, dbRec := range dbRecords {
		logger.Info(dbRec.Roles, "->")

		if rec, err := tf.DBUserToGQLUser(dbRec); err != nil {
			// logger.Errorfn(entity, err)
		} else {
			record.List = append(record.List, rec)
		}
	}
	return record, db.Error
}

func getCurrentUser(ctx context.Context) *dbm.User {
	return ctx.Value(ocontext.ProjectContextKeys.UserCtxKey).(*dbm.User)
}
