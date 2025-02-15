package config

import (
	"github.com/mahdi-cpp/api-go-emqx/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var DB *gorm.DB
var err error

// docker run --name postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin@123456 -e POSTGRES_DB=mqtt -p 5432:5432 -d postgres

func DatabaseInit() {
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=PostgreSQL user=admin password=admin@123456 dbname=mqtt port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "api_v1.", // schema name
			SingularTable: false,
		}})

	if err != nil {
		println("Failed to connect database gallery\"")
		os.Exit(1)
	}

	err := DB.AutoMigrate(&model.Node{})
	if err != nil {
		return
	}

	err = DB.AutoMigrate(&model.Temperature{})
	if err != nil {
		return
	}

}
