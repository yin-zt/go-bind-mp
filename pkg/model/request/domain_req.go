package request

type DomainListReq struct {
	DomainRecord string `json:"doaminrecord" form:"domainrecord"`
	Type         string `json:"type" form:"type"`
	Status       string `json:"status" form:"status"`
	PageNum      int    `json:"pageNum" form:"pageNum"`
	PageSize     int    `json:"pageSize" form:"pageSize"`
}

type DomainAddReq struct {
	DomainRecord string `json:"doaminrecord" validate:"excludes=_,required"`
	Type         string `json:"type" validate:"oneof=1 2 3,required"`
	Resolution   string `json:"resolution" validate:"required"`
	Monitor      string `json:"monitor" validate:"oneof=1 2,required"`
	Protocol     string `json:"protocal" validate:"oneof=http tcp,required"`
	Port         uint   `json:"port" validate:"lt=65535"`
	Notify       string `json:"notify" validate:"required"`
	BelongSystem string `json:"belongsystem"`
	BelongViewId []uint `json:"belongviewid" validate:"required"`
	BelongZoneId []uint `json:"belongzoneid" validate:"required"`
}

// DomainDeleteReq 批量删除资源结构体
type DoaminDeleteReq struct {
	DomainIds []uint `json:"domainids" validate:"required"`
}

// DomainUpdateReq 更新资源结构体; 只可更新不影响dns服务的域名信息
type DomainUpdateReq struct {
	ID           uint   `json:"id" validate:"required"`
	Monitor      string `json:"monitor"`
	Protocol     string `json:"protocol"`
	Port         uint   `json:"port"`
	Notify       string `json:"notify"`
	BelongSystem string `json:"belong_system"`
}
