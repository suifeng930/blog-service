package dao

import "github.com/go-programming-tour-book/blog-service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth,error) {

	auth := model.Auth{AppSecret: appSecret, AppKey: appKey}
	return auth.Get(d.engine)

}
