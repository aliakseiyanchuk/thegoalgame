package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInventorySelect_WithEpic(t *testing.T) {
	inv := WorkCenterInventory{}

	inv.AddAll([]WorkUnit{
		NewWorkUnit(3),
		NewWorkUnit(1),
	},
	)

	remaining, selected := inv.Select(2)
	assert.Equal(t, 1, len(remaining))
	assert.Equal(t, 3, remaining[0].Size)

	assert.Equal(t, 1, len(selected))
	assert.Equal(t, 1, selected[0].Size)
}

func TestInventorySelect_SelectingEpic(t *testing.T) {
	inv := WorkCenterInventory{}

	inv.AddAll([]WorkUnit{
		NewWorkUnit(3),
		NewWorkUnit(1),
		NewWorkUnit(1),
	},
	)

	remaining, selected := inv.Select(4)
	assert.Equal(t, 2, len(selected))
	assert.Equal(t, 3, selected[0].Size)
	assert.Equal(t, 1, selected[1].Size)

	assert.Equal(t, 1, len(remaining))
	assert.Equal(t, 1, remaining[0].Size)
}
