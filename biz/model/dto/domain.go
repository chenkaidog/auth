package dto

type DomainCreateReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type DomainCreateResp struct {
}

type DomainUpdateReq struct {
	DomainID    string `query:"domain_id" json:"-" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DomainUpdateResp struct {
}

type DomainDeleteReq struct {
	DomainID string `query:"domain_id" json:"-" binding:"required"`
}

type DomainDeleteResp struct {
}

type DomainDetailQueryReq struct {
	DomainID string `query:"domain_id" binding:"required"`
}

type Domain struct {
	DomainID    string `json:"domain_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DomainDetailQueryResp struct {
	Domain
}

type DomainListQueryReq struct {
	Page int `query:"page" binding:"required,gt=0"`
	Size int `query:"size" binding:"required,gt=0"`
}

type DomainListQueryResp struct {
	Page    int       `json:"page"`
	Size    int       `json:"size"`
	Total   int       `json:"total"`
	Domains []*Domain `json:"domains"`
}
