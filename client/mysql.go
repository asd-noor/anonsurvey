package client

import (
	"anonsurvey/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	*gorm.DB
}

func (m *MySQL) ConnectDB(conf config.DBConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.User, conf.Pass, conf.Host, conf.Port, conf.Schema)

	g, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	m.DB = g

	return nil
}
