package model

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   databaseSetting.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,                        // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connStr, // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), gormConfig)
	//替换 gorm  create update 流程中自带的回调函数
	db.Callback().Create().Replace("gorm:before_create", updateTimeStampForBeforeCreateCallback)
	db.Callback().Update().Replace("gorm:before_update", updateTimeStampForBeforeUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete",deleteCallback)
	if err != nil {
		log.Printf("connstr :%v  err: %+v \n", connStr, err.Error())
		panic(err)
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}
	return db, nil

}

func updateTimeStampForBeforeCreateCallback(db *gorm.DB) {
	db.Statement.SetColumn("CreatedOn", time.Now().Unix())
}

func updateTimeStampForBeforeUpdateCallback(db *gorm.DB) {
	db.Statement.SetColumn("ModifiedOn", time.Now().Unix())

}

//   实现软删除
func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		if db.Statement.Schema != nil {
			db.Statement.SQL.Grow(100)
			deleteField := db.Statement.Schema.LookUpField("DeletedOn")
			isDelField := db.Statement.Schema.LookUpField("IsDel")
			if !db.Statement.Unscoped && deleteField != nil && isDelField != nil {
				//soft delete
				if db.Statement.SQL.String() == "" {
					nowTime := time.Now().Unix()
					db.Statement.AddClause(
						clause.Set{
							{Column: clause.Column{Name: deleteField.DBName},
								Value: nowTime},
							{Column: clause.Column{Name: isDelField.DBName},
								Value: 1},
						})
					db.Statement.AddClauseIfNotExists(clause.Update{})
					db.Statement.Build("UPDATE", "SET", "WHERE")
				}
			} else {
				//delete
				if db.Statement.SQL.String() == "" {
					db.Statement.AddClauseIfNotExists(clause.Delete{})
					db.Statement.AddClauseIfNotExists(clause.From{})
					db.Statement.Build("DELETE", "SET", "WHERE")
				}
			}
			//must need where
			if _, ok := db.Statement.Clauses["WHERE"]; !db.AllowGlobalUpdate && !ok {
				db.AddError(gorm.ErrMissingWhereClause)
				return
			}
			// exec sql
			if !db.DryRun && db.Error == nil {
				result, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
				if err != nil {
					db.RowsAffected, _ = result.RowsAffected()
				} else {
					db.AddError(err)
				}
			}
		}
	}

}
