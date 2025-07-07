package dto

type PermissionStatus int

const (
	PermissionStatusValid   PermissionStatus = 1
	PermissionStatusInvalid PermissionStatus = 2
)

type PermissionEffect int

const (
	PermissionEffectAllow PermissionEffect = 1
	PermissionEffectDeny  PermissionEffect = 2
)

type Permission struct {
	PermissionID string           `json:"permission_id"`
	Role         Role             `json:"role"`
	Resource     Resource         `json:"resource"`
	Action       string           `json:"action"`
	Effect       PermissionEffect `json:"effect"`
	Status       PermissionStatus `json:"status"`
}

type PermissionCreateReq struct {
	RoleID     string           `query:"role_id" json:"-" binding:"required"`
	ResourceID string           `json:"resource_id" binding:"required"`
	Action     string           `json:"action" binding:"required"`
	Effect     PermissionEffect `json:"effect" binding:"required"`
	Status     PermissionStatus `json:"status" binding:"required"`
}

type PermissionCreateResp struct {
}

type PermissionUpdateReq struct {
	PermissionID string           `query:"permission_id" json:"-" binding:"required"`
	Effect       PermissionEffect `json:"effect"`
	Status       PermissionStatus `json:"status"`
}

type PermissionUpdateResp struct {
}

type PermissionDeleteReq struct {
	PermissionID string `query:"permission_id" json:"-" binding:"required"`
}

type PermissionDeleteResp struct {
}

type PermissionDetailQueryReq struct {
	PermissionID string `json:"permission_id" binding:"required"`
}

type PermissionDetailQueryResp struct {
	Permission
}

type PermissionListQueryReq struct {
	Page   int    `query:"page" binding:"required,gt=0"`
	Size   int    `query:"size" binding:"required,gt=0"`
	RoleID string `query:"role_id" binding:"required"`
}

type PermissionListQueryResp struct {
	Page       int           `json:"page"`
	Size       int           `json:"size"`
	Total      int64         `json:"total"`
	Permissons []*Permission `json:"permissions"`
}
