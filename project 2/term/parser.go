package term

import (
	"errors"
	// "strconv"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <start>    ::= <term> | \epsilon
// <term>     ::= ATOM | NUM | VAR | <compound>
// <compound> ::= <functor> LPAR <args> RPAR
// <functor>  ::= ATOM
// <args>     ::= <term> | <term> COMMA <args>
//

// Parser is the interface for the term parser.
// Do not change the definition of this interface.
type Parser interface {
	Parse(string) (*Term, error)
}

// parser struct holds the information we need during parsing
type parser struct {
	shared_terms map[string]*Term // storing the processed term using a map form its "literal" to the term, notice that the term with the same literal would shared the same object
	toks         []*Token         // tokens
	cp           int              // the current pointer to the examing Token
}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	return &parser{
		shared_terms: make(map[string]*Term),
		toks:         make([]*Token, 0),
		cp:           0}
}

// clean up the mess of the last execution for a new start
func (p *parser) clear() {
	p.shared_terms = make(map[string]*Term)
	p.toks = make([]*Token, 0)
	p.cp = 0
}

// Try to get the shared term.
// If the term is new to the memory, then it would be added to buffer.
// If the term exists, return it.
func (p *parser) getSharedTerm(t *Term) *Term {
	literal := t.Literal
	if p.shared_terms[literal] == nil {
		if t.Typ == TermCompound {
			t.Literal = "" // required by the definition of term
		}
		p.shared_terms[literal] = t
		return t
	} else {
		return p.shared_terms[literal]
	}
}

// Peek at the next i-th unprocessed token.
func (p *parser) peekToken(i int) *Token {
	if i+p.cp >= len(p.toks) || i+p.cp < 0 {
		return nil
	}
	return p.toks[p.cp+i]
}

// Get the next unprocessed token
func (p *parser) nextToken() *Token {
	if p.cp >= len(p.toks) {
		return &Token{tokenEOF, ""}
	}
	t := p.toks[p.cp]
	p.cp += 1
	return t
}

func (p *parser) unfinished() bool {
	return p.cp != len(p.toks)
}

func (p *parser) Parse(input string) (*Term, error) {

	// clear up the mess left in the last execution
	p.clear()

	// tokenization
	lex := newLexer(input)
	for tok, err := lex.next(); tok.typ != tokenEOF; tok, err = lex.next() {
		if err != nil {
			return nil, err
		}
		p.toks = append(p.toks, tok)
	}

	// parsing to build the DAG
	root, err := p.start()
	if p.unfinished() {
		return nil, errors.New("invalid input")
	}
	return root, err
}

// the following functions directly express the given grammar

// return the root of the DAG
// <start>    ::= <term> | \epsilon
func (p *parser) start() (*Term, error) {
	if len(p.toks) == 0 {
		return nil, nil
	}
	return p.term()
}

// return a Term representing <term>
// <term>     ::= ATOM | NUM | VAR | <compound>
func (p *parser) term() (*Term, error) {
	// Check if it is a compound term
	t := p.peekToken(1)
	if t != nil && t.typ == tokenLpar {
		// parse it as compound term
		cterm, err := p.compound()
		if err != nil {
			return nil, err
		}
		return p.getSharedTerm(cterm), nil
	}

	// ATOM | NUM | VAR |
	tok := p.nextToken()
	if tok.typ == tokenAtom {
		t := &Term{TermAtom, tok.literal, nil, nil}
		return p.getSharedTerm(t), nil
	}
	if tok.typ == tokenNumber {
		t := &Term{TermNumber, tok.literal, nil, nil}
		return p.getSharedTerm(t), nil
	}
	if tok.typ == tokenVariable {
		t := &Term{TermVariable, tok.literal, nil, nil}
		return p.getSharedTerm(t), nil
	}
	return nil, errors.New("unexpected reach of code")
}

// return a term representing <compound>
// <compound> ::= <functor> LPAR <args> RPAR
func (p *parser) compound() (*Term, error) {
	ft, fte := p.functor()
	if fte != nil {
		return nil, fte
	}
	ats, ate := p.args()
	if ate != nil {
		return nil, ate
	}

	// get the compound literal
	literal := ft.Literal
	literal += "("
	for i, t := range ats {
		literal += t.Literal
		if i+1 != len(ats) {
			literal += ","
		}
	}
	literal += ")"

	return &Term{TermCompound, literal, ft, ats}, nil
}

// return a Term representing <functor>
// <functor>  ::= ATOM
func (p *parser) functor() (*Term, error) {
	tok := p.nextToken()
	if tok.typ == tokenAtom {
		ft := &Term{TermAtom, tok.literal, nil, nil}
		return p.getSharedTerm(ft), nil
	}
	return nil, errors.New("unexpected reached of code")
}

// return a Term representing <args>
// <args>     ::= (<term>) | (<term> COMMA <args>)
func (p *parser) args() ([]*Term, error) {
	terms := make([]*Term, 0)
	tok := p.nextToken()

	// skip "("
	if tok.typ != tokenLpar {
		return nil, errors.New("expected (")
	}
	for tok.typ != tokenEOF && tok.typ != tokenRpar {
		nextTerm, err := p.term()
		if err != nil {
			return nil, err
		}
		terms = append(terms, nextTerm)
		// skip ","
		tok = p.nextToken()
		if (tok.typ != tokenComma) && (tok.typ != tokenRpar) {
			return nil, errors.New("expected , or )")
		}
	}
	if tok.typ != tokenRpar {
		return nil, errors.New("expected )")
	}

	return terms, nil
}
