package handlers

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type DatapointHandler interface {
	HandleMetric(metric pmetric.Metric)
	getMetricTypeHandle(string) func(metric pmetric.Metric)
	handleGauge(metric pmetric.Metric)
	handleSum(metric pmetric.Metric)
	handleHistogram(metric pmetric.Metric)
	handleSummary(metric pmetric.Metric)
}

type DatapointHandlerHolder interface {
	GetCPUHandler() DatapointHandler
	GetNetworkHandler() DatapointHandler
	GetDiskIOHandler() DatapointHandler
	GetMemoryHandler() DatapointHandler
	GetFilesystemHandler() DatapointHandler
}

var _ DatapointHandlerHolder = (*OakestraDatapointHandlerHolder)(nil)

var datapointHandlerHolder DatapointHandlerHolder

func GetDatapointHandlerHolder() DatapointHandlerHolder {
	return datapointHandlerHolder
}

type OakestraDatapointHandlerHolder struct {
	cpuDatapointHandler        DatapointHandler
	networkDatapointHandler    DatapointHandler
	diskIODatapointHandler     DatapointHandler
	memoryDatapointHandler     DatapointHandler
	filesystemDatapointHandler DatapointHandler
}

func InitOakestraDatapointHandlerHolder() error {
	odhh := &OakestraDatapointHandlerHolder{
		cpuDatapointHandler:        newCPUDatapointHandler(),
		networkDatapointHandler:    newNoopDatapointHandler(),
		diskIODatapointHandler:     newNoopDatapointHandler(),
		memoryDatapointHandler:     newNoopDatapointHandler(),
		filesystemDatapointHandler: newNoopDatapointHandler(),
	}

	datapointHandlerHolder = odhh
	return nil
}

func (odhh *OakestraDatapointHandlerHolder) GetCPUHandler() DatapointHandler {
	return odhh.cpuDatapointHandler
}

func (odhh *OakestraDatapointHandlerHolder) GetNetworkHandler() DatapointHandler {
	return odhh.networkDatapointHandler
}

func (odhh *OakestraDatapointHandlerHolder) GetDiskIOHandler() DatapointHandler {
	return odhh.diskIODatapointHandler
}

func (odhh *OakestraDatapointHandlerHolder) GetMemoryHandler() DatapointHandler {
	return odhh.memoryDatapointHandler
}

func (odhh *OakestraDatapointHandlerHolder) GetFilesystemHandler() DatapointHandler {
	return odhh.filesystemDatapointHandler
}
