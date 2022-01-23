package main

import (
	"context"
	"flag"
	"github.com/1005281342/gozerodemo/api/internal/config"
	"github.com/1005281342/gozerodemo/api/internal/handler"
	"github.com/1005281342/gozerodemo/api/internal/metrics"
	"github.com/1005281342/gozerodemo/api/internal/svc"
	"github.com/rookie-ninja/rk-entry/entry"
	"github.com/rookie-ninja/rk-zero/boot"
	"github.com/tal-tech/go-zero/core/conf"
)

var (
	configFile = flag.String("f", "etc/demo-api.yaml", "the config file")
)

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// Bootstrap basic entries from boot config.
	rkentry.RegisterInternalEntriesFromConfig("./boot.yaml")

	// Bootstrap zero entry from boot config
	res := rkzero.RegisterZeroEntriesWithConfig("./boot.yaml")

	// Get ZeroEntry
	zeroEntry := res["go-zero"].(*rkzero.ZeroEntry)
	metrics.Init(c.Namespace, zeroEntry.PromEntry.Registerer)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(zeroEntry.Server, ctx)

	// Bootstrap zero entry
	zeroEntry.Bootstrap(context.Background())

	// Wait for shutdown signal
	rkentry.GlobalAppCtx.WaitForShutdownSig()

	// Interrupt zero entry
	zeroEntry.Interrupt(context.Background())
}
