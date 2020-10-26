package datas

import (
	"github.com/gookit/color"
	"time"

	"gin-vue-admin/model"
	"gorm.io/gorm"
)

var Authorities = []model.SysAuthority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "普通用户", ParentId: "0"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "普通用户子角色", ParentId: "888"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: "测试角色", ParentId: "0"},
}

func InitSysAuthority(db *gorm.DB) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
			color.Danger.Println("sys_authorities表的初始数据已存在!")
			return nil
		}
		if tx.Create(&Authorities).Error != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}
