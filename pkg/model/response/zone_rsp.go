package response

import "github.com/yin-zt/go-bind-mp/pkg/model"

type ZoneListRsp struct {
	Total int64        `json:"total"`
	Zones []model.Zone `json:"zones"`
}
