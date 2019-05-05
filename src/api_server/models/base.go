package models

import (
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// init function
func init() {
	DBMS := "mysql"
	USER := "root"
	PASS := "intellilink"
	PROTOCOL := "tcp(0.0.0.0:3306)"
	DBNAME := "anitya"
	OPTIONS := "parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTIONS

	var err error
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(&Project{})
	db.Debug().AutoMigrate(&Role{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Debug().AutoMigrate(&User{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE").AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE")
	db.Debug().AutoMigrate(&ProbePoint{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Debug().AutoMigrate(&DashboardConfig{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Debug().AutoMigrate(&HTTPMonitoringConfig{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Debug().AutoMigrate(&DNSMonitoringConfig{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Debug().AutoMigrate(&NotificationHistory{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Debug().AutoMigrate(&NotificationSlack{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
}

// GetDB function
func GetDB() *gorm.DB {
	return db
}
