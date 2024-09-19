package handlers

import (
	"fmt"
	"go.opentelemetry.io/collector/pdata/pmetric"
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
		machine, _ := dataPoint.Attributes().Get("machine")
		switch dataPoint.ValueType().String() {
		case "Double":
			fmt.Printf("%s, %s, %s, %f, %s/n",
				metric.Type().String(),
				metric.Name(),
				dataPoint.ValueType().String(),
				dataPoint.DoubleValue(), machine.Str())
		case "Int":
			fmt.Printf("%s, %s, %s, %d, %s/n",
				metric.Type().String(),
				metric.Name(),
				dataPoint.ValueType().String(),
				dataPoint.IntValue(),
				machine.Str())
		default:
			fmt.Printf("unsupported datapoint type")
			return
		}
	}
}

func (file *FilesystemDatapointHandler) handleGauge(metric pmetric.Metric) {
	for i := 0; i < metric.Gauge().DataPoints().Len(); i++ {
		dataPoint := metric.Gauge().DataPoints().At(i)
		fmt.Printf("%s, %s, %f/n", metric.Type().String(), metric.Name(), dataPoint.DoubleValue())
	}
}

func (file *FilesystemDatapointHandler) handleHistogram(metric pmetric.Metric) {
	for i := 0; i < metric.Histogram().DataPoints().Len(); i++ {
		dataPoint := metric.Histogram().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		fmt.Printf("%s, %s, %v/n", metric.Type().String(), metric.Name(), dataPoint.Attributes())
	}
}

func (file *FilesystemDatapointHandler) handleSummary(metric pmetric.Metric) {
	for i := 0; i < metric.Summary().DataPoints().Len(); i++ {
		dataPoint := metric.Summary().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		fmt.Printf("%s, %s, %v/n", metric.Type().String(), metric.Name(), dataPoint.Attributes())
	}
}
