package minimax

import (
	"math"
)

// A Decision by either the score maximiser or the score minimiser. The
// maximiser decides first. Each decision may have a number of responses,
// creating a decision tree. The root of that tree will be a decision containing
// all the possible first decisions of the maximiser.
//
type Decision struct {

	// Any associated application reference for this decision, such as a string
	// or a struct.
	//
	Reference interface{}

	// The numeric score of this decision. For final decisions this score
	// may be assigned by the developer or by an algorithm when the decision
	// tree is built. All intermediary decisions will have a value assigned when
	// the decision tree is evaluated - see the Evaluate function.
	//
	Score float64

	// Possible responses to this decision.
	//
	Responses []*Decision
}

type strategy struct {
	initial float64
	compare func(x, y float64) bool
}

var maximising = &strategy{
	initial: -math.MaxFloat64,
	compare: func(x, y float64) bool { return x > y },
}

var minimising = &strategy{
	initial: math.MaxFloat64,
	compare: func(x, y float64) bool { return x < y },
}

func whichStrategy(maximiser bool) *strategy {
	if maximiser {
		return maximising
	}
	return minimising
}

// Evaluate the given decision, typically the root of the decision tree. This
// will search depth first to propogate scores from the final decisions up to
// the root.
//
func Evaluate(decision *Decision, maximiser bool) {
	s := whichStrategy(maximiser)
	decision.Score = s.initial
	for _, r := range decision.Responses {
		if r.Responses != nil {
			//
			// Evaluate the response from the opposite perspective.
			//
			Evaluate(r, !maximiser)
		}
		if s.compare(r.Score, decision.Score) {
			decision.Score = r.Score
		}
	}
}

// Select the best responses to the given decision.
//
func Select(decision *Decision, maximiser bool) (responses []*Decision) {
	s := whichStrategy(maximiser)
	best := s.initial
	for _, r := range decision.Responses {
		if s.compare(r.Score, best) {
			best = r.Score
			responses = responses[:0]
			responses = append(responses, r)
		} else if r.Score == best {
			responses = append(responses, r)
		}
	}
	return
}
