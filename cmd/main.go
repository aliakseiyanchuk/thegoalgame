package main

import (
	"flag"
	"fmt"
	"github.com/aliakseiyanchuk/thegoalgame/lib"
)

var pipelineSize int
var wrcMin int
var wrcMax int
var useAlternatingInventory bool
var runs int

var plotAchievedOutput bool
var plotLag bool
var plotAllWrcInventories bool
var plotAllWrcUnusedCapacity bool
var plotRelativeUnusedCapacity bool
var plotAll bool

func supplyInventory() lib.Inventory {
	if useAlternatingInventory {
		return lib.CreateAlternatingBottomlessInventory(wrcMin + (wrcMax-wrcMin)/2)
	} else {
		return &lib.SimpleBottomlessInventory{}
	}
}

func noSpecificOutputRequested() bool {
	return !plotAll &&
		!plotAchievedOutput &&
		!plotLag &&
		!plotAllWrcInventories &&
		!plotAllWrcUnusedCapacity &&
		!plotRelativeUnusedCapacity
}

func init() {
	flag.IntVar(&pipelineSize, "ps", 5, "Production pipeline size")
	flag.IntVar(&wrcMin, "wrc-min", 1, "Work center minimal capacity during a single run")
	flag.IntVar(&wrcMax, "wrc-max", 6, "Work center maximal capacity during a single run")
	flag.BoolVar(&useAlternatingInventory, "alternating-inventory", false, "Use epic/non-epic alternating inventory")
	flag.IntVar(&runs, "r", 20, "Number of runs to perform")

	flag.BoolVar(&plotAchievedOutput, "plot-achieved-output", false, "Plot expected vs achieved output")
	flag.BoolVar(&plotLag, "plot-lag", false, "Plot output lag relative to baseline")
	flag.BoolVar(&plotAllWrcInventories, "plot-wrc-inventories", false, "Plot inventories work centers have accumulated")
	flag.BoolVar(&plotAllWrcUnusedCapacity, "plot-wrc-unused-capacity", false, "Plot unused capacity in work center")
	flag.BoolVar(&plotRelativeUnusedCapacity, "plot-unused-capacity", false, "Plot total unused capacity across all work center in a run")

	flag.BoolVar(&plotAll, "G", false, "Output all plots")
}

func main() {
	flag.Parse()

	if noSpecificOutputRequested() {
		fmt.Println("Warn: you haven't requested any specific output to be produced")
		fmt.Println("This run will print how inventory will evolve across work centers.")
		plotAllWrcInventories = true
	}

	pl := lib.CreateProductionLineWithSourceInventory(pipelineSize, supplyInventory())
	c := make([]lib.ProductionLineRun, runs)

	for i := 0; i < runs; i++ {
		c[i] = pl.Run(wrcMin, wrcMax)
	}

	g := lib.CreateGrapher(c, wrcMin, wrcMax)
	if plotAchievedOutput || plotAll {
		g.PlotAchievedOutput()
	}
	if plotLag || plotAll {
		g.PlotLag()
	}
	if plotAllWrcInventories || plotAll {
		g.PlotAllWorkCenterInventories()
	}
	if plotAllWrcUnusedCapacity || plotAll {
		g.PlotAllWorkCenterUnusedCapacity()
	}
	if plotRelativeUnusedCapacity || plotAll {
		g.PlotRelativeCumulativeUnusedCapacity()
	}
}
