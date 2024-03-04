package volcengine

import (
	"context"

	operatorMetrics "github.com/cilium/cilium/operator/metrics"
	operatorOption "github.com/cilium/cilium/operator/option"
	"github.com/cilium/cilium/pkg/ipam"
	ipamMetrics "github.com/cilium/cilium/pkg/ipam/metrics"
	"github.com/cilium/cilium/pkg/logging"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/metrics"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "ipam-allocator-volc-engine")

type AllocatorVolcEngine struct {
	// TODO: api client
}

func (a *AllocatorVolcEngine) Init(ctx context.Context) error {
	// TODO: init API.
	return nil
}

func (a *AllocatorVolcEngine) Start(ctx context.Context) error {
	var iMetrics ipam.MetricsAPI

	log.Info("Starting VolcEngine ENI allocator...")

	if operatorOption.Config.EnableMetrics {
		iMetrics = ipamMetrics.NewPrometheusMetrics(metrics.Namespace, "volcengine", operatorMetrics.Registry)
	} else {
		iMetrics = &ipamMetrics.NoOpMetrics{}
	}
	return nil
}
