package lib

type ProductionLine struct {
	line   []*WorkCenter
	output *OutputCollector
}

func (pl *ProductionLine) WorkCenterCount() int {
	return len(pl.line)
}

func (pl *ProductionLine) Run(from, to int) ProductionLineRun {
	plSize := len(pl.line)

	rv := CreateProductionRunLog(plSize)

	for i := 0; i < plSize; i++ {
		wrc := pl.line[i]
		wrcRun := wrc.Run(from, to)

		rv.WorkCenterStat[i] = wrcRun

		wrc.Inventory.Set(wrcRun.RemainingInventory)
		wrc.Output.AddAll(wrcRun.PassOnInventory)
	}

	rv.AchievedOutput = pl.output.CollectedSize()

	return rv
}

func CreateProductionLine(size int) *ProductionLine {
	rv := ProductionLine{}

	first := WorkCenter{
		Inventory: &SimpleBottomlessInventory{},
		Index:     0,
	}

	rv.line = make([]*WorkCenter, size)
	rv.line[0] = &first

	for i := 1; i < size; i++ {
		rv.line[i] = &WorkCenter{
			Inventory: &WorkCenterInventory{},
			Index:     i,
		}
	}

	for i := 0; i < size-1; i++ {
		rv.line[i].Output = rv.line[i+1].Inventory
	}

	rv.output = &OutputCollector{}
	rv.line[size-1].Output = rv.output

	return &rv
}
