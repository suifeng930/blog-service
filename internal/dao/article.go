package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
)

func (d *Dao)CreateArticle(param *model.Article) error {

	article :=model.Article{
		Title: param.Title,
		Desc: param.Desc,
		Content: param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State: param.State,
		Model: &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *model.Article) error{

	article :=model.Article{Model:&model.Model{ID: param.ID}}
	values :=map[string]interface{}{
		"modified_by":param.ModifiedBy,
		"state":param.State,
	}
	if param.Title!="" {
		values["title"]=param.Title
	}
	if param.CoverImageUrl!="" {
		values["cover_image_url"]=param.CoverImageUrl
	}
	if param.Desc!="" {
		values["desc"]=param.Desc
	}
	if param.Content!="" {
		values["content"]=param.Content
	}
	return article.Update(d.engine,values)


}

func (d *Dao)GetArticle(id uint32,state int64	) (model.Article,error)  {
	article :=model.Article{Model:&model.Model{ID: id},State: state}
	return article.GetArticleByIdAndState(d.engine)

}


func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}
