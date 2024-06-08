package lib

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
)

var allColours = [...]asciigraph.AnsiColor{
	asciigraph.AntiqueWhite,
	asciigraph.Aqua,
	asciigraph.Aquamarine,
	asciigraph.Beige,
	asciigraph.Bisque,
	asciigraph.Blue,
	asciigraph.BlueViolet,
	asciigraph.Brown,
	asciigraph.BurlyWood,
	asciigraph.CadetBlue,
	asciigraph.Chartreuse,
	asciigraph.Chocolate,
	asciigraph.Coral,
	asciigraph.CornflowerBlue,
	asciigraph.Cornsilk,
	asciigraph.Crimson,
	asciigraph.Cyan,
	asciigraph.DarkBlue,
	asciigraph.DarkBlue,
	// Do we have enough?
}

type Grapher struct {
	runs        []ProductionLineRun
	minCapacity int
	maxCapacity int
}

func CreateGrapher(runs []ProductionLineRun, minCapacity, maxCapacity int) *Grapher {
	return &Grapher{
		runs:        runs,
		minCapacity: minCapacity,
		maxCapacity: maxCapacity,
	}
}

func (g *Grapher) PlotAchievedOutput() {
	data := make([][]float64, 2)

	avg := float64(g.minCapacity) + float64(g.maxCapacity-g.minCapacity)/2

	numRuns := len(g.runs)
	expectedBand := make([]float64, numRuns)
	actualBand := make([]float64, numRuns)

	for i := range g.runs {
		expectedBand[i] = float64(i+1) * avg
		actualBand[i] = float64(g.runs[i].AchievedOutput)
	}

	data[0] = expectedBand
	data[1] = actualBand

	graph := asciigraph.PlotMany(data,
		asciigraph.Precision(0),
		asciigraph.Width(60),
		asciigraph.Height(15),
		asciigraph.SeriesColors(asciigraph.Blue, asciigraph.Pink),
		asciigraph.SeriesLegends("Expected", "Achieved"),
		asciigraph.Caption(fmt.Sprintf("Achieved %d of %d mean", int(actualBand[numRuns-1]), int(expectedBand[numRuns-1]))),
	)

	fmt.Println(graph)
}

func (g *Grapher) PlotLag() {
	avg := float64(g.minCapacity) + float64(g.maxCapacity-g.minCapacity)/2

	lag := make([]float64, len(g.runs))

	for i := range g.runs {
		lag[i] = avg*float64(i) - float64(g.runs[i].AchievedOutput)
	}

	graph := asciigraph.Plot(lag,
		asciigraph.Precision(0),
		asciigraph.Width(60),
		asciigraph.Height(15),
		asciigraph.SeriesColors(asciigraph.Pink),
		asciigraph.SeriesLegends("Output Lag"),
		asciigraph.Caption("Output Lag"),
	)

	fmt.Println(graph)
}

func (g *Grapher) PlotAllWorkCenterInventories() {

	numWorkCenters := len(g.runs[0].WorkCenterStat)
	inventoryData := make([][]float64, numWorkCenters)
	for i := range inventoryData {
		inventoryData[i] = make([]float64, len(g.runs))
	}

	for runIndex := range g.runs {
		for wrcIndex := 0; wrcIndex < numWorkCenters; wrcIndex++ {
			inventoryData[wrcIndex][runIndex] = float64(g.runs[runIndex].WorkCenterStat[wrcIndex].InventorySizeAfterRun)
		}
	}

	wrcLabels := make([]string, numWorkCenters)
	for i := range wrcLabels {
		wrcLabels[i] = fmt.Sprintf("Work Center %d", i+1)
	}

	graph := asciigraph.PlotMany(inventoryData,
		asciigraph.Precision(0),
		asciigraph.Width(60),
		asciigraph.Height(15),
		asciigraph.SeriesColors(allColours[:numWorkCenters]...),
		asciigraph.SeriesLegends(wrcLabels...),
		asciigraph.Caption("Work Center Inventory"),
	)

	fmt.Println(graph)
}

