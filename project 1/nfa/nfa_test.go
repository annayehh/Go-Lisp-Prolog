package nfa

import (
	"reflect"
	"testing"
)

func dagTransitions(st state, sym rune) []state {
	/*
	 * 0 -a-> 1
	 * 0 -a-> 2
	 * 1 -b-> 3
	 * 2 -c-> 3
	 */
	return map[state]map[rune][]state{
		0: map[rune][]state{
			'a': []state{1, 2},
		},
		1: map[rune][]state{
			'b': []state{3},
		},
		2: map[rune][]state{
			'c': []state{3},
		},
	}[st][sym]
}

func expTransitions(st state, sym rune) []state {
	/*
	 * 0 -a-> 1
	 * 0 -a-> 2
	 * 0 -b-> 2
	 * 1 -b->0
	 */
	return map[state]map[rune][]state{
		0: map[rune][]state{
			'a': []state{1, 2},
			'b': []state{2},
		},
		1: map[rune][]state{
			'b': []state{0},
		},
		2: map[rune][]state{},
	}[st][sym]
}

func langTransitions(st state, sym rune) []state {
	/*
	 * 0 -a-> 0
	 * 0 -b-> 1
	 * 1 -a-> 1
	 * 1 -b-> 0
	 */
	return map[state]map[rune][]state{
		0: map[rune][]state{
			'a': []state{0},
			'b': []state{1},
		},
		1: map[rune][]state{
			'a': []state{1},
			'b': []state{0},
		},
	}[st][sym]
}

func multiTransitions(st state, sym rune) []state {
	/*
	 * 0 -a-> 0
	 * 0 -a-> 1
	 * 1 -a-> 0
	 */
	return map[state]map[rune][]state{
		0: map[rune][]state{
			'a': []state{0, 1},
		},
		1: map[rune][]state{
			'a': []state{0},
		},
	}[st][sym]
}

func TestReachable(t *testing.T) {
	tests := []struct {
		label        string
		nfa          TransitionFunction
		start, final state
		input        []rune
		exp          bool
		expected     []state
	}{
		{"dagTransitions", dagTransitions, 0, 0, nil, true, []state{0}},
		{"dagTransitions", dagTransitions, 0, 3, nil, false, nil},
		{"dagTransitions", dagTransitions, 0, 3, []rune{'a', 'b'}, true, []state{0, 1, 3}},
		{"dagTransitions", dagTransitions, 0, 3, []rune{'a', 'c'}, true, []state{0, 2, 3}},
		{"dagTransitions", dagTransitions, 0, 1, []rune{'a'}, true, []state{0, 1}},
		{"dagTransitions", dagTransitions, 0, 3, []rune{'a', 'a'}, false, nil},
		{"dagTransitions", dagTransitions, 0, 3, []rune{'a'}, false, nil},
		{"dagTransitions", dagTransitions, 0, 1, []rune{'b'}, false, nil},
		{"dagTransitions", dagTransitions, 0, 0, []rune{'b'}, false, nil},

		{"expTransitions", expTransitions, 0, 0, []rune{'a', 'b'}, true, []state{0, 1, 0}},
		{"expTransitions", expTransitions, 0, 2, []rune{'a', 'b', 'a'}, true, []state{0, 1, 0, 2}},
		{"expTransitions", expTransitions, 0, 2, []rune{'a', 'b', 'a', 'b', 'a'}, true, []state{0, 1, 0, 1, 0, 2}},
		{"expTransitions", expTransitions, 0, 0, []rune{'a', 'a'}, false, nil},
		{"expTransitions", expTransitions, 0, 2, []rune{'a', 'b', 'a', 'b'}, false, nil},

		{"langTransitions", langTransitions, 0, 0, []rune{'a', 'b', 'b'}, true, []state{0, 0, 1, 0}},
		{"langTransitions", langTransitions, 0, 1, []rune{'a', 'a', 'b'}, true, []state{0, 0, 0, 1}},
		{"langTransitions", langTransitions, 0, 0, []rune{'a', 'a', 'a', 'a', 'a'}, true, []state{0, 0, 0, 0, 0, 0}},
		{"langTransitions", langTransitions, 0, 0, nil, true, []state{0}},
		{"langTransitions", langTransitions, 0, 1, []rune{'a', 'a'}, false, nil},
		{"langTransitions", langTransitions, 0, 0, []rune{'a', 'b', 'a', 'a'}, false, nil},

		// TODO add more tests for 100% test coverage
	}

	for _, test := range tests {
		func() {
			defer func() {
				if recover() != nil {
					t.Errorf("Reachable panicked on (%s, %d, %d, %v)",
						test.label, test.start, test.final, string(test.input))
				}
			}()
			ans, actual := Reachable(test.nfa, test.start, test.final, test.input)

			if ans != test.exp || !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Reachable failed on (%s, %d, %d, %v); expected (%t, %v), got (%t, %v).",
					test.label, test.start, test.final, string(test.input),
					test.exp, test.expected, ans, actual)
			}
		}()
	}
}

func TestMultiReachable(t *testing.T) {
	func() {
		defer func() {
			if recover() != nil {
				t.Error("Reachable panicked on TestMultiReachable.")
			}
		}()
		ans, actual := Reachable(multiTransitions, 0, 0, []rune{'a', 'a'})

		if !ans {
			t.Error("Reachable failed on TestMultiReachable; expected true but got false.")
		} else {
			if len(actual) != 3 {
				t.Errorf("Reachable failed on TestMultiReachable; expected sequence of length 3 but got length %d.",
					len(actual))
			} else if !(actual[0] == 0 && (actual[1] == 1 || actual[1] == 0) && actual[2] == 0) {
				t.Errorf("Reachable failed on TestMultiReachable; expected sequence of states [0, 1, 0] or [0, 0, 0] but got %v.",
					actual)
			}
		}
	}()
}
