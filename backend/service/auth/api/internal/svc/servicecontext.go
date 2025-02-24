package svc

import (
	"chat_app/service/auth/api/internal/config"
	"chat_app/service/auth/model"
	"chat_app/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	ObjSync   sync.ObjSync
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		ObjSync:   *sync.NewObjSync(0),
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DataSource)),
	}
}
