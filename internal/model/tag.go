package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"gorm.io/gorm"
	"log"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State int64  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"

}

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}



/**
 * Model : 指定运行DB 操作的模型实例，默认解析该结构体的名字为表名，格式为大写驼峰转小写下划线驼峰，若特殊情况，也可以编写该结构体的TableName()
 * 用于返回其对应的表名
 * Where : 设置筛选条件， 接受map \struct \ string 作为条件
 * Offset() : 偏移量，用于指定开始返回记录之前要跳过的记录数
 * Limit() : 限制检索的记录数
 * Find() : 查找符合筛选条件的记录
 * Update() : 更新所选字段
 * Delete() : 删除数据
 * Count() : 统计行为，用于统计模型的记录数
 *
 */

func (t Tag) Count(db *gorm.DB) (int64,error) {
	var count int64
	if t.Name!="" {
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?", t.State)
	err := db.Model(&t).Where("is_del=?", 0).Count(&count).Error
	if err!=nil {
		return 0, err
	}
	return count,nil

}
func (t Tag) List(db *gorm.DB,pageOffset,pageSize int) ([]*Tag,error)  {
	var tags []*Tag
	var err error
	if pageOffset>=0 &&pageSize >0	 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name!=""{
		db =db.Where("name=?",t.Name)
	}
	db =db.Where("state=?",t.State)
	if  err=db.Where("is_del=?", 0).Find(&tags).Error; err!=nil {
		return nil, err
	}
	return tags,nil

}

func (t Tag)Create(db *gorm.DB)error  {
	return db.Create(&t).Error

}

func (t Tag)Update	(db *gorm.DB,values interface{}) error  {
	err :=db.Model(&Tag{}).Where("id=? and is_del=?",t.ID,0).Updates(values).Error
	if err!=nil {
		log.Println("db.Model(&Tag{}).Where(\"id=? and is_del=?\",t.ID).Updates(values).Error ",err)
		return err

	}
	return nil

}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del=?",t.Model.ID,0).Delete(&t).Error

}