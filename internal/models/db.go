package models

import (
	"oe/conf"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Namespace puts all tables names under a common
// namespace
var Namespace string

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

// Connect will connect to that storage engine
func Connect(config *conf.GlobalConfiguration, log logrus.FieldLogger) (*ORM, error) {
	if config.DB.Namespace != "" {
		Namespace = config.DB.Namespace
	}

	if config.DB.Dialect == "" {
		config.DB.Dialect = config.DB.Driver
	}
	db, err := gorm.Open(config.DB.Dialect, config.DB.Driver, config.DB.URL)
	if err != nil {
		return nil, errors.Wrap(err, "opening database connection")
	}
	// db.SetLogger(NewDBLogger(log))
	db.LogMode(true)

	err = db.DB().Ping()
	if err != nil {
		return nil, errors.Wrap(err, "checking database connection")
	}

	if config.DB.Automigrate {
		migDB := db.New()
		// migDB.SetLogger(NewDBLogger(log.WithField("task", "migration")))
		if err := AutoMigrate(migDB); err != nil {

			return nil, errors.Wrap(err, "migrating tables")
		}
	}

	db.Set("gorm:auto_preload", true)

	orm := &ORM{DB: db}

	return orm, nil
}

func tableName(defaultName string) string {
	if Namespace != "" {
		return Namespace + "_" + defaultName
	}
	return defaultName
}

// AutoMigrate runs the gorm automigration for all models
func AutoMigrate(db *gorm.DB) error {
	db = db.AutoMigrate(
		Product{},
		ProductVariant{},
		Seller{},
		Bank{},
		Role{},
		Permission{},
		LoginHistory{},
		RolePermission{},
		// SellerProduct{},
		User{},
		Address{},
		Category{},
		Image{},
	)
	return db.Error
}

// FindUserByJWT finds the user that is related to the APIKey token
func (o *ORM) FindUserByJWT(userID string) (*User, error) {
	db := o.DB.New()
	up := &User{}

	// db = db.Preload("Roles")
	db = db.Preload("Roles.Permissions")

	if err := db.Preload("Roles").Where("id = ?", userID).First(up).Error; err != nil {
		return nil, err
	}
	logrus.Info(up.Roles, up.Permissions, "------>")
	return up, nil
}

func (o *ORM) FindUserByEmail(email string) (*User, error) {
	db := o.DB.New()
	up := &User{}

	db = db.Preload("Roles.Permissions")
	if err := db.Preload("Roles").Where("email = ?", email).First(up).Error; err != nil {
		return nil, err
	}
	return up, nil
}

func (o *ORM) FindSellerUser(user *User) (*Seller, error) {
	db := o.DB.New()

	seller := Seller{}

	db = db.Where("user_id = ?", user.ID).First(&seller)

	return &seller, db.Error
}
