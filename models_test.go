package gotoml

import "testing"

func TestTakeAction(t *testing.T) {
	tests := []struct {
		givenCell         Cell
		givenMode         ActionMode
		givenEpsilG       float64
		expectedAction    int
		expectedActionMax int
	}{
		{
			givenCell: Cell{
				Weight:  1,
				Actions: []float64{0.5, 0, 0.8, 0, 0.9, 0, 0, 0},
			},
			givenMode:         0,
			givenEpsilG:       1,
			expectedActionMax: 3,
		},
		{
			givenCell: Cell{
				Weight:  1,
				Actions: []float64{0.5, 0, 0.8, 0, 0.9, 0, 0, 0},
			},
			givenMode:         1,
			givenEpsilG:       1,
			expectedActionMax: 7,
		},
		{
			givenCell: Cell{
				Weight:  1,
				Actions: []float64{0.5, 0, 0.8, 0, 0.9, 0, 0, 0},
			},
			givenMode:      0,
			givenEpsilG:    0,
			expectedAction: 2,
		},
		{
			givenCell: Cell{
				Weight:  1,
				Actions: []float64{0.5, 0, 0.8, 0, 0.9, 0, 0, 0},
			},
			givenMode:      1,
			givenEpsilG:    0,
			expectedAction: 4,
		},
		{
			givenCell: Cell{
				Weight:  1,
				Actions: []float64{0, 0, 0, 0, 0, 0, 0, 0},
			},
			givenMode:      0,
			givenEpsilG:    0,
			expectedAction: 0,
		},
		{
			givenCell: Cell{
				Weight:  1,
				Actions: []float64{0, 0, 0, 0, 0, 0, 0, 0},
			},
			givenMode:      1,
			givenEpsilG:    0,
			expectedAction: 0,
		},
		{
			givenCell: Cell{
				Weight:  1,
				Actions: []float64{1, 1, 1, 1, 1, 1, 1, 1},
			},
			givenMode:      0,
			givenEpsilG:    0,
			expectedAction: 0,
		},
		{
			givenCell: Cell{
				Weight:  1,
				Actions: []float64{1, 1, 1, 1, 1, 1, 1, 1},
			},
			givenMode:      1,
			givenEpsilG:    0,
			expectedAction: 0,
		},
	}

	for i, chk := range tests {
		if i <= 1 {
			for j := 0; j < 100; j++ {
				gotAction := chk.givenCell.TakeAction(chk.givenEpsilG, chk.givenMode)
				if gotAction > chk.expectedActionMax {
					t.Errorf("given EpsilonG : [%v] \n expected max : [%v] got : [%v]", chk.givenEpsilG, chk.expectedActionMax, gotAction)
				}
			}

		} else {
			gotAction := chk.givenCell.TakeAction(chk.givenEpsilG, chk.givenMode)
			if gotAction != chk.expectedAction {
				t.Errorf("given: [%v] \n expected : [%v] got : [%v]", chk.givenCell.Actions, chk.expectedAction, gotAction)
			}
		}
	}
}
