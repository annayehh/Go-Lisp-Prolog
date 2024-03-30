:- initialization main.

main :-
    consult(['transitions.pl', 'nfa.pl']),
    (show_coverage(run_tests) ; true),
    halt.

:- begin_tests(nfa).

test(nfaExp1, all(X == [[0,2]])) :- reachable(expTransitions, 0, 2, [a], X).
test(nfaExp2, all(X == [[0,2]])) :- reachable(expTransitions, 0, 2, [b], X).
test(nfaExp3, all(X ==[[0, 1, 0, 1]])) :- reachable(expTransitions, 0, 1, [a, b, a], X).
test(nfaExp4, [fail])   :- reachable(expTransitions, 0, 1, [a, b, a, b], _).
test(nfaExp5, all(X == [[0, 1, 0, 2]])) :- reachable(expTransitions, 0, 2, [a, b, a], X).

test(nfaFoo1a, all(X == [[0, 1, 3]])) :- reachable(fooTransitions, 0, 3, [a, b], X).
test(nfaFoo1b, [fail]) :- reachable(fooTransitions, 0, 3, [a, b], [1, 3]).

test(nfaFoo2, all(X == [[0, 2, 3]])) :- reachable(fooTransitions, 0, 3, [a, c], X).
test(nfaFoo3, all(X == [[0, 1]])) :- reachable(fooTransitions, 0, 1, [a], X).
test(nfaFoo4, [fail])   :- reachable(fooTransitions, 0, 3, [a, a], _).
test(nfaFoo5, [fail])   :- reachable(fooTransitions, 0, 3, [a], _).
test(nfaFoo6, [fail])   :- reachable(fooTransitions, 0, 1, [b], _).

test(nfaLang1, all(X == [[0, 0, 1, 0]])) :- reachable(langTransitions, 0, 0, [a, b, b], X).
test(nfaLang2, all(X == [[0, 0, 0, 1]])) :- reachable(langTransitions, 0, 1, [a, a, b], X).
test(nfaLang3, all(X == [[0, 0, 0, 0, 0, 0]])) :- reachable(langTransitions, 0, 0, [a, a, a, a, a], X).
test(nfaLang4, [fail])   :- reachable(langTransitions, 0, 1, [a, a], _).
test(nfaLang5, [fail])   :- reachable(langTransitions, 0, 0, [a, b, a, a], _).

:- end_tests(nfa).
