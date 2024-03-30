package sexpr

import (
	"errors"
	"math/big" // You will need to use this package in your implementation.
)

// ErrEval is the error value returned by the Evaluator if the contains
// an invalid token.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrEval = errors.New("eval error")

func (expr *SExpr) Eval() (*SExpr, error) {
    // Error handling: check if the expression is nil
    if expr == nil {
        return nil, ErrEval // Use a predefined error for evaluation errors
    }

    // Handle NIL (empty list)
    if expr.isNil() {
        // If the expression is an empty list, return the Lisp representation of NIL
        return mkNil(), nil
    }

    // Handle atoms (numbers and symbols)
    if expr.isAtom() {
        // Numbers evaluate to themselves
        if expr.isNumber() {
            return expr, nil
        }
        // Symbol handling
        if expr.isSymbol() {
            // Special case for the symbol "NIL"
            if expr.atom.literal == "NIL" {
                return mkNil(), nil
            } else if expr.atom.literal == "T" {
				return mkSymbol("T"), nil
			} else if expr.atom.literal == "+" {
				return mkNumber(big.NewInt(0)), nil
			} else if expr.atom.literal == "*" {
				return mkNumber(big.NewInt(1)), nil
			}
            // Look up other symbols in the environment
            return expr, nil
        }

        // If the atom is not a number or a recognized symbol, evaluate it as an atom
        // This might involve checking for predefined symbols or constants
        return expr.EvalAtom()
    }

    // Handle lists (function calls and special forms)
    // If the expression is a list, delegate to EvalList for further evaluation
    return expr.EvalList()
}


// evalList evaluates a list S-expression, interpreting the first element as a function or special form.
func (expr *SExpr) EvalList() (*SExpr, error) {

    operator := expr.car
    operands := expr.cdr


    if operator.isSymbol() {
        switch operator.atom.literal {
        case "QUOTE":
            // QUOTE simply returns the quoted expression
            return operands.car, nil
        case "CAR":
            return operands.car.EvalCar()
        case "CDR":
            return operands.car.EvalCdr()
        case "CONS":
            return EvalCons(operands.car, operands.cdr.car)
        case "LENGTH":
            return operands.car.EvalLength()
        case "+", "*":
            return EvalArithmetic(operator.atom.literal, operands)
        // Add cases for ATOM, LISTP, ZEROP as needed
		case "ATOM":
			return operands.car.EvalAtom()
		case "LISTP":
			return operands.car.EvalListp()
		case "ZEROP":
			return operands.car.EvalZerop()
		default:
			return nil, ErrEval
        }
    }

    return nil, ErrEval
}

// EvalCar returns the first element of the list.
func (expr *SExpr) EvalCar() (*SExpr, error) {
    evaluatedExpr, err := expr.Eval()
    if err != nil {
        return nil, err
    }
    
    if evaluatedExpr.isNil() {
		return mkNil(), nil
	}
	if evaluatedExpr.isConsCell() {
        return evaluatedExpr.car, nil
	}
    return nil, ErrEval
}

// EvalCdr returns the list without its first element.
func (expr *SExpr) EvalCdr() (*SExpr, error) {
    evaluatedExpr, err := expr.Eval()
    if err != nil {
        return nil, ErrEval
    }
    if evaluatedExpr.isNil() {
		return mkNil(), nil
	}
	if evaluatedExpr.isConsCell() {
        return evaluatedExpr.cdr, nil
	}
    return nil, ErrEval
}

// EvalCons creates a new list by prepending an element to another list.
func EvalCons(first, second *SExpr) (*SExpr, error) {
    firstEvaluated, err := first.Eval()
    if err != nil {
        return nil, ErrEval
    }
    secondEvaluated, err := second.Eval()
    if err != nil {
        return nil, ErrEval
    }
    return mkConsCell(firstEvaluated, secondEvaluated), nil
}

// EvalLength calculates the length of a list.
func (expr *SExpr) EvalLength() (*SExpr, error) {
    evaluatedExpr, err := expr.Eval()
    if err != nil {

        return nil, ErrEval
    }
    length := 0
    for !evaluatedExpr.isNil() {
        length++
        evaluatedExpr = evaluatedExpr.cdr
    }
    return mkNumber(big.NewInt(int64(length))), nil
}

// evalArithmetic performs the arithmetic operation represented by the operator on the operands.
func EvalArithmetic(operator string, operands *SExpr) (*SExpr, error) {
    result := big.NewInt(0)
    if operator == "+" {
        result.SetInt64(0) // Start with 0 for addition
    } else if operator == "*" {
        result.SetInt64(1) // Start with 1 for multiplication
    }
    for !operands.isNil() {
        operand, err := operands.car.Eval()
        if err != nil {
            return nil, ErrEval
        }
        if !operand.isNumber() {
            return nil, ErrEval
        }

        if operator == "+" {
            result.Add(result, operand.atom.num)
        } else if operator == "*" {
            result.Mul(result, operand.atom.num)
        }

        operands = operands.cdr
    }
    return mkNumber(result), nil
}

// EvalAtom returns T if the expression is an atom, otherwise NIL.
func (expr *SExpr) EvalAtom() (*SExpr, error) {
    evaluatedExpr, err := expr.Eval()
    if err != nil {
        return nil, err
    }
    if evaluatedExpr.isAtom() {
        return mkSymbol("T"), nil  // "T" for true in Lisp
    }
    return mkNil(), nil  // NIL for false
}

// EvalListp returns T if the expression is a list, otherwise NIL.
func (expr *SExpr) EvalListp() (*SExpr, error) {
    evaluatedExpr, err := expr.Eval()
    if err != nil {
        return nil, err
    }
    if evaluatedExpr.isConsCell() || evaluatedExpr.isNil() {
        return mkSymbol("T"), nil  // "T" for true
    }
    return mkNil(), nil  // NIL for false
}

// EvalZerop returns T if the expression is the number 0, otherwise NIL.
func (expr *SExpr) EvalZerop() (*SExpr, error) {
    evaluatedExpr, err := expr.Eval()
    if err != nil {
        return nil, err
    }
    if evaluatedExpr.isNumber() && evaluatedExpr.atom.num.Cmp(big.NewInt(0)) == 0 {
        return mkSymbol("T"), nil  // "T" for true
    }
    return mkNil(), nil  // NIL for false
}
