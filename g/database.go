package g

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
)

type DBPool struct {
	Falcon    *xorm.Engine
	Graph     *xorm.Engine
	Uic       *xorm.Engine
	Dashboard *xorm.Engine
	Alarm     *xorm.Engine
}

var (
	dbp    DBPool
	format string = `%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`
)

func Con() DBPool {
	return dbp
}

func NewEngine(user string, password string, host string, port int, db string) (*xorm.Engine, error) {
	dns := fmt.Sprintf(format, user, password, host, port, db)
	engine, err := xorm.NewEngine("mysql", dns)
	if err != nil {
		return nil, err
	}

	engine.Logger().SetLevel(core.LOG_ERR)
	engine.Ping()
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(10)
	// if g.AppConfig().Debug {
	// 	engine.ShowSQL(true)
	// } else {
	// 	engine.ShowSQL(false)
	// }
	return engine, nil
}

func InitDB(loggerlevel bool, vip *viper.Viper) (err error) {
	// var p *sql.DB
	portal, err := NewEngine("", "", "", 3306, "")
	if err != nil {
		return fmt.Errorf("connect to falcon_portal: %s", err.Error())
	}
	dbp.Falcon = portal

	// var g *sql.DB
	graphd, err := NewEngine("", "", "", 3306, "")
	if err != nil {
		return fmt.Errorf("connect to graph: %s", err.Error())
	}
	dbp.Graph = graphd

	// var u *sql.DB
	uicd, err := NewEngine("", "", "", 3306, "")
	if err != nil {
		return fmt.Errorf("connect to uic: %s", err.Error())
	}
	dbp.Uic = uicd

	// var d *sql.DB
	dashd, err := NewEngine("", "", "", 3306, "")
	if err != nil {
		return fmt.Errorf("connect to dashboard: %s", err.Error())
	}
	dbp.Dashboard = dashd

	// var alm *sql.DB
	almd, err := NewEngine("", "", "", 3306, "")
	if err != nil {
		return fmt.Errorf("connect to alarms: %s", err.Error())
	}
	dbp.Alarm = almd

	return
}

func CloseDB() (err error) {
	err = dbp.Falcon.Close()
	if err != nil {
		return
	}
	err = dbp.Graph.Close()
	if err != nil {
		return
	}
	err = dbp.Uic.Close()
	if err != nil {
		return
	}
	err = dbp.Dashboard.Close()
	if err != nil {
		return
	}
	err = dbp.Alarm.Close()
	if err != nil {
		return
	}
	return
}
