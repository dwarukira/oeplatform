package tf

import (
	"errors"
	"strconv"

	gql "oe/internal/gql/models"
	dbm "oe/internal/models"

	"github.com/gofrs/uuid"
)

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gql.UserInput, update bool, ids ...string) (o *dbm.User, err error) {
	o = &dbm.User{
		Name:        i.Name,
		FirstName:   i.FirstName,
		LastName:    i.LastName,
		NickName:    i.NickName,
		Description: i.Description,
		Location:    i.Location,
		UserStatus:  "active",
	}
	if i.Email == nil && !update {
		return nil, errors.New("field [email] is required")
	}
	if i.Password == nil && !update {
		return nil, errors.New("field [password] is required")
	}
	if i.Email != nil {
		o.Email = *i.Email
	}
	if i.Password != nil {
		o.Password = *i.Password
	}
	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}
		o.ID = updID.String()
	}
	return o, err
}

// GQLRoleToDBRole transforms [role] db input to gql type
func GQLRoleToDBRole(i *gql.RoleInput, update bool, ids ...string) (o *dbm.Role, err error) {
	o = &dbm.Role{
		Description: *i.Description,
		Name:        i.Name,
	}
	return o, err
}

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	// status, _ := DBUserStatusGQLUserStatus(&i.UserStatus)

	roles := []*gql.Role{}
	for _, vrec := range i.Roles {
		if rec, err := DBRoleToGQLRole(&vrec); err != nil {

		} else {
			roles = append(roles, rec)
		}
	}

	o = &gql.User{
		AvatarURL:   i.AvatarURL,
		ID:          i.ID,
		Email:       i.Email,
		Name:        i.Name,
		FirstName:   i.FirstName,
		LastName:    i.LastName,
		NickName:    i.NickName,
		Description: i.Description,
		Location:    i.Location,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		LastLogin:   i.LastLogin,
		Roles:       roles,
	}
	return o, err
}

func DBPermissionToGQLPermission(i *dbm.Permission) (o *gql.Permission, err error) {
	o = &gql.Permission{
		ID:          strconv.Itoa(i.ID),
		Tag:         i.Tag,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		Description: &i.Description,
	}

	return o, err
}

// DBRoleToGQLRole transforms [role] db input to gql type
func DBRoleToGQLRole(i *dbm.Role) (o *gql.Role, err error) {
	permissions := []*gql.Permission{}
	for _, vrec := range i.Permissions {
		if rec, err := DBPermissionToGQLPermission(&vrec); err != nil {

		} else {
			permissions = append(permissions, rec)
		}
	}

	o = &gql.Role{
		ID:          strconv.Itoa(i.ID),
		Name:        i.Name,
		Description: &i.Description,
		Permissions: permissions,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}
	return o, err
}

func DBUserStatusGQLUserStatus(i *dbm.UserStatus) (o *gql.UserStatus, err error) {
	i.Scan(o)
	return o, err
}

// DBUserToGQLAuthUser transforms [user] db input to gql type
func DBUserToGQLAuthUser(i *dbm.User, token *string) (o *gql.Authuser, err error) {

	o = &gql.Authuser{
		User: &gql.User{
			AvatarURL:   i.AvatarURL,
			ID:          i.ID,
			Email:       i.Email,
			Name:        i.Name,
			FirstName:   i.FirstName,
			LastName:    i.LastName,
			NickName:    i.NickName,
			Description: i.Description,
			Location:    i.Location,
			CreatedAt:   i.CreatedAt,
			UpdatedAt:   i.UpdatedAt,
		},
		Token: token,
	}

	return o, err
}
