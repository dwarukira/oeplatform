package resolvers

import (
	"context"
	"fmt"
	"oe/internal/gql/models"
	"oe/internal/gql/tf"
	"oe/pkg/mail"

	dbm "oe/internal/models"

	"github.com/pkg/errors"
)

func (r *mutationResolver) CreateSeller(ctx context.Context, input models.SellerInput) (*models.Seller, error) {
	user := getCurrentUser(ctx)

	if user == nil {
		return &models.Seller{}, fmt.Errorf("Access denied")
	}

	// user := middleware.ForContext(ctx)

	return sellerCreateUpdate(r, input, user, false)
}

func (r *queryResolver) Sellers(ctx context.Context, id *string) (*models.Sellers, error) {
	return sellerList(r, id)
}

func sellerCreateUpdate(r *mutationResolver, input models.SellerInput, user *dbm.User, update bool, ids ...string) (*models.Seller, error) {
	dbo, err := tf.GQLInputSellerToDBSeller(&input, user, update)

	if err != nil {
		return nil, err
	}

	db := r.ORM.DB.New().Begin()

	seller, _ := r.ORM.FindSellerUser(user)
	// &seller.ID != nil {
	// 	return nil, errors.New("Can't create a seller for this user. One exist")
	// }

	if len(seller.ID) > 0 {
		return nil, errors.New("Can't create a seller for this user. One exist")
	}
	role := dbm.Role{}

	db = db.Model(role).Where("name = ?", "Seller").First(&role)
	dbo.User.Roles = append(dbo.User.Roles, role)
	db = db.Create(dbo).First(dbo)

	gql, err := tf.DBSellerToGQLSeller(dbo)

	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}
	db = db.Commit()
	// TODO update the password
	if !update {
		mail.SendSimpleMessage("https://api.mailgun.net/v3/sandbox561903decfef4126bf7dce12f4e91d98.mailgun.org", "842c142f15d727261444bc3634cfe13b-c27bf672-10c53ddc")
	}
	return gql, err
}

func sellerList(r *queryResolver, id *string) (*models.Sellers, error) {
	whereID := "id = ?"
	record := &models.Sellers{}
	dbRecords := []*dbm.Seller{}

	db := r.ORM.DB.New()

	db = db.Preload("User").Preload("Bank")
	if id != nil {
		db = db.Where(whereID, *id)
	}

	db = db.Find(&dbRecords).Count((&record.Count))

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBSellerToGQLSeller(dbRec); err != nil {

		} else {
			record.List = append(record.List, rec)
		}
	}

	return record, db.Error
}
