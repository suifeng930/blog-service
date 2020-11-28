package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"log"
)


/**
 主要定义了Request 结构体作为接口入参的基准。 由于本项目并不复杂，所以直接吧Request 结构体放在了service层一便于使用
若后续业务不断增长，程序越来越复杂，service层也变得冗余，则可以考虑抽离一层接口校验层。以方便解耦
 */

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service)CountTag	(param *CountTagRequest) (int64,error)  {

	return svc.dao.CountTag(param.Name,int64(param.State))

}

func (svc *Service) GetTagList(param *TagListRequest,pager *app.Pager) ([]*model.Tag,error) {
	return svc.dao.GetTagList(param.Name,int64(param.State),pager.Page,pager.PageSize)

}


func (svc *Service) CreateTag(param *CreateTagRequest) error {
	err := svc.dao.CreateTag(param.Name, int64(param.State), param.CreatedBy)
	log.Println("CreateTag(param *CreateTagRequest) :",err)
	return err

}


func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID,param.Name,int64(param.State),param.ModifiedBy)

}


func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)

}
