package request

type ViewListReq struct {
	ViewName string `json:"viewname" form:"viewname"`
	Acl      string `json:"acl" form:"acl"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type ViewAddReq struct {
	ViewName  string `json:"viewname" form:"viewname"`
	Acl       string `json:"acl" form:"acl"`
	Remark    string `json:"remark" form:"remark"`
	ZonesId   []uint `json:"zones_id" validate:"required"`
	DomainsId []uint `json:"domains_id" validate:"required"`
}

type ViewDeleteReq struct {
	ViewIds []uint `json:"viewids" validate:"required"`
}

// ViewUpdateReq 更新视图资源结构体; 只可更新不影响dns服务的视图信息
type ViewUpdateReq struct {
	ID     uint   `json:"id" validate:"required"`
	Remark string `json:"remark"`
}
