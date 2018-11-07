package db

import (
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"

	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var (
	app       = beego.AppConfig.String("appname")
	Engine    *xorm.Engine
	HasEngine bool
)

func ConDb() (*xorm.Engine, error) {
	dbtype := beego.AppConfig.String("DBType")
	switch dbtype {
	case "mysql":
		return xorm.NewEngine("mysql", beego.AppConfig.String("mysql::DBConnect"))

	case "postgres":
		return xorm.NewEngine("postgres", beego.AppConfig.String("postgres::DBConnect"))

	case "memory":
		return xorm.NewEngine("tidb", "memory://tidb/tidb")

	case "goleveldb":
		if beego.AppConfig.String("goleveldb::DBConnect") != "" {
			return xorm.NewEngine("tidb", beego.AppConfig.String("goleveldb::DBConnect"))
		}
		return xorm.NewEngine("tidb", "goleveldb://"+app+"/data/tidb/tidb")

	case "boltdb":
		if beego.AppConfig.String("goleveldb::DBConnect") != "" {
			return xorm.NewEngine("tidb", beego.AppConfig.String("boltdb::DBConnect"))
		}
		return xorm.NewEngine("tidb", "boltdb://"+app+"/data/tidb/tidb")

	case "sqlite":
		if beego.AppConfig.String("sqlite::DBConnect") != "" {
			return xorm.NewEngine("sqlite3", beego.AppConfig.String("sqlite::DBConnect"))
		}
		return xorm.NewEngine("sqlite3", app+"/data/sqlite.db")
	}
	return nil, errors.New("Unknown database type..")
}

func SetEngine() (*xorm.Engine, error) {
	var _error error
	if Engine, _error = ConDb(); _error != nil {
		return nil, fmt.Errorf("Fail to connect to database: %s", _error.Error())
	}

	if err := Engine.Ping(); err != nil {
		return nil, err
	}

	Engine.SetColumnMapper(core.GonicMapper{})

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 10240)
	Engine.SetDefaultCacher(cacher)

	logDir := "logs"
	if _, e := os.Open(logDir); e != nil {
		os.MkdirAll(logDir, os.ModePerm)
	}

	logPath := path.Join(logDir, "xorm.log")
	f, err := os.Create(logPath)
	if err != nil {
		return Engine, fmt.Errorf("Fail to create xorm.log: %s", err.Error())
	}

	Engine.SetLogger(xorm.NewSimpleLogger(f))
	Engine.ShowSQL(false)

	if location, err := time.LoadLocation("Asia/Shanghai"); err == nil {
		Engine.TZLocation = location
	}

	return Engine, err

}

func init() {

	var _error error
	if Engine, _error = SetEngine(); _error != nil {
		log.Fatal("consumer.infrastructure.db.initdb.init() errors:", _error.Error())
	}

	_error = Engine.CreateTables(&User{}, &Category{}, &Node{}, &Topic{}, &Reply{}, &File{})
	if _error != nil {
		log.Fatal("consumer.infrastructure.db.initdb.init() CreateTables errors:", _error.Error())
	}
	//	Engine.CreateTables(&Category{})
	//	Engine.CreateTables(&Node{})

	inits()

}

func Ping() error {
	return Engine.Ping()
}

func inits() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("The db init success!")
}
