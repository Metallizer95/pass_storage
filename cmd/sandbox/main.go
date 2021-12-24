package main

import (
	"go.uber.org/dig"
)

func main() {

}

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(NewConfig)
	_ = container.Provide(NewDB)
	_ = container.Provide(NewServer)
	return container
}
