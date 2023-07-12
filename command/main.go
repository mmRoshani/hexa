package main

import (
	"os"

	"github.com/mmRoshani/hexa/internal/adapters/app/api"
	"github.com/mmRoshani/hexa/internal/adapters/core/arithmetic"
	gRPC "github.com/mmRoshani/hexa/internal/adapters/framework/left/grpc"
	"github.com/mmRoshani/hexa/internal/adapters/framework/right/db"
	"github.com/mmRoshani/hexa/internal/ports"
)

func main() {
	var err error

	// env fetch

	dbaseDriver := os.Getenv("DB_DRIVER")
	dataSourceName := os.Getenv("DB_DS_NAME")

	// ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	// dep injection

	dbaseAdapter = db.NewAdapter(dbaseDriver, dataSourceName)
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)
	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()

}
