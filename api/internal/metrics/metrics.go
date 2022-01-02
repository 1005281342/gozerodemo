package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	rkprom "github.com/rookie-ninja/rk-prom"
)

var metricsSet *rkprom.MetricsSet

const (
	helloFailed = "hello_failed"
	hello       = "hello"
)

func Init(namespace string, registerer prometheus.Registerer) {
	metricsSet = rkprom.NewMetricsSet(namespace, "demo", registerer)

	var err = metricsSet.RegisterCounter(helloFailed)
	if err != nil {
		panic(err)
	}
	err = metricsSet.RegisterCounter(hello)
	if err != nil {
		panic(err)
	}
}

func GetMetricsSet() *rkprom.MetricsSet {
	return metricsSet
}

func HelloFailed() {
	metricsSet.GetCounterWithValues(helloFailed).Inc()
}

func Hello() {
	metricsSet.GetCounterWithValues(hello).Inc()
}
