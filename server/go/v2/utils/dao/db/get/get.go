package get

import (
	"fmt"
	"log"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/liov/hoper/go/v2/utils/fs"
)

var (
	OrmDB *gorm.DB
)

func init() {
	path, err := fs.FindFile("add-config.toml")
	if err != nil {
		log.Fatal(err)
	}
	var config = struct {
		GetOrm struct {
			Type     string
			User     string
			Password string
			Host     string
		}
	}{}
	err = configor.New(&configor.Config{Debug: true}).
		Load(&config, path)
	if err != nil {
		log.Fatal(err)
	}
	var url string
	if config.GetOrm.Type == "mysql" {
		url = fmt.Sprintf("%s:%s@tcp(%s:3306)/test?charset=utf8mb4&parseTime=True&loc=Local", config.GetOrm.User, config.GetOrm.Password, config.GetOrm.Host)
	} else if config.GetOrm.Type == "postgres" {
		url = fmt.Sprintf("host=%s user=%s dbname=test sslmode=disable password=%s",
			config.GetOrm.Host, config.GetOrm.User, config.GetOrm.Password)
	}

	OrmDB, err = gorm.Open(config.GetOrm.Type, url)
	if err != nil {
		log.Fatalln(err)
	}
	OrmDB.SingularTable(true)
	OrmDB.LogMode(true)
}