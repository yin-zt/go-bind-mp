package response

import "github.com/yin-zt/go-bind-mp/pkg/model"

type ViewListRsp struct {
	Total int64        `json:"total"`
	Views []model.View `json:"views"`
}
