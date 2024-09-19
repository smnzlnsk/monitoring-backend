package metrics

import (
	"fmt"
	"github.com/smnzlnsk/monitoring-backend/metrics/handlers"
	"strings"
)

// In this file we define the decoding of the pmetric.Metrics package into their data points

type MetricDecoder interface {
	getDatapointHandler(string) handlers.DatapointHandler
	DecodeMetrics(rawMetrics []byte) error
}

var _ MetricDecoder = (*OakestraMetricDecoder)(nil)

var metricDecoder MetricDecoder

func GetDecoder() MetricDecoder {
	if metricDecoder == nil {
		panic("metric decoder not initialized")
	}
	return metricDecoder
}

type OakestraMetricDecoder struct {
	datapointHandlerHolder handlers.DatapointHandlerHolder
	// create map with function references for quick dereferencing below
	handlerCategoryTable map[string]handlers.DatapointHandler
}

func InitOakestraDecoder() error {
	// first set up datapoint handlers
	err := handlers.InitOakestraDatapointHandlerHolder()
	if err != nil {
		return err
	}

	omd := &OakestraMetricDecoder{
		datapointHandlerHolder: handlers.GetDatapointHandlerHolder(),
		handlerCategoryTable:   make(map[string]handlers.DatapointHandler),
	}
	// set up references
	omd.handlerCategoryTable["cpu"] = omd.datapointHandlerHolder.GetCPUHandler()
	omd.handlerCategoryTable["network"] = omd.datapointHandlerHolder.GetNetworkHandler()
	omd.handlerCategoryTable["memory"] = omd.datapointHandlerHolder.GetMemoryHandler()
	omd.handlerCategoryTable["filesystem"] = omd.datapointHandlerHolder.GetFilesystemHandler()

	metricDecoder = omd
	return nil
}

func (omd *OakestraMetricDecoder) DecodeMetrics(rawMetrics []byte) error {
	md, err := unmarshaler.UnmarshalMetrics(rawMetrics)
	if err != nil {
		return fmt.Errorf("problem unmarshaling metrics -> dropping package")
	}
	for i := 0; i < md.ResourceMetrics().Len(); i++ {
		resourceMetrics := md.ResourceMetrics().At(i)
		for j := 0; j < resourceMetrics.ScopeMetrics().Len(); j++ {
			scopeMetrics := resourceMetrics.ScopeMetrics().At(j)
			for k := 0; k < scopeMetrics.Metrics().Len(); k++ {
				metric := scopeMetrics.Metrics().At(k)
				// handle metric through direct call of the appropriate handler's HandleMetrics function
				// the split up metrics.Name() usually returns an array of size 3
				datapointHandler := omd.getDatapointHandler(strings.Split(metric.Name(), ".")[1])
				if datapointHandler == nil {
					continue
				}
				datapointHandler.HandleMetric(metric)
			}
		}
	}
	return nil
}

func (omd *OakestraMetricDecoder) getDatapointHandler(datapointType string) handlers.DatapointHandler {
	h := omd.handlerCategoryTable[datapointType]
	if h == nil {
		fmt.Printf("datapoint handler not initialized or implemented for type: %s\n", datapointType)
	}
	return h
}
