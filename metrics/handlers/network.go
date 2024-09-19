package handlers

import (
	"fmt"
	"go.opentelemetry.io/collector/pdata/pmetric"
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

func (net *NetworkDatapointHandler) handleGauge(metric pmetric.Metric) {
	for i := 0; i < metric.Gauge().DataPoints().Len(); i++ {
		_ = metric.Gauge().DataPoints().At(i)
		// fmt.Printf("%s, %s, %f/n", metric.Type().String(), metric.Name(), dataPoint.DoubleValue())
	}
}

func (net *NetworkDatapointHandler) handleHistogram(metric pmetric.Metric) {
	for i := 0; i < metric.Histogram().DataPoints().Len(); i++ {
		dataPoint := metric.Histogram().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		fmt.Printf("%s, %s, %v/n", metric.Type().String(), metric.Name(), dataPoint.Attributes())
	}
}

func (net *NetworkDatapointHandler) handleSummary(metric pmetric.Metric) {
	for i := 0; i < metric.Summary().DataPoints().Len(); i++ {
		dataPoint := metric.Summary().DataPoints().At(i)
		// TODO: Implement me!
		// ! Currently not contained in the metrics data package we receive
		fmt.Printf("%s, %s, %v/n", metric.Type().String(), metric.Name(), dataPoint.Attributes())
	}
}
