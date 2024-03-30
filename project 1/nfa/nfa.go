package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

// You may define helper functions here.

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) (bool, []state) {
	var initial_route []state = []state{start}
	if input == nil {
		if start == final {
			return true, initial_route
		}
		return false, nil
	}

	var paths [][]state = [][]state{{start}} //Paths will store all possible paths toward the final
	for i, val := range input {
		newPaths := [][]state{}
		for j, path := range paths {
			var curState state
			var nextStates []state
			curState = path[len(path)-1] //get the last state in current path
			nextStates = transitions(curState, val)
			//if nextStates == nil {
			//continue
			//}
			for k, nextState := range nextStates {
				newPath := append([]state{}, path...)
				newPath = append(newPath, nextState)
				newPaths = append(newPaths, newPath)
				k = k
			}
			j = j
		}
		paths = newPaths
		i = i
	}

	for _, path := range paths {
		if (len(path) == len(input)+1) && (path[len(path)-1] == final) {
			return true, path
		}
	}
	return false, nil

}
