package main

import (
	"fmt"
	"go.uber.org/dig"
	"time"
)

func main() {
	timestr := "06.12.2019 09:12:48"
	layout := "02.01.2006 15:04:05"
	t, err := time.Parse(layout, timestr)
	if err != nil {
		fmt.Println(err)
		return
	}
	now := time.Now()
	fmt.Println(now.Sub(t))
}

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(NewConfig)
	_ = container.Provide(NewDB)
	_ = container.Provide(NewServer)
	return container
}
