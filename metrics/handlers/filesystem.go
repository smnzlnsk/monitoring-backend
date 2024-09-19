package handlers

import (
	"github.com/smnzlnsk/monitoring-backend/logging"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

var _ DatapointHandler = (*FilesystemDatapointHandler)(nil)

var filesystemHandler DatapointHandler

type FilesystemDatapointHandler struct{}

func newFilesystemDatapointHandler() DatapointHandler {
	filesystemHandler = &FilesystemDatapointHandler{}
	return filesystemHandler
}

func (file *FilesystemDatapointHandler) HandleMetric(metric pmetric.Metric) {
	// fetch handler function according to metric type and consume it
	file.getMetricTypeHandle(metric.Type().String())(metric)
}

func (file *FilesystemDatapointHandler) getMetricTypeHandle(metricType string) func(metric pmetric.Metric) {
	switch metricType {
	case "Sum":
		return file.handleSum
	case "Gauge":
		return file.handleGauge
	case "Histogram":
		return file.handleHistogram
	case "Summary":
		return file.handleSummary
	default:
		return nil
	}
}

func (file *FilesystemDatapointHandler) handleSum(metric pmetric.Metric) {
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

func (file *FilesystemDatapointHandler) handleGauge(metric pmetric.Metric) {
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

func (file *FilesystemDatapointHandler) handleHistogram(metric pmetric.Metric) {
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

func (file *FilesystemDatapointHandler) handleSummary(metric pmetric.Metric) {
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
