package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// BaseModel defines the common columns that all db structs should hold, usually
// db structs based on this have no soft delete
type BaseModel struct {
	ID        string     `gorm:"PRIMARY_KEY;TYPE:uuid;"`
	CreatedAt time.Time  `gorm:"INDEX;NOT NULL;DEFAULT:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"INDEX"`
}

// BeforeCreate set Model's primary key value to uuid
// http://doc.gorm.io/crud.html#setting-primary-key-in-callbacks
func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()

	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err
	}

	scope.SetColumn("ID", id.String())
	return nil
}

// BaseModelSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `gorm:"INDEX"`
}

// BaseModelSeq defines the common columns that all db structs should hold, with an INT key
type BaseModelSeq struct {
	ID        int        `gorm:"PRIMARY_KEY,AUTO_INCREMENT;"`
	CreatedAt time.Time  `gorm:"INDEX;NOT NULL;DEFAULT:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"INDEX"`
}

// BaseModelSeqSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSeqSoftDelete struct {
	BaseModelSeq
	DeletedAt *time.Time `gorm:"INDEX"`
}
