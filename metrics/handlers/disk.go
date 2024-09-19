package handlers

import (
	"github.com/smnzlnsk/monitoring-backend/logging"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

var _ DatapointHandler = (*DiskIODatapointHandler)(nil)

var diskIOHandler DatapointHandler

type DiskIODatapointHandler struct{}

func newDiskIODatapointHandler() DatapointHandler {
	diskIOHandler = &DiskIODatapointHandler{}
	return diskIOHandler
}

func (dio *DiskIODatapointHandler) HandleMetric(metric pmetric.Metric) {
	// fetch handler function according to metric type and consume it
	dio.getMetricTypeHandle(metric.Type().String())(metric)
}
func (dio *DiskIODatapointHandler) getMetricTypeHandle(metricType string) func(metric pmetric.Metric) {
	switch metricType {
	case "Sum":
		return dio.handleSum
	case "Gauge":
		return dio.handleGauge
	case "Histogram":
		return dio.handleHistogram
	case "Summary":
		return dio.handleSummary
	default:
		return nil
	}
}

func (dio *DiskIODatapointHandler) handleSum(metric pmetric.Metric) {
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

func (dio *DiskIODatapointHandler) handleGauge(metric pmetric.Metric) {
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

func (dio *DiskIODatapointHandler) handleHistogram(metric pmetric.Metric) {
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

func (dio *DiskIODatapointHandler) handleSummary(metric pmetric.Metric) {
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
