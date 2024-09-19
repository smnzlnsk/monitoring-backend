package handlers

import (
	"github.com/smnzlnsk/monitoring-backend/logging"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

var _ DatapointHandler = (*NetworkDatapointHandler)(nil)

var networkHandler DatapointHandler

type NetworkDatapointHandler struct{}

func newNetworkDatapointHandler() DatapointHandler {
	networkHandler = &NetworkDatapointHandler{}
	return networkHandler
}

func (net *NetworkDatapointHandler) HandleMetric(metric pmetric.Metric) {
	// fetch handler function according to metric type and consume it
	net.getMetricTypeHandle(metric.Type().String())(metric)
}

func (net *NetworkDatapointHandler) getMetricTypeHandle(metricType string) func(metric pmetric.Metric) {
	switch metricType {
	case "Sum":
		return net.handleSum
	case "Gauge":
		return net.handleGauge
	case "Histogram":
		return net.handleHistogram
	case "Summary":
		return net.handleSummary
	default:
		return nil
	}
}

func (net *NetworkDatapointHandler) handleSum(metric pmetric.Metric) {
	for i := 0; i < metric.Sum().DataPoints().Len(); i++ {
		dataPoint := metric.Sum().DataPoints().At(i)
		switch dataPoint.ValueType().String() {
		case "Double":
			logging.Logger.Info("",
				zap.String("name", metric.Name()),
				zap.Any("attributes", dataPoint.Attributes().AsRaw()),
				zap.String("type", dataPoint.ValueType().String()),
				zap.Any("value", dataPoint.DoubleValue()),
			)
		case "Int":
			logging.Logger.Info("",
				zap.String("name", metric.Name()),
				zap.Any("attributes", dataPoint.Attributes().AsRaw()),
				zap.String("type", dataPoint.ValueType().String()),
				zap.Any("value", dataPoint.IntValue()),
			)
		default:
			logging.Logger.Warn("unsupported datapoint type")
			return
		}
	}
}

func (net *NetworkDatapointHandler) handleGauge(metric pmetric.Metric) {
	for i := 0; i < metric.Gauge().DataPoints().Len(); i++ {
		dataPoint := metric.Gauge().DataPoints().At(i)
		logging.Logger.Info("",
			zap.String("name", metric.Name()),
			zap.Any("attributes", dataPoint.Attributes().AsRaw()),
			zap.String("type", dataPoint.ValueType().String()),
			zap.Any("value", dataPoint.IntValue()),
		)
	}
}

func (net *NetworkDatapointHandler) handleHistogram(metric pmetric.Metric) {
	for i := 0; i < metric.Histogram().DataPoints().Len(); i++ {
		dataPoint := metric.Histogram().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		logging.Logger.Info("",
			zap.String("name", metric.Name()),
			zap.Any("attributes", dataPoint.Attributes().AsRaw()),
		)
	}
}

func (net *NetworkDatapointHandler) handleSummary(metric pmetric.Metric) {
	for i := 0; i < metric.Summary().DataPoints().Len(); i++ {
		dataPoint := metric.Summary().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		logging.Logger.Info("",
			zap.String("name", metric.Name()),
			zap.Any("attributes", dataPoint.Attributes().AsRaw()),
		)
	}
}
