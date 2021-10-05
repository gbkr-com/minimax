package minimax

type factory struct {
	count int
}

func (f *factory) make(ref interface{}) *Decision {
	f.count++
	return &Decision{Reference: ref}
}

// Build is a utility to assist building a decision tree, returning the root
// of that tree and the number of decisions within it. The rules for the tree
// are supplied from the Builder interface.
//
func Build(builder Builder, initial interface{}) (*Decision, int) {
	f := new(factory)
	root := f.make(initial)
	build(root, builder, f)
	return root, f.count
}

func build(decision *Decision, builder Builder, f *factory) {
	responses := builder.Responses(decision.Reference)
	if len(responses) == 0 {
		decision.Score = builder.Estimate(decision.Reference)
	} else {
		for _, ref := range responses {
			response := f.make(ref)
			decision.Responses = append(decision.Responses, response)
			if final, value := builder.IsFinal(ref); final {
				response.Score = value
			} else {
				build(response, builder, f)
			}
		}
	}
}

// The Builder interface used by the Build function.
//
type Builder interface {

	// Responses returns decision references for all the responses to the given
	// decision reference.
	//
	Responses(reference interface{}) []interface{}

	// IsFinal returns true, and gives a score for, the decision reference if
	// it is a final decision.
	//
	IsFinal(reference interface{}) (bool, float64)

	// Estimate a score for a decision reference for which no responses can be
	// created.
	//
	Estimate(reference interface{}) float64
}
