package gotoml

import "testing"

func TestTakeAction(t *testing.T) {
	tests := []struct {
		givenCell         Cell
		givenEpsilG       float64
		expectedAction    int
		expectedActionMax int
	}{
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{0.5, 0, 0.8, 0},
			},
			givenEpsilG:       1,
			expectedActionMax: 3,
		},
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{0.5, 0, 0.8, 0},
			},
			givenEpsilG:       0.5,
			expectedActionMax: 3,
		},
		//greedy only
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{0, 0, 0, 0},
			},
			givenEpsilG:    0,
			expectedAction: 0,
		},
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{-1, 0, 0, 0},
			},
			givenEpsilG:    0,
			expectedAction: 1,
		},
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{-1, -1, -1, -1},
			},
			givenEpsilG:    0,
			expectedAction: 0,
		},
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{0.9, 0, 0, 0},
			},
			givenEpsilG:    0,
			expectedAction: 0,
		},
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{0.5, 0.9, 0.8, 0},
			},
			givenEpsilG:    0,
			expectedAction: 1,
		},
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{0.5, 0, 0.8, 0},
			},
			givenEpsilG:    0,
			expectedAction: 2,
		},
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{0.5, 0, 0.8, 0.9},
			},
			givenEpsilG:    0,
			expectedAction: 3,
		},
		{
			givenCell: &Cell4Dir{
				Weight:  1,
				Actions: []float64{1, 1, 1, 1},
			},
			givenEpsilG:    0,
			expectedAction: 0,
		},
	}

	for i, chk := range tests {
		if i <= 1 {
			// random
			for j := 0; j < 100; j++ {
				gotAction := chk.givenCell.TakeAction(chk.givenEpsilG)
				if gotAction > chk.expectedActionMax {
					t.Errorf("given EpsilonG : [%v] \n expected max : [%v] got : [%v]", chk.givenEpsilG, chk.expectedActionMax, gotAction)
				}
			}

		} else {
			// greedy
			gotAction := chk.givenCell.TakeAction(chk.givenEpsilG)
			if gotAction != chk.expectedAction {
				t.Errorf("given: [%v] \n expected : [%v] got : [%v]", chk.givenCell.GetActions(), chk.expectedAction, gotAction)
			}
		}
	}
}

func BenchmarkTakeActionExploreFull(b *testing.B) {
	var CellTest Cell
	CellTest = &Cell4Dir{Actions: []float64{0.1, 0.2, 1, -1}}
	for i := 0; i < b.N; i++ {
		CellTest.TakeAction(1)
	}
}

func BenchmarkTakeActionExploreMid(b *testing.B) {
	var CellTest Cell
	CellTest = &Cell4Dir{Actions: []float64{0.1, 0.2, 1, -1}}
	for i := 0; i < b.N; i++ {
		CellTest.TakeAction(0.5)
	}
}

func BenchmarkTakeActionExploit(b *testing.B) {
	var CellTest Cell
	CellTest = &Cell4Dir{Actions: []float64{0.1, 0.2, 1, -1}}
	for i := 0; i < b.N; i++ {
		CellTest.TakeAction(0)
	}
}

func TestGrid(t *testing.T) {
	tests := []struct {
		givenDim         []int
		givenInterf      Cell
		givenCellCoord   []int
		givenCell        Cell
		expectedErr      error
		expectedCellBool bool
		expectedCellVal  Cell
	}{
		{
			givenDim:         []int{10, 3},
			givenInterf:      Cell4Dir{},
			givenCellCoord:   []int{0, 0},
			givenCell:        Cell4Dir{Actions: []float64{1, 1, 1, 1}},
			expectedErr:      nil,
			expectedCellBool: true,
			expectedCellVal:  Cell4Dir{Actions: []float64{-1, 1, 1, -1}},
		},
		{
			givenDim:         []int{10, 3},
			givenInterf:      Cell4Dir{},
			givenCellCoord:   []int{10, 0},
			givenCell:        Cell4Dir{Actions: []float64{1, 1, 1, 1}},
			expectedErr:      nil,
			expectedCellBool: false,
		},
		{
			givenDim:    []int{0, 3},
			givenInterf: Cell4Dir{},
			expectedErr: ErrGridWrongDimensions,
		},
		{
			givenDim:    []int{10, -3},
			givenInterf: Cell4Dir{},
			expectedErr: ErrGridWrongDimensions,
		},
	}

	for _, chk := range tests {
		gridTest, gotErr := NewGrid(chk.givenDim[0], chk.givenDim[1], chk.givenInterf)
		if chk.expectedErr != gotErr {
			t.Errorf("given %v expected %v got %v", chk.givenDim, chk.expectedErr, gotErr)
		}
		if gotErr == nil {
			if gridTest.Height != chk.givenDim[0] {
				t.Errorf("height given %v got grid height %v", chk.givenDim[0], gridTest.Height)
			}
			if gridTest.Width != chk.givenDim[1] {
				t.Errorf("height given %v got grid width %v", chk.givenDim[1], gridTest.Width)
			}

			gotBool := gridTest.PutCell(chk.givenCellCoord[0], chk.givenCellCoord[1], chk.givenCell)
			if chk.expectedCellBool != gotBool {
				t.Errorf("given Cell  %v in %v got %v expected %v", chk.givenCell, chk.givenDim, gotBool, chk.expectedCellBool)
			}
			if gotBool && chk.expectedCellBool == gotBool {
				gotCell, _ := gridTest.GetCell(chk.givenCellCoord[0], chk.givenCellCoord[1])
				if chk.expectedCellVal.GetAction(0) != gotCell.GetAction(0) {
					t.Errorf("expected Cell  %v got %v ", chk.expectedCellVal, gotCell)
				}
				if chk.expectedCellVal.GetAction(1) != gotCell.GetAction(1) {
					t.Errorf("expected Cell  %v got %v ", chk.expectedCellVal, gotCell)
				}
				if chk.expectedCellVal.GetAction(2) != gotCell.GetAction(2) {
					t.Errorf("expected Cell  %v got %v ", chk.expectedCellVal, gotCell)
				}
				if chk.expectedCellVal.GetAction(3) != gotCell.GetAction(3) {
					t.Errorf("expected Cell  %v got %v ", chk.expectedCellVal, gotCell)
				}
			}
		}
	}
}
