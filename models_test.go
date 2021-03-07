package gotoml

import "testing"

func TestTakeAction(t *testing.T) {
	tests := []struct {
		givenNode         Noder
		givenEpsilG       float64
		expectedAction    int
		expectedActionMax int
	}{
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv5{0, 0.5, 0, 0.8, 0},
			},
			givenEpsilG:       1,
			expectedActionMax: 4,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv5{0, 0.5, 0, 0.8, 0},
			},
			givenEpsilG:       0.5,
			expectedActionMax: 4,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv4{0.5, 0, 0.8, 0},
			},
			givenEpsilG:       1,
			expectedActionMax: 3,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv4{0.5, 0, 0.8, 0},
			},
			givenEpsilG:       0.5,
			expectedActionMax: 3,
		},
		//greedy only
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv5{0, 0, 0, 0, 0},
			},
			givenEpsilG:    0,
			expectedAction: 0,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv5{-1, -1, 0, 0, 0},
			},
			givenEpsilG:    0,
			expectedAction: 2,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv5{-1, -1, -1, -1, -1},
			},
			givenEpsilG:    0,
			expectedAction: 0,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv5{0.9, 0, 0, 0, 0},
			},
			givenEpsilG:    0,
			expectedAction: 0,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv5{0, 0.5, 0.9, 0.8, 0},
			},
			givenEpsilG:    0,
			expectedAction: 2,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv5{0, 0.5, 0, 0.8, 0},
			},
			givenEpsilG:    0,
			expectedAction: 3,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv5{0, 0.5, 0, 0.8, 0.9},
			},
			givenEpsilG:    0,
			expectedAction: 4,
		},
		{
			givenNode: &Node{
				Reward:    1,
				ActionsQv: &ActionsQv4{1, 1, 1, 1},
			},
			givenEpsilG:    0,
			expectedAction: 0,
		},
	}

	for i, chk := range tests {
		if i <= 3 {
			// random
			for j := 0; j < 100; j++ {
				gotAction := chk.givenNode.TakeAction(chk.givenEpsilG)
				if gotAction > chk.expectedActionMax {
					t.Errorf("given EpsilonG : [%v] \n expected max : [%v] got : [%v]", chk.givenEpsilG, chk.expectedActionMax, gotAction)
				}
			}

		} else {
			// greedy
			gotAction := chk.givenNode.TakeAction(chk.givenEpsilG)
			if gotAction != chk.expectedAction {
				t.Errorf("given: [%v] \n expected : [%v] got : [%v]", chk.givenNode.GetActionsQv(), chk.expectedAction, gotAction)
			}
		}
	}
}

func BenchmarkTakeActionExploreFull(b *testing.B) {
	var NodeTest Noder
	NodeTest = &Node{ActionsQv: &ActionsQv5{0, 0.1, 0.2, 1, -1}}
	for i := 0; i < b.N; i++ {
		NodeTest.TakeAction(1)
	}
}

func BenchmarkTakeActionExploreMid(b *testing.B) {
	var NodeTest Noder
	NodeTest = &Node{ActionsQv: &ActionsQv5{0, 0.1, 0.2, 1, -1}}
	for i := 0; i < b.N; i++ {
		NodeTest.TakeAction(0.5)
	}
}

func BenchmarkTakeActionExploit(b *testing.B) {
	var NodeTest Noder
	NodeTest = &Node{ActionsQv: &ActionsQv5{0, 0.1, 0.2, 1, -1}}
	for i := 0; i < b.N; i++ {
		NodeTest.TakeAction(0)
	}
}

func TestGrid(t *testing.T) {
	tests := []struct {
		givenDim         []int
		givenNodeCoord   []int
		givenNode        Noder
		expectedErr      error
		expectedNodeBool bool
		expectedNodeVal  Noder
	}{
		{
			givenDim:         []int{10, 3},
			givenNodeCoord:   []int{0, 0},
			givenNode:        &Node{ActionsQv: &ActionsQv5{0, 1, 1, 1, 1}},
			expectedErr:      nil,
			expectedNodeBool: true,
			expectedNodeVal:  &Node{ActionsQv: &ActionsQv5{0, -1, 1, 1, -1}},
		},
		{
			givenDim:         []int{10, 3},
			givenNodeCoord:   []int{10, 0},
			givenNode:        &Node{ActionsQv: &ActionsQv5{0, 1, 1, 1, 1}},
			expectedErr:      nil,
			expectedNodeBool: false,
		},
		{
			givenDim:         []int{10, 3},
			givenNodeCoord:   []int{0, 0},
			givenNode:        &Node{ActionsQv: &ActionsQv4{1, 1, 1, 1}},
			expectedErr:      nil,
			expectedNodeBool: true,
			expectedNodeVal:  &Node{ActionsQv: &ActionsQv4{-1, 1, 1, -1}},
		},
		{
			givenDim:         []int{10, 3},
			givenNodeCoord:   []int{10, 0},
			givenNode:        &Node{ActionsQv: &ActionsQv4{1, 1, 1, 1}},
			expectedErr:      nil,
			expectedNodeBool: false,
		},
		{
			givenDim:    []int{0, 3},
			expectedErr: ErrGridWrongDimensions,
		},
		{
			givenDim:    []int{10, -3},
			expectedErr: ErrGridWrongDimensions,
		},
	}

	for _, chk := range tests {
		gridTest, gotErr := NewGrid(chk.givenDim[0], chk.givenDim[1], &Node{})
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

			gotBool := gridTest.PutNode(chk.givenNodeCoord[0], chk.givenNodeCoord[1], chk.givenNode)
			if chk.expectedNodeBool != gotBool {
				t.Errorf("given Noder  %v in %v got %v expected %v", chk.givenNode, chk.givenDim, gotBool, chk.expectedNodeBool)
			}
			if gotBool && chk.expectedNodeBool == gotBool {
				gotNode := gridTest.GetNode(chk.givenNodeCoord[0], chk.givenNodeCoord[1])
				if chk.expectedNodeVal.GetActionQv(0) != gotNode.GetActionQv(0) {
					t.Errorf("expected Noder  %v got %v ", chk.expectedNodeVal, gotNode)
				}
				if chk.expectedNodeVal.GetActionQv(1) != gotNode.GetActionQv(1) {
					t.Errorf("expected Noder  %v got %v ", chk.expectedNodeVal, gotNode)
				}
				if chk.expectedNodeVal.GetActionQv(2) != gotNode.GetActionQv(2) {
					t.Errorf("expected Noder  %v got %v ", chk.expectedNodeVal, gotNode)
				}
				if chk.expectedNodeVal.GetActionQv(3) != gotNode.GetActionQv(3) {
					t.Errorf("expected Noder  %v got %v ", chk.expectedNodeVal, gotNode)
				}
			}
		}
	}
}
