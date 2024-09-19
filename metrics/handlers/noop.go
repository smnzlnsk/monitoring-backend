package handlers

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

var _ DatapointHandler = (*NoopDatapointHandler)(nil)

type NoopDatapointHandler struct{}

func newNoopDatapointHandler() DatapointHandler {
	return &NoopDatapointHandler{}
}

func (noop *NoopDatapointHandler) HandleMetric(metric pmetric.Metric) {
	// fetch handler function according to metric type and consume it
	noop.getMetricTypeHandle(metric.Type().String())(metric)
}

func (noop *NoopDatapointHandler) getMetricTypeHandle(metricType string) func(metric pmetric.Metric) {
	switch metricType {
	case "Sum":
		return noop.handleSum
	case "Gauge":
		return noop.handleGauge
	case "Histogram":
		return noop.handleHistogram
	case "Summary":
		return noop.handleSummary
	default:
		return nil
	}
}

func (noop *NoopDatapointHandler) handleSum(_ pmetric.Metric) {
}

func (noop *NoopDatapointHandler) handleGauge(_ pmetric.Metric) {
}

func (noop *NoopDatapointHandler) handleHistogram(_ pmetric.Metric) {
}

func (noop *NoopDatapointHandler) handleSummary(_ pmetric.Metric) {
}
