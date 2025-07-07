package dto

type RoleStatus int

const (
	RoleStatusValid   RoleStatus = 1
	RoleStatusInvalid RoleStatus = 2
)

type Role struct {
	RoleID      string     `json:"role_id"`
	DomainID    string     `json:"domain_id"`
	ParentID    string     `json:"parent_id"`
	Name        string     `json:"name"`
	Status      RoleStatus `json:"status"`
	Description string     `json:"description"`
}

type RoleCreateReq struct {
	DomainID    string     `query:"domain_id" json:"-" binding:"required"`
	ParentID    string     `json:"parent_id"`
	Name        string     `json:"name" binding:"required"`
	Status      RoleStatus `json:"status" binding:"required"`
	Description string     `json:"description"`
}

type RoleCreateResp struct {
}

type RoleUpdateReq struct {
	RoleID      string     `query:"role_id" json:"-" binding:"required"`
	Name        string     `json:"name"`
	Status      RoleStatus `json:"status"`
	Description string     `json:"description"`
}

type RoleUpdateResp struct {
}

type RoleDeleteReq struct {
	RoleID string `query:"role_id" json:"-" binding:"required"`
}

type RoleDeleteResp struct {
}

type RoleDetailQueryReq struct {
	RoleID string `query:"role_id" binding:"required"`
}

type RoleDetailQueryResp struct {
	Role
}

type RoleListQueryReq struct {
	Page     int        `query:"page" binding:"required,gt=0"`
	Size     int        `query:"size" binding:"required,gt=0"`
	DomainID string     `query:"domain_id" binding:"required"`
	Name     string     `query:"name"`
	Status   RoleStatus `query:"status"`
}

type RoleListQueryResp struct {
	Page  int     `json:"page"`
	Size  int     `json:"size"`
	Total int     `json:"total"`
	Roles []*Role `json:"roles"`
}