func (g *Grapher) PlotAllWorkCenterUnusedCapacity() {

	numWorkCenters := len(g.runs[0].WorkCenterStat)
	capacityData := make([][]float64, numWorkCenters)
	for i := range capacityData {
		capacityData[i] = make([]float64, len(g.runs))
	}

	for runIndex := range g.runs {
		for wrcIndex := 0; wrcIndex < numWorkCenters; wrcIndex++ {
			capacityData[wrcIndex][runIndex] = float64(g.runs[runIndex].WorkCenterStat[wrcIndex].UnusedRunCapacity)
		}
	}

	wrcLabels := make([]string, numWorkCenters)
	for i := range wrcLabels {
		wrcLabels[i] = fmt.Sprintf("Work Center %d", i+1)
	}

	graph := asciigraph.PlotMany(capacityData,
		asciigraph.Precision(0),
		asciigraph.Width(60),
		asciigraph.Height(15),
		asciigraph.SeriesColors(allColours[:numWorkCenters]...),
		asciigraph.SeriesLegends(wrcLabels...),
		asciigraph.Caption("Work Center Starving"),
	)

	fmt.Println(graph)
}

func (g *Grapher) PlotCumulativeUnusedCapacity() {

	numWorkCenters := len(g.runs[0].WorkCenterStat)
	cumul := [][]float64{
		make([]float64, len(g.runs)),
		make([]float64, len(g.runs)),
		make([]float64, len(g.runs)),
	}

	for runIndex := range g.runs {
		runUnused := float64(0)
		runAvailable := float64(0)
		for wrcIndex := 0; wrcIndex < numWorkCenters; wrcIndex++ {
			runAvailable += float64(g.runs[runIndex].WorkCenterStat[wrcIndex].RunCapacity)
			runUnused += float64(g.runs[runIndex].WorkCenterStat[wrcIndex].UnusedRunCapacity)
		}

		cumul[0][runIndex] = runAvailable
		cumul[1][runIndex] = runUnused
		cumul[2][runIndex] = runUnused / runAvailable
	}

	graph := asciigraph.PlotMany(cumul[0:2],
		asciigraph.Precision(0),
		asciigraph.Width(60),
		asciigraph.Height(15),
		asciigraph.SeriesColors(asciigraph.Blue, asciigraph.Red),
		asciigraph.SeriesLegends("Capacity Available", "Capacity Unused"),
		asciigraph.Caption("Total Starving Per Run"),
	)

	fmt.Println(graph)
}
func (g *Grapher) PlotRelativeCumulativeUnusedCapacity() {

	numWorkCenters := len(g.runs[0].WorkCenterStat)
	cumul := [][]float64{
		make([]float64, len(g.runs)),
		make([]float64, len(g.runs)),
		make([]float64, len(g.runs)),
	}

	for runIndex := range g.runs {
		runUnused := float64(0)
		runAvailable := float64(0)
		for wrcIndex := 0; wrcIndex < numWorkCenters; wrcIndex++ {
			runAvailable += float64(g.runs[runIndex].WorkCenterStat[wrcIndex].RunCapacity)
			runUnused += float64(g.runs[runIndex].WorkCenterStat[wrcIndex].UnusedRunCapacity)
		}

		cumul[0][runIndex] = runAvailable
		cumul[1][runIndex] = runUnused
		cumul[2][runIndex] = runUnused / runAvailable
	}

	data := cumul[2]
	for i := range data {
		data[i] = data[i] * float64(100)
	}

	graph := asciigraph.Plot(data,
		asciigraph.Precision(2),
		asciigraph.Width(60),
		asciigraph.Height(15),
		asciigraph.SeriesColors(asciigraph.Red),
		asciigraph.SeriesLegends("Unused capacity, %"),
		asciigraph.Caption("Unused Capacity %"),
	)

	fmt.Println(graph)
}
