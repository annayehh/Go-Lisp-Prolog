reachable(Nfa, FinalState, FinalState, [], [FinalState]).
reachable(Nfa, StartState, FinalState, [Input|AfterInput], Visited) :-
    transition(Nfa, StartState, Input, NextStates),
    member(NextState, NextStates),
    reachable(Nfa, NextState, FinalState, AfterInput, VisitedTail),
    Visited = [StartState|VisitedTail].
