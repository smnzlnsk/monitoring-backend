package handlers

import (
	"fmt"
	"go.opentelemetry.io/collector/pdata/pmetric"
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

func (dio *DiskIODatapointHandler) handleGauge(metric pmetric.Metric) {
	for i := 0; i < metric.Gauge().DataPoints().Len(); i++ {
		dataPoint := metric.Gauge().DataPoints().At(i)
		fmt.Printf("%s, %s, %f/n", metric.Type().String(), metric.Name(), dataPoint.DoubleValue())
	}
}

func (dio *DiskIODatapointHandler) handleHistogram(metric pmetric.Metric) {
	for i := 0; i < metric.Histogram().DataPoints().Len(); i++ {
		dataPoint := metric.Histogram().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		fmt.Printf("%s, %s, %v/n", metric.Type().String(), metric.Name(), dataPoint.Attributes())
	}
}

func (dio *DiskIODatapointHandler) handleSummary(metric pmetric.Metric) {
	for i := 0; i < metric.Summary().DataPoints().Len(); i++ {
		dataPoint := metric.Summary().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		fmt.Printf("%s, %s, %v/n", metric.Type().String(), metric.Name(), dataPoint.Attributes())
	}
}
