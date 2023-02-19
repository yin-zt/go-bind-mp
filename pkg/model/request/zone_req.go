package request

type ZoneListReq struct {
	ZoneName string `json:"zonename" form:"zonename"`
	AllowIps string `json:"allowips" form:"allowips"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type ZoneAddReq struct {
	ZoneName  string `json:"zonename" validate:"required"`
	AllowIps  string `json:"allowips" validate:"required"`
	SerialNu  uint   `json:"serialnu" validate:"required"`
	ViewsId   []uint `json:"views_id" validate:"required"`
	DomainsId []uint `json:"domains_id" validate:"required"`
}

type ZoneDeleteReq struct {
	ZoneIds []uint `json:"zoneids" validate:"required"`
}
