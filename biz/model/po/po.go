package po

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type GormModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt
}

type Account struct {
	GormModel
	AccountID string `gorm:"column:account_id"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	Salt      string `gorm:"column:salt"`
	Status    string `gorm:"column:status"`
}

func (Account) TableName() string {
	return "account"
}

type User struct {
	GormModel
	AccountID   string `gorm:"column:account_id"`
	UserID      string `gorm:"column:user_id"`
	Name        string `gorm:"column:name"`
	Gender      string `gorm:"column:gender"`
	Phone       string `gorm:"column:phone"`
	Email       string `gorm:"column:email"`
	Description string `gorm:"column:description"`
}

func (User) TableName() string {
	return "user"
}

type Domain struct {
	GormModel
	DomainID    string `gorm:"column:domain_id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	DeletedAt   soft_delete.DeletedAt
}

func (Domain) TableName() string {
	return "domain"
}

type Role struct {
	GormModel
	RoleID       string `gorm:"column:role_id"`
	DomainID     string `gorm:"column:domain_id"`
	ParentRoleID string `gorm:"column:parent_role_id"`
	Name         string `gorm:"column:name"`
	Status       string `gorm:"column:status"`
	Description  string `gorm:"column:description"`
}

func (Role) TableName() string {
	return "role"
}

type Resource struct {
	GormModel
	ResourceID  string `gorm:"column:resource_id"`
	DomainID    string `gorm:"column:domain_id"`
	Name        string `gorm:"column:name"`
	Status      string `gorm:"column:status"`
	Description string `gorm:"column:description"`
}

// TableName 返回 Resource 结构体对应的数据库表名
func (Resource) TableName() string {
	return "resource"
}

type UserRole struct {
	GormModel
	RelationID string    `gorm:"column:relation_id"`
	UserID     string    `gorm:"column:user_id"`
	RoleID     string    `gorm:"column:role_id"`
	Status     string    `gorm:"column:status"`
	ExpireAt   time.Time `gorm:"column:expire_at"`
}

func (UserRole) TableName() string {
	return "user_role"
}

type Permission struct {
	GormModel
	PermissionID string `gorm:"column:permission_id"`
	RoleID       string `gorm:"column:role_id"`
	ResourceID   string `gorm:"column:resource_id"`
	Action       string `gorm:"column:action"`
	Effect       string `gorm:"column:effect"`
	Status       string `gorm:"column:status"`
}

func (Permission) TableName() string {
	return "permission"
}
