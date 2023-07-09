// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable_test

// Import the Go standard library package testing as well as tstable and tserr
import (
	"testing" // testing

	"github.com/thorstenrie/tserr"   // tserr
	"github.com/thorstenrie/tstable" // tstable
)

// TestGrids tests all grids of AllGrids. The test fails, if the retrieved string representation of the test table
// with the specified grid does not equal to the contents of the test data golden file.
func TestGrids(t *testing.T) {
	// retrieve test table
	tbl := testTable(t)
	// Iterate over all grids of AllGrids
	for name, grid := range tstable.AllGrids {
		// Set the grid of the test table
		e := tbl.SetGrid(grid)
		// The test fails if SetGrid fails
		if e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{Op: "SetGrid", Fn: name, Err: e}))
		}
		// Evaluate the test table
		evalTable(name, tbl, t)
	}
}
