package handlers

import (
	"fmt"
	"go.opentelemetry.io/collector/pdata/pmetric"
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

func (mem *MemoryDatapointHandler) handleGauge(metric pmetric.Metric) {
	for i := 0; i < metric.Gauge().DataPoints().Len(); i++ {
		_ = metric.Gauge().DataPoints().At(i)
		// fmt.Printf("%s, %s, %f/n", metric.Type().String(), metric.Name(), dataPoint.DoubleValue())
	}
}

func (mem *MemoryDatapointHandler) handleHistogram(metric pmetric.Metric) {
	for i := 0; i < metric.Histogram().DataPoints().Len(); i++ {
		dataPoint := metric.Histogram().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		fmt.Printf("%s, %s, %v/n", metric.Type().String(), metric.Name(), dataPoint.Attributes())
	}
}

func (mem *MemoryDatapointHandler) handleSummary(metric pmetric.Metric) {
	for i := 0; i < metric.Summary().DataPoints().Len(); i++ {
		dataPoint := metric.Summary().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		fmt.Printf("%s, %s, %v/n", metric.Type().String(), metric.Name(), dataPoint.Attributes())
	}
}
