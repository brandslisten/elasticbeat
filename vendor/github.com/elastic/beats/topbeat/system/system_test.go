// +build !integration

package system

import (
	"runtime"
	"testing"

	"github.com/elastic/gosigar"
	"github.com/stretchr/testify/assert"
)

func TestGetSystemLoad(t *testing.T) {

	if runtime.GOOS == "windows" {
		return //no load data on windows
	}

	load, err := GetSystemLoad()

	assert.NotNil(t, load)
	assert.Nil(t, err)

	assert.True(t, (load.Load1 > 0))
	assert.True(t, (load.Load5 > 0))
	assert.True(t, (load.Load15 > 0))
}

func TestGetMemory(t *testing.T) {
	mem, err := GetMemory()

	assert.NotNil(t, mem)
	assert.Nil(t, err)

	assert.True(t, (mem.Total > 0))
	assert.True(t, (mem.Used > 0))
	assert.True(t, (mem.Free >= 0))
	assert.True(t, (mem.ActualFree >= 0))
	assert.True(t, (mem.ActualUsed > 0))
}

func TestGetSwap(t *testing.T) {

	if runtime.GOOS == "windows" {
		return //no load data on windows
	}

	swap, err := GetSwap()

	assert.NotNil(t, swap)
	assert.Nil(t, err)

	assert.True(t, (swap.Total >= 0))
	assert.True(t, (swap.Used >= 0))
	assert.True(t, (swap.Free >= 0))
}

func TestMemPercentage(t *testing.T) {

	m := MemStat{
		Mem: gosigar.Mem{
			Total: 7,
			Used:  5,
			Free:  2,
		},
	}
	AddMemPercentage(&m)
	assert.Equal(t, m.UsedPercent, 0.71)

	m = MemStat{
		Mem: gosigar.Mem{Total: 0},
	}
	AddMemPercentage(&m)
	assert.Equal(t, m.UsedPercent, 0.0)
}

func TestActualMemPercentage(t *testing.T) {

	m := MemStat{
		Mem: gosigar.Mem{
			Total:      7,
			ActualUsed: 5,
			ActualFree: 2,
		},
	}
	AddMemPercentage(&m)
	assert.Equal(t, m.ActualUsedPercent, 0.71)

	m = MemStat{
		Mem: gosigar.Mem{
			Total: 0,
		},
	}
	AddMemPercentage(&m)
	assert.Equal(t, m.ActualUsedPercent, 0.0)
}
