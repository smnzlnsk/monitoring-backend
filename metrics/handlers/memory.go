package handlers

import (
	"github.com/smnzlnsk/monitoring-backend/logging"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

var _ DatapointHandler = (*MemoryDatapointHandler)(nil)

var memoryHandler DatapointHandler

type MemoryDatapointHandler struct{}

func newMemoryDatapointHandler() DatapointHandler {
	memoryHandler = &MemoryDatapointHandler{}
	return memoryHandler
}

func (mem *MemoryDatapointHandler) HandleMetric(metric pmetric.Metric) {
	// fetch handler function according to metric type and consume it
	mem.getMetricTypeHandle(metric.Type().String())(metric)
}

func (mem *MemoryDatapointHandler) getMetricTypeHandle(metricType string) func(metric pmetric.Metric) {
	switch metricType {
	case "Sum":
		return mem.handleSum
	case "Gauge":
		return mem.handleGauge
	case "Histogram":
		return mem.handleHistogram
	case "Summary":
		return mem.handleSummary
	default:
		return nil
	}
}

func (mem *MemoryDatapointHandler) handleSum(metric pmetric.Metric) {
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

func (mem *MemoryDatapointHandler) handleGauge(metric pmetric.Metric) {
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

func (mem *MemoryDatapointHandler) handleHistogram(metric pmetric.Metric) {
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

func (mem *MemoryDatapointHandler) handleSummary(metric pmetric.Metric) {
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
