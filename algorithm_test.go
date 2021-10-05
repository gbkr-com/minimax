package minimax

import (
	"testing"
)

func TestConvoy(t *testing.T) {
	//
	// https://cs.stanford.edu/people/eroberts/courses/soco/projects/1998-99/game-theory/Minimax.html
	//
	plan := &Decision{
		Reference: "",
		Responses: []*Decision{
			{
				Reference: "Allies to North",
				Responses: []*Decision{
					{
						Reference: "Enemy to North",
						Score:     2,
					},
					{
						Reference: "Enemy to South",
						Score:     2,
					},
				},
			},
			{
				Reference: "Allies to South",
				Responses: []*Decision{
					{
						Reference: "Enemy to North",
						Score:     1,
					},
					{
						Reference: "Enemy to South",
						Score:     3,
					},
				},
			},
		},
	}
	//
	// Terminal cases.
	//
	Evaluate(plan.Responses[1], false)
	if plan.Responses[1].Score != 1.0 {
		t.Error()
	}
	Evaluate(plan.Responses[1], true)
	if plan.Responses[1].Score != 3.0 {
		t.Error()
	}
	//
	// Full tree.
	//
	Evaluate(plan, true)
	action := Select(plan, true)
	if len(action) != 1 {
		t.Error()
	}
	if action[0] != plan.Responses[0] {
		t.Error()
	}
}
