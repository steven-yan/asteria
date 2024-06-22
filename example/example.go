package main

import (
	"github.com/mylxsw/asteria/filter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/log"
)

func main() {
	log.AddGlobalFilter(filter.WithStacktrace(level.Error))
	//log.Debug("Hello, world")
	//log.Warning("Hello, world")
	//log.F(log.M{
	//	"key": "value",
	//}).Errorf("oops")
	log.WithFields(log.Fields{
		"key123": "valueabc",
	}).Error("oops")
}
