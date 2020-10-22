package main

import (
	"context"

	"github.com/godcong/chronos"
	"github.com/godcong/fate"
	"github.com/godcong/fate/config"
)

func main() {
	//出生日期
	born := chronos.New("2020/12/28")
	//姓氏
	lastName := "方"
	cfg := config.DefaultConfig()
	cfg.BaguaFilter = false
	cfg.ZodiacFilter = false
	cfg.SupplyFilter = false
	cfg.HardFilter = false
	cfg.StrokeMin = 3
	cfg.StrokeMax = 24
	cfg.Database = config.Database{
		Host:         "localhost",
		Port:         "3306",
		User:         "root",
		Pwd:          "123456",
		Name:         "fate",
		MaxIdleCon:   0,
		MaxOpenCon:   0,
		Driver:       "mysql",
		File:         "",
		Dsn:          "",
		ShowSQL:      false,
		ShowExecTime: false,
	}
	cfg.FileOutput = config.FileOutput{
		OutputMode: config.OutputModeLog,
		Path:       "name.log",
	}

	f := fate.NewFate(lastName, born.Solar().Time(), fate.ConfigOption(cfg))

	e := f.MakeName(context.Background())
	if e != nil {
		panic(e)
	}
}
