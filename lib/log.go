package lib

type WorkCenterRun struct {
	WorkCenterIndex        int
	InventorySizeBeforeRun int
	InventorySizeAfterRun  int
	InventoryAtStart       []WorkUnit
	RunCapacity            int

	PassOnInventory    []WorkUnit
	RemainingInventory []WorkUnit
	UnusedRunCapacity  int
}

type ProductionLineRun struct {
	AchievedOutput int
	WorkCenterStat []WorkCenterRun
}

func CreateProductionRunLog(workCenters int) ProductionLineRun {
	return ProductionLineRun{
		AchievedOutput: 0,
		WorkCenterStat: make([]WorkCenterRun, workCenters),
	}
}

func (wcr *WorkCenterRun) InventoryAtEnd() int {
	rv := 0

	for _, ri := range wcr.RemainingInventory {
		rv = rv + ri.Size
	}

	return rv
}

func (wcr *WorkCenterRun) UnusedCapacity() int {
	return max(0, wcr.RunCapacity-wcr.InventorySizeBeforeRun)
}
