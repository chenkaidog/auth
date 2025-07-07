package dto

type ResourceStatus int

const (
	ResourceStatusValid   ResourceStatus = 1
	ResourceStatusInvalid ResourceStatus = 2
)

type Resource struct {
	ResourceID  string         `json:"resource_id"`
	DomainID    string         `json:"domain_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Status      ResourceStatus `json:"status"`
}

type ResourceCreateReq struct {
	DomainID    string         `query:"domain_id" json:"-" binding:"required"`
	Name        string         `json:"name" binding:"required"`
	Status      ResourceStatus `json:"status" binding:"required"`
	Description string         `json:"description"`
}

type ResourceCreateResp struct {
}

type ResourceUpdateReq struct {
	ResourceID  string         `query:"resource_id" json:"-" binding:"required"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Status      ResourceStatus `json:"status"`
}

type ResourceUpdateResp struct {
}

type ResourceDeleteReq struct {
	ResourceID string `query:"resource_id" json:"-" binding:"required"`
}

type ResourceDeleteResp struct{}

type ResourceDetailQueryReq struct {
	ResourceID string `query:"resource_id" json:"-" binding:"required"`
}
type ResourceDetailQueryResp struct {
	Resource
}

type ResourceListQueryReq struct {
	DomainID string         `query:"domain_id" binding:"required"`
	Page     int            `query:"page" binding:"required,gt=0"`
	Size     int            `query:"size" binding:"required,gt=0"`
	Name     string         `query:"name"`
	Status   ResourceStatus `query:"status"`
}

type ResourceListQueryResp struct {
	Page      int         `json:"page"`
	Size      int         `json:"size"`
	Total     int         `json:"total"`
	Resources []*Resource `json:"resources"`
}
