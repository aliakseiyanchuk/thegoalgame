package main

import (
	"flag"
	"fmt"
	"github.com/aliakseiyanchuk/thegoalgame/lib"
	"os"
	"time"
)

var pipelineSize int
var wrcMin int
var wrcMax int
var matchBoxBehaviour string
var runCycles int
var simulationRuns int
var pauseBetweenSumulations int

var plotAchievedOutput bool
var plotLag bool
var plotAllWrcInventories bool
var plotAllWrcStarvedCapacity bool
var plotRelativeStarvedCapacity bool
var plotAll bool

const MatchBoxEpicAlternating = "epic-alternating"
const MatchBoxSimple = "simple"

func supplyInventory() lib.Inventory {
	if MatchBoxEpicAlternating == matchBoxBehaviour {
		return lib.CreateAlternatingBottomlessInventory(float64(wrcMin) + float64(wrcMax-wrcMin)/2)
	} else if matchBoxBehaviour != MatchBoxSimple {
		fmt.Printf("Unrecognized option for match box behaviour: %s", matchBoxBehaviour)
		fmt.Println("Supported options are:")

		mbOptions := []string{MatchBoxSimple, MatchBoxEpicAlternating}

		for _, v := range mbOptions {
			fmt.Printf(" - %s", v)
		}

		os.Exit(1)
	}

	return &lib.SimpleBottomlessInventory{}
}

func noSpecificOutputRequested() bool {
	return !plotAll &&
		!plotAchievedOutput &&
		!plotLag &&
		!plotAllWrcInventories &&
		!plotAllWrcStarvedCapacity &&
		!plotRelativeStarvedCapacity
}

func init() {
	flag.IntVar(&pipelineSize, "ps", 5, "Production pipeline size")
	flag.IntVar(&wrcMin, "wrc-min", 1, "Work center minimal capacity during a single run")
	flag.IntVar(&wrcMax, "wrc-max", 6, "Work center maximal capacity during a single run")
	flag.StringVar(&matchBoxBehaviour, "m", "simple", "Behaviour option for a \"match box\" when the first stage of pipeline tries to draw \"matches\"")

	flag.IntVar(&runCycles, "c", 20, "Number of cycles to perform")
	flag.IntVar(&simulationRuns, "R", 1, "Number of simulation runs to repeat")
	flag.IntVar(&pauseBetweenSumulations, "P", 5, "Pause between simulations")

	flag.BoolVar(&plotAchievedOutput, "plot-achieved-output", false, "Plot expected vs achieved output")
	flag.BoolVar(&plotLag, "plot-lag", false, "Plot output lag relative to baseline")
	flag.BoolVar(&plotAllWrcInventories, "plot-wrc-inventories", false, "Plot inventories work centers have accumulated")
	flag.BoolVar(&plotAllWrcStarvedCapacity, "plot-wrc-starving", false, "Plot unused capacity in work center")
	flag.BoolVar(&plotRelativeStarvedCapacity, "plot-starving", false, "Plot total unused capacity across all work center in a run")

	flag.BoolVar(&plotAll, "G", false, "Output all plots")
}

func main() {
	flag.Parse()

	if noSpecificOutputRequested() {
		fmt.Println("Warn: you haven't requested any specific output to be produced")
		fmt.Println("This run will print how inventory will evolve across work centers.")
		plotAllWrcInventories = true
	}

	for sc := 0; sc < simulationRuns; sc++ {
		pl := lib.CreateProductionLineWithSourceInventory(pipelineSize, supplyInventory())
		c := make([]lib.ProductionLineRun, runCycles)

		for i := 0; i < runCycles; i++ {
			c[i] = pl.Run(wrcMin, wrcMax)
		}

		g := lib.CreateGrapher(c, wrcMin, wrcMax)
		if plotAchievedOutput || plotAll {
			g.PlotAchievedOutput(sc)
		}
		if plotLag || plotAll {
			g.PlotLag(sc)
		}
		if plotAllWrcInventories || plotAll {
			g.PlotAllWorkCenterInventories(sc)
		}
		if plotAllWrcStarvedCapacity || plotAll {
			g.PlotAllWorkCenterUnusedCapacity(sc)
		}
		if plotRelativeStarvedCapacity || plotAll {
			g.PlotRelativeCumulativeUnusedCapacity(sc)
		}

		if sc < simulationRuns-1 {
			time.Sleep(time.Second * time.Duration(pauseBetweenSumulations))
		}
	}
}
