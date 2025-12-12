//go:build wireinject
// +build wireinject

package main

import (
	"emqx-kafka/internal/ioc"
	"emqx-kafka/internal/server"

	"github.com/google/wire"
)

func InitBridge() *server.Bridge {
	panic(wire.Build(
		ioc.LoadConfig,
		ioc.InitSarmaClient,
		server.NewBridge,
	))
	return &server.Bridge{}
}
