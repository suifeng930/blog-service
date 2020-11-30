package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         int64  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"

}



type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}

func (a Article)Create(db *gorm.DB) error  {

	err := db.Create(&a).Error
	if err!=nil {
		return errors.Wrap(err,"create article table is fail")
	}
	return nil

}

func (a Article) Delete(db *gorm.DB)error  {

	err := db.Where("id=? and is_del=?", a.ID, 0).Delete(&a).Error
	if err!=nil {
		return errors.Wrap(err,"delete article table is fail")

	}
	return nil

}

func (a Article) Update(db *gorm.DB,values interface{}) error {

	err :=db.Model(&Article{}).Where("id=? and is_del=?",a.ID,0).Updates(values).Error
	if err!=nil {
	log.Println("db.Model(&Tag{}).Where(\"id=? and is_del=?\",t.ID).Updates(values).Error ",err)
	return errors.Wrap(err,"update article table is fail")

	}
	return nil

}

func (a Article) GetArticleByIdAndState(db *gorm.DB) (Article,error) {
	var article Article
	db = db.Where("id=? and state=? and is_del=?", a.ID, a.State, 0)
	err := db.Find(&article).Error
	if err!=nil && err!=gorm.ErrRecordNotFound {
		return article, errors.Wrap(err,"get Article is fail ")
	}
	return article,nil

}
