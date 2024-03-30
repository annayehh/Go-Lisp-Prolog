package unify

import (
	"errors"
	// "fmt"
	// "hw4/disjointset"
	"hw4/term"
)

// ErrUnifier is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrUnifier = errors.New("unifier error")

// UnifyResult is the result of unification. For example, for a variable term
// `s`, `UnifyResult[s]` is the term which `s` is unified with.
type UnifyResult map[*term.Term]*term.Term

// Unifier is the interface for the term unifier.
// Do not change the definition of this interface
type Unifier interface {
	Unify(*term.Term, *term.Term) (UnifyResult, error)
}

// Implement the Parser interface.
type UnifierImpl struct {
    // unionset size
    size map[*term.Term]int 
    // parents
    class map[*term.Term]*term.Term
    // zeta[C] is a variable only if C consists entirely of variables
    zeta map[*term.Term]*term.Term
    // vars[C] -> all vars in C
    vars map[*term.Term][]*term.Term
    // result of this Unify
    visited map[*term.Term]bool
    acyclic map[*term.Term]bool
    ans UnifyResult
}

// NewUnifier creates a struct of a type that satisfies the Unifier interface.
func NewUnifier() Unifier {
    return &UnifierImpl{}
}

// refer to UnificationChapter.pdf p461 

func (u *UnifierImpl) Unify(termL,termR *term.Term) (UnifyResult, error) {
    u.size = make(map[*term.Term]int)
    u.class = make(map[*term.Term]*term.Term)
    u.zeta = make(map[*term.Term]*term.Term)
    u.vars = make(map[*term.Term][]*term.Term)
    u.ans = make(UnifyResult)
    u.visited = make(map[*term.Term]bool)
    u.acyclic = make(map[*term.Term]bool)
    err := u.UnifyClosure(termL, termR)
    //fmt.Printf("%s %s %s\n", termL.String(), termR.String(), err);
    if err != nil {
        return nil, err
    }
    err = u.FindSolution(termL)
    if err != nil {
        return nil, err
    }
    return u.ans, nil
}

func (u *UnifierImpl) UnifyClosure(termL,termR *term.Term) error {
    termL = u.Find(termL)
    termR = u.Find(termR)
    if termL != termR {
        zetaL := u.zeta[termL]
        zetaR := u.zeta[termR]
        if zetaL.Typ != term.TermVariable && zetaR.Typ != term.TermVariable {
            var zetaLlen, zetaRlen int
            if zetaL.Args == nil {
                zetaLlen = 0
            } else {
                zetaLlen = len(zetaL.Args)
            }
            if zetaR.Args == nil {
                zetaRlen = 0
            } else {
                zetaRlen = len(zetaR.Args)
            }
            if (zetaL.Functor != nil && zetaR.Functor != nil && zetaL.Functor.Literal != zetaR.Functor.Literal) || zetaL.Literal != zetaR.Literal || zetaLlen != zetaRlen {
                //if zetaL.Functor != nil && zetaR.Functor != nil {
                //    fmt.Printf("%s %s %s\n", zetaL.Functor.Literal, zetaR.Functor.Literal, ErrUnifier);
                //}
                return ErrUnifier
            } else {
                if err := u.Union(termL, termR); err != nil {
                    return ErrUnifier
                }
                for index := range zetaL.Args {
                    if err := u.UnifyClosure(zetaL.Args[index], zetaR.Args[index]); err != nil {
                        return ErrUnifier
                    }
                }
            }
        } else {
            if err := u.Union(termL, termR); err != nil {
                return ErrUnifier
            }
        }
    }
    return nil
}

func (u *UnifierImpl) Union(termL,termR *term.Term) error {
    if termL != u.class[termL] {
        return ErrUnifier
    }

    if termR != u.class[termR] {
        return ErrUnifier
    }

    // vars total max size = nlogn
    if u.size[termL] < u.size[termR] {
        termR, termL = termL, termR
    }

    u.size[termL] += u.size[termR]
    u.vars[termL] = append(u.vars[termL], u.vars[termR]...)
    if u.zeta[termL].Typ == term.TermVariable {
        u.zeta[termL] = u.zeta[termR]
    }
    u.class[termR] = termL

    return nil
}

func (u *UnifierImpl) Find(t *term.Term) *term.Term {
    if termParent, ok := u.class[t]; ok {
        if termParent == t {
            return t
        } else {
            u.class[t] = u.Find(termParent)
            return u.class[t]
        }
    }
    // if not in class ,should be init
    u.size[t] = 1
    u.class[t] = t
    u.zeta[t] = t
    u.vars[t] = make([]*term.Term, 0)
    if t.Typ == term.TermVariable {
        u.vars[t] = append(u.vars[t], t)
    }
    u.visited[t] = false
    u.acyclic[t] = false
    return t;
}

func (u *UnifierImpl) FindSolution(t *term.Term) error {
    t = u.zeta[u.Find(t)]
    if u.acyclic[t] {
        return nil
    }
    if u.visited[t] {
        return ErrUnifier
    }
    if t.Typ == term.TermCompound {
        u.visited[t] = true
        for index := range t.Args {
            if err := u.FindSolution(t.Args[index]); err != nil {
                return ErrUnifier
            }
        }
        u.visited[t] = false
    }
    u.acyclic[t] = true
    tParent := u.Find(t)
    for index := range u.vars[tParent] {
        x := u.vars[tParent][index]
        if x != t {
            u.ans[x] = t
        }
    }
    return nil
}
