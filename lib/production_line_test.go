package lib

import (
	"testing"
)

func TestCreateProductionLine(t *testing.T) {
	pl := CreateProductionLine(5)

	numRuns := 20
	minWrcCapacity := 1
	maxWrcCapacity := 6
	c := make([]ProductionLineRun, numRuns)

	for i := 0; i < numRuns; i++ {
		c[i] = pl.Run(minWrcCapacity, maxWrcCapacity)
	}

	g := CreateGrapher(c, minWrcCapacity, maxWrcCapacity)
	g.PlotAchievedOutput()
	g.PlotLag()
	g.PlotAllWorkCenterInventories()
	g.PlotAllWorkCenterUnusedCapacity()
	g.PlotCumulativeUnusedCapacity()
	g.PlotRelativeCumulativeUnusedCapacity()
}

func TestCreateProductionLineWithAlternatingEpic(t *testing.T) {
	pl := CreateProductionLineWithSourceInventory(5, &AlternatingEpicBottomlessInventory{epicSize: 3})

	numRuns := 20
	minWrcCapacity := 1
	maxWrcCapacity := 6

	c := make([]ProductionLineRun, numRuns)

	for i := 0; i < numRuns; i++ {
		c[i] = pl.Run(minWrcCapacity, maxWrcCapacity)
	}

	g := CreateGrapher(c, minWrcCapacity, maxWrcCapacity)
	g.PlotAchievedOutput()
	g.PlotLag()
	g.PlotAllWorkCenterInventories()
	g.PlotAllWorkCenterUnusedCapacity()
	g.PlotCumulativeUnusedCapacity()
	g.PlotRelativeCumulativeUnusedCapacity()
}
