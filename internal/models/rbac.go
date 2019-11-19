package models

import (
	"oe/pkg/logger"

	"github.com/jinzhu/gorm"
)

// Role defines a role for the user
type Role struct {
	BaseModelSeq
	Name        string       `gorm:"not null"`
	Description string       `gorm:"size:1024"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

// Permission defines a permission scope for the user
type Permission struct {
	BaseModelSeq
	Tag         string `gorm:"not null"`
	Description string `gorm:"size:1024"`
}

// RolePermission defines relation between a role and permissions
type RolePermission struct {
	RoleID       int
	PermissionID int
}

// AfterSave hook to
func (r *RolePermission) AfterSave(scope *gorm.Scope) (err error) {
	db := scope.DB().New()
	ur := []UserRole{}
	db.Model(&UserRole{}).Where("role_id = ?", r.RoleID).Find(&ur)
	p := &Permission{}
	db.Model(p).First(p, r.PermissionID)
	for i := 0; i < len(ur); i++ {
		u := &User{}
		if err := db.Where("id = ?", ur[i].UserID).First(u).Association("Permissions").Append(p).Error; err != nil {
			logger.Warnf("Could not save permission: %s for user: %s", p.Tag, u.ID)
		}
	}
	db.Where("role_id = ? AND permission_id = ?", r.RoleID, r.PermissionID).Delete(r)
	return nil
}

// AfterDelete hook to
func (r *RolePermission) AfterDelete(scope *gorm.Scope) (err error) {
	db := scope.DB().New()
	ur := []UserRole{}
	db.Model(&UserRole{}).Where("role_id = ?", r.RoleID).Find(&ur)
	p := &Permission{}
	db.Model(p).First(p, r.PermissionID)
	for i := 0; i < len(ur); i++ {
		u := &User{}
		if err := db.Where("id = ?", ur[i].UserID).First(u).Association("Permissions").Delete(p).Error; err != nil {
			logger.Warnf("Could not delete permission: %s for user: %s", p.Tag, u.ID)
		}
	}
	return nil
}

func (o *ORM) DefaultRole() (*Role, error) {
	db := o.DB.New()
	role := &Role{}

	if err := db.Where("name = ?", "User").First(role).Association("Permissions").Error; err != nil {
		// logger.Warnf("Could not save permission: %s for user: %s", p.Tag, u.ID)
	}

	return role, nil
}
