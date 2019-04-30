package db

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

func ConnectGORM() *gorm.DB {
    DBMS     := "mysql"
    USER     := "root"
    PASS     := "intellilink"
    PROTOCOL := "tcp(0.0.0.0:3306)"
    DBNAME   := "anitya"

    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
    db,err := gorm.Open(DBMS, CONNECT)

    if err != nil {
        panic(err.Error())
    }
    return db
}
