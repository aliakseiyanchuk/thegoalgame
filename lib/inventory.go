package lib

import (
	"cmp"
	"math"
	"slices"
)

type InventoryReceiver interface {
	Add(unit WorkUnit)
	AddAll(units []WorkUnit)
}

type Inventory interface {
	InventoryReceiver
	Select(size int) ([]WorkUnit, []WorkUnit)
	TotalSize() int
	Set([]WorkUnit)
	Get() []WorkUnit
}

func WorkUnitsSize(units ...WorkUnit) int {
	rv := 0
	for _, unit := range units {
		rv += unit.Size
	}

	return rv
}

type WorkCenterInventory struct {
	units []WorkUnit
}

func (inv *WorkCenterInventory) Get() []WorkUnit {
	rv := make([]WorkUnit, len(inv.units))
	for i := range inv.units {
		rv[i] = inv.units[i].Clone()
	}

	return rv
}

func (inv *WorkCenterInventory) Add(unit WorkUnit) {
	rv := append(inv.units, unit)
	slices.SortFunc(rv, func(a, b WorkUnit) int {
		return cmp.Compare(a.Size, b.Size)
	})

	inv.units = rv
}

func (inv *WorkCenterInventory) AddAll(units []WorkUnit) {
	rv := append(inv.units, units...)
	slices.SortFunc(rv, func(a, b WorkUnit) int {
		return cmp.Compare(b.Size, a.Size)
	})

	inv.units = rv
}

func (inv *WorkCenterInventory) Set(units []WorkUnit) {
	inv.units = units
}

func (inv *WorkCenterInventory) Select(size int) ([]WorkUnit, []WorkUnit) {
	var returned []WorkUnit
	var remaining []WorkUnit

	capacityRemaining := size
	for i := 0; i < len(inv.units); i++ {
		unit := inv.units[i]

		if capacityRemaining >= unit.RequiredCapacity {
			capacityRemaining = capacityRemaining - unit.Size
			returned = append(returned, unit)
		} else {
			remaining = append(remaining, unit)
		}
	}

	return remaining, returned
}

func (inv *WorkCenterInventory) TotalSize() int {
	return WorkUnitsSize(inv.units...)
}

//------------------------------------------------------------------------------------
// Initial inventory base

type InitialInventoryBase struct {
}

func (b InitialInventoryBase) Get() []WorkUnit {
	return []WorkUnit{}
}

func (b InitialInventoryBase) Add(unit WorkUnit) {
	//Nothing to do
}

func (b InitialInventoryBase) AddAll(units []WorkUnit) {
	// Nothing to do
}

func (b InitialInventoryBase) Set(units []WorkUnit) {
	// Do nothing.
}

func (b InitialInventoryBase) TotalSize() int {
	return math.MaxInt
}

//------------------------------------------------------------------------------------
// Simple bottomless inventory

type SimpleBottomlessInventory struct {
	InitialInventoryBase
}

func (b SimpleBottomlessInventory) Select(size int) ([]WorkUnit, []WorkUnit) {
	rv := make([]WorkUnit, size)

	for i := 0; i < size; i++ {
		rv[i].Size = 1
		rv[i].RequiredCapacity = 1
	}

	return []WorkUnit{}, rv
}

func CreateAlternatingBottomlessInventory(meanSize float64) Inventory {
	return &AlternatingEpicBottomlessInventory{
		meanSize:        meanSize,
		epicEmittedLast: false,
	}
}

type AlternatingEpicBottomlessInventory struct {
	InitialInventoryBase
	meanSize        float64
	epicEmittedLast bool
}

func (b *AlternatingEpicBottomlessInventory) Select(_ int) ([]WorkUnit, []WorkUnit) {
	var passing []WorkUnit

	if !b.epicEmittedLast {
		epicSize := int(math.Floor(b.meanSize))

		seedingEpic := []WorkUnit{
			{
				Size:             epicSize,
				RequiredCapacity: epicSize,
			},
		}

		passing = seedingEpic
		b.epicEmittedLast = true
	} else {
		passing = make([]WorkUnit, int(math.Ceil(b.meanSize)))
		for i := range passing {
			passing[i].Size = 1
			passing[i].RequiredCapacity = 1
		}
		b.epicEmittedLast = false
	}

	return []WorkUnit{}, passing
}
