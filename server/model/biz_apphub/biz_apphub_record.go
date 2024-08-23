package biz_apphub

// biz_apphub 结构体  BizAppHubRecord
type BizAppHubRecord struct {
	AppId uint `json:"appId" form:"appId" gorm:"column:app_id;comment:应用id;"`
	BizAppHub
	//UpdatedUser string `json:"updatedUser"  gorm:"column:updatedUser;comment:操作用户;type:varchar(255);"`
}

// TableName biz_apphub BizAppHub自定义表名 biz_apphub
func (BizAppHubRecord) TableName() string {
	return "biz_apphub_record"
}
