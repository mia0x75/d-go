package g

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type DBPool struct {
	Falcon    *xorm.Engine
	Graph     *xorm.Engine
	Uic       *xorm.Engine
	Dashboard *xorm.Engine
	Alarm     *xorm.Engine
}

type DailInfo struct {
	Host     string
	User     string
	Password string
	Port     int
	DbName   string
	Options  string
	Debug    bool
}

var (
	dbp               DBPool
	dailInfo          DailInfo
	format            string = `%s:%s@tcp(%s:%d)/%s`
	formatWithOptions string = `%s:%s@tcp(%s:%d)/%s?%s`
)

func Con() DBPool {
	return dbp
}

func NewEngine(d DailInfo) (*xorm.Engine, error) {
	var dns string
	if len(d.Options) > 0 {
		dns = fmt.Sprintf(formatWithOptions, d.User, d.Password, d.Host, d.Port, d.DbName, d.Options)
	} else {
		dns = fmt.Sprintf(format, d.User, d.Password, d.Host, d.Port, d.DbName)
	}
	engine, err := xorm.NewEngine("mysql", dns)
	if err != nil {
		return nil, err
	}

	engine.Logger().SetLevel(core.LOG_ERR)
	engine.Ping()
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(10)
	engine.ShowSQL(d.Debug)
	return engine, nil
}

func InitDB(vip *viper.Viper) (err error) {
	mapstructure.Decode(vip.GetStringMap("db.portal"), &dailInfo)
	portal, err := NewEngine(dailInfo)
	if err != nil {
		return fmt.Errorf("connect to falcon_portal: %s", err.Error())
	}
	dbp.Falcon = portal

	mapstructure.Decode(vip.GetStringMap("db.graph"), &dailInfo)
	graphd, err := NewEngine(dailInfo)
	if err != nil {
		return fmt.Errorf("connect to graph: %s", err.Error())
	}
	dbp.Graph = graphd

	mapstructure.Decode(vip.GetStringMap("db.uic"), &dailInfo)
	uicd, err := NewEngine(dailInfo)
	if err != nil {
		return fmt.Errorf("connect to uic: %s", err.Error())
	}
	dbp.Uic = uicd

	mapstructure.Decode(vip.GetStringMap("db.dashboard"), &dailInfo)
	dashd, err := NewEngine(dailInfo)
	if err != nil {
		return fmt.Errorf("connect to dashboard: %s", err.Error())
	}
	dbp.Dashboard = dashd

	mapstructure.Decode(vip.GetStringMap("db.alarms"), &dailInfo)
	almd, err := NewEngine(dailInfo)
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
