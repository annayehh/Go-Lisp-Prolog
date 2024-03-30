package nfa

import "sync"

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

// You may define helper functions here.
var explore func(path []state, inputIndex int)

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) (bool, []state) {
	var wg sync.WaitGroup
	resultsChan := make(chan []state, 10000) // Buffer based on expected concurrency

	explore = func(path []state, inputIndex int) {
		defer wg.Done()
		if inputIndex >= len(input) {
			// Reached the end of the input; check if we're at the final state.
			if path[len(path)-1] == final {
				resultsChan <- path
				return
			}
			return // End of input, not at final state.
		} else {
			currentState := path[len(path)-1]
			nextStates := transitions(currentState, input[inputIndex])
			for _, nextState := range nextStates {
				newPath := append([]state(nil), path...)
				newPath = append(newPath, nextState)
				wg.Add(1)
				go explore(newPath, inputIndex+1)
			}
		}
	}

	// Wait for all worker goroutines to finish.
	wg.Add(1)
	go explore([]state{start}, 0) // start at the beginning of the input string
	wg.Wait()

	// Close resultsChan once all goroutines finish.
	close(resultsChan)

	// collect results from results Channel
	for path := range resultsChan {
		if len(path) > 0 && path[len(path)-1] == final {
			return true, path
		}
	}

	return false, nil
}
