package lib

import (
	"math/rand"
)

type WorkUnit struct {
	Size             int
	RequiredCapacity int
}

func (wu WorkUnit) Clone() WorkUnit {
	return WorkUnit{
		Size:             wu.Size,
		RequiredCapacity: wu.RequiredCapacity,
	}
}

func NewWorkUnit(size int) WorkUnit {
	return WorkUnit{size, size}
}

func (wu WorkUnit) CanSelect(units int) bool {
	return wu.RequiredCapacity <= units
}

type WorkCenter struct {
	Index     int
	Inventory Inventory
	Output    InventoryReceiver
}

func randomValueFrom(from, to int) int {
	return from + rand.Intn(to-from+1)
}

func (wc *WorkCenter) Run(from, to int) WorkCenterRun {
	runCapacity := randomValueFrom(from, to)
	toRemain, toPass := wc.Inventory.Select(runCapacity)

	rv := WorkCenterRun{
		WorkCenterIndex:        wc.Index,
		InventoryAtStart:       wc.Inventory.Get(),
		InventorySizeBeforeRun: wc.Inventory.TotalSize(),

		RunCapacity: runCapacity,

		RemainingInventory: toRemain,
		PassOnInventory:    toPass,

		InventorySizeAfterRun: WorkUnitsSize(toRemain...),
		UnusedRunCapacity:     runCapacity - WorkUnitsSize(toPass...),
	}

	return rv
}

type OutputCollector struct {
	units []WorkUnit
}

func (o *OutputCollector) Add(unit WorkUnit) {
	o.units = append(o.units, unit)
}

func (o *OutputCollector) AddAll(units []WorkUnit) {
	o.units = append(o.units, units...)
}

func (o *OutputCollector) CollectedSize() int {
	rv := 0
	for _, unit := range o.units {
		rv += unit.Size
	}

	return rv
}
