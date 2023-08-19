package infra

import (
	"fmt"
	"go-gacha-system/pkg/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const driverName = "mysql"

type MySQLConnector struct {
	DB *sqlx.DB
}

func NewMySQLConnector() *MySQLConnector {
	conf := config.LoadDBConfig()
	dsn := mysqlConnDSN(*conf)
	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	return &MySQLConnector{
		DB: db,
	}
}

func mysqlConnDSN(mysqlInfo config.MySQLInfo) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		mysqlInfo.MySQLUser,
		mysqlInfo.MySQLPassword,
		mysqlInfo.MySQLHost,
		mysqlInfo.MySQLPort,
		mysqlInfo.MySQLDBName)

	return dsn
}
