package models

import (
	"database/sql/driver"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserStatus string

const (
	Active    UserStatus = "active"
	InActive  UserStatus = "inactive"
	Suspended UserStatus = "suspended"
)

func (e *UserStatus) Scan(value interface{}) error {
	*e = UserStatus(value.([]byte))
	return nil
}

func (e UserStatus) Value() (driver.Value, error) {
	return string(e), nil
}

// User defines a user for the app
type User struct {
	BaseModelSoftDelete        // We don't to actually delete the users, maybe audit
	Email               string `gorm:"not null;unique_index:idx_email"`
	Password            string
	Name                *string
	NickName            *string
	FirstName           *string
	LastName            *string
	Location            *string
	UserStatus          UserStatus `json:"status" sql:"type:user_status"`
	AvatarURL           *string    `gorm:"size:1024"`
	Description         *string    `gorm:"size:1024"`
	LastLogin           *time.Time `gorm:"INDEX;NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	Profiles            []UserProfile
	Roles               []Role       `gorm:"many2many:user_roles;association_autoupdate:false;association_autocreate:false"`
	Permissions         []Permission `gorm:"many2many:user_permissions;association_autoupdate:false;association_autocreate:false"`
	Logins              []LoginHistory
	Addresses           []Address
}

func (user *User) GetPermissions() []Permission {
	var permissions []Permission

	for _, role := range user.Roles {
		for _, per := range role.Permissions {
			permissions = append(permissions, per)
		}
	}

	return permissions
}

// TableName returns the database table name for the users model.
// func (User) TableName() string {
// 	return tableName("users")
// }

type LoginHistory struct {
	BaseModelSeq
	IP        string
	User      User   `gorm:"association_autoupdate:false;association_autocreate:false"`
	UserID    string `gorm:"not null;index"`
	UserAgent string
}

// TableName returns the database table name for the users model.
// func (LoginHistory) TableName() string {
// 	return tableName("login_histories")
// }

// UserProfile saves all the related OAuth Profiles
type UserProfile struct {
	BaseModelSeq
	Email          string    `gorm:"unique_index:idx_email_provider_external_user_id"`
	User           User      `gorm:"association_autoupdate:false;association_autocreate:false"`
	UserID         uuid.UUID `gorm:"not null;index"`
	Provider       string    `gorm:"not null;index;unique_index:idx_email_provider_external_user_id;default:'DB'"` // DB means database or no ExternalUserID
	ExternalUserID string    `gorm:"not null;index;unique_index:idx_email_provider_external_user_id"`              // User ID
	Name           string
	NickName       string
	FirstName      string
	LastName       string
	Location       string  `gorm:"size:1024"`
	AvatarURL      string  `gorm:"size:1024"`
	Description    *string `gorm:"size:1024"`
}

func (user *User) CheckPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}

	return true
}

// UserProfile returns the database table name for the users model.
// func (UserProfile) TableName() string {
// 	return tableName("user_profiles")
// }

// UserAPIKey generated api keys for the users
type UserAPIKey struct {
	BaseModelSeq
	User        User      `gorm:"association_autoupdate:false;association_autocreate:false"`
	UserID      uuid.UUID `gorm:"not null;index"`
	APIKey      string    `gorm:"size:128;unique_index"`
	Name        string
	Permissions []Permission `gorm:"many2many:user_permissions;association_autoupdate:false;association_autocreate:false"`
}

// UserRole relation between an user and its roles
type UserRole struct {
	UserID uuid.UUID `gorm:"index"`
	RoleID int       `gorm:"index"`
}

// UserPermission relation between an user and its permissions
type UserPermission struct {
	UserID       uuid.UUID `gorm:"index"`
	PermissionID int       `gorm:"index"`
}

// Hooks

// User

// BeforeSave hook for User
// func (user *User) BeforeSave(scope *gorm.Scope) {
// 	if user.Password != "" {
// 		if pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 11); err == nil {
// 			scope.SetColumn("Password", pw)
// 		}
// 	}
// 	return
// }

// AfterSave hook (assigning roles, fill all permissions for example)

// UserAPIKey

// BeforeSave hook for UserAPIKey
// func (k *UserAPIKey) BeforeSave(scope *gorm.Scope) error {
// 	db := scope.NewDB()
// 	if k.Name == "" {
// 		u := &User{}
// 		if err := db.Where("id = ?", k.UserID).First(u).Error; err != nil {
// 			return err
// 		}
// 		// k.Name =
// 	}
// 	if hash, err := bcrypt.GenerateFromPassword([]byte(k.UserID.String()), 0); err == nil {
// 		hasher := sha1.New()
// 		hasher.Write(hash)
// 		scope.SetColumn("APIKey", hex.EncodeToString(hasher.Sum(nil)))
// 	}
// 	return nil
// }

func (o *ORM) GetUser(id string) (*User, error) {
	db := o.DB.New()
	up := &User{}
	db.Preloads("Logins")
	db.Preloads("Roles")
	db.Preloads("Roles.Permissions")
	db.Preloads("Address")
	if err := db.First(up, id).Error; err != nil {
		return nil, err
	}

	return up, nil
}
