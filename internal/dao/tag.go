package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

func (d *Dao) CountTag(name string,state int64)(int64,error)  {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)

}

func (d *Dao)GetTagList(name string,state int64,page,pageSize int) ([]*model.Tag,error)  {

	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)

	return tag.List(d.engine,pageOffset,pageSize)

}

func (d *Dao)CreateTag(name string,state int64, createBy string)error  {
	tag := model.Tag{

		Model: &model.Model{CreatedBy: createBy},
		Name:  name,
		State: state,
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state int64, modifiedBy string) error {
	tag :=model.Tag{
		Model:&model.Model{ID: id},
	}
	values :=map[string]interface{}{
		"state":state,
		"modified_by":modifiedBy,
	}
	if name!="" {
		values["name"]=name

	}
	return tag.Update(d.engine,values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag :=model.Tag{
		Model:&model.Model{ID: id},
	}
	return tag.Delete(d.engine)
}

func (d *Dao) GetTag(id uint32, state int64) (model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}