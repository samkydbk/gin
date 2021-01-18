package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connection() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "z_", //表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, //使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		//SkipDefaultTransaction: true  //跳过默认事务
		//DisableAutomaticPing: true,  //检查数据库的可用性 true禁用
		//DisableForeignKeyConstraintWhenMigrating: true,  //禁用外键约束
		//更多 https://gorm.io/zh_CN/docs/gorm_config.html
	})
	if err != nil {
		panic(err)
	}
	return db
}
