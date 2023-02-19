package response

import "github.com/yin-zt/go-bind-mp/pkg/model"

type DomainListRsp struct {
	Total   int64          `json:"total"`
	Domains []model.Domain `json:"domains"`
}
