% A list is a 1-D array of numbers.
% A matrix is a 2-D array of numbers, stored in row-major order.

% You may define helper functions here.

% are_adjacent(List, A, B) returns true iff A and B are neighbors in List.
are_adjacent(List, A, B) :-
   (append(_, [A, B | _], List) ; append(_, [B, A | _], List)).

% matrix_transpose(Matrix, Answer) returns true iff Answer is the transpose of
% the 2D matrix Matrix.
matrix_transpose([], []).
matrix_transpose([[]|_], []).
matrix_transpose(Matrix, [Row|Result]) :-
    get_columns(Matrix, Row, RestMatrix),
    matrix_transpose(RestMatrix, Result).
    fail.

get_columns([], [], []).
get_columns([[X|Xs]|Matrix], [X|Column], [Xs|RestMatrix]) :-
    get_columns(Matrix, Column, RestMatrix).

% are_neighbors(Matrix, A, B) returns true iff A and B are neighbors in the 2D
% matrix Matrix.
are_neighbors(Matrix, A, B) :-
    adjacent_position(Matrix, A, PositionA),
    adjacent_position(Matrix, B, PositionB),
    are_adjacent_positions(PositionA, PositionB).

% adjacent_position(Matrix, Element, Position) returns the position of the Element in the Matrix.
adjacent_position(Matrix, Element, [Row, Col]) :-
    nth0(Row, Matrix, RowList),
    nth0(Col, RowList, Element).

are_adjacent_positions([RowA, ColA], [RowB, ColB]) :-
    (RowA =:= RowB, abs(ColA - ColB) =:= 1) ; (ColA =:= ColB, abs(RowA - RowB) =:= 1).
