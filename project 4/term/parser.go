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

// Implement the Parser interface.
type ParserImpl struct {
    lex     *lexer
    peekTok *Token
    terms map[string]*Term
}

func (p *ParserImpl) getTerm(term *Term) (*Term) {
    key := term.String()
    if term, ok := p.terms[key]; ok {
        return term
    }
    p.terms[key] = term
    return term
}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	return &ParserImpl{
        lex: nil,
        peekTok: nil,
        terms: make(map[string]*Term),
    }
}

// Helper function which returns the next token.
func (p *ParserImpl) nextToken() (*Token, error) {
    if tok := p.peekTok; tok != nil {
        p.peekTok = nil
        return tok, nil
    }

    tok, err := p.lex.next()
    if err != nil {
        return nil, ErrParser
    }

    return tok, nil
}

// Helper function which puts a token back as the next token.
func (p *ParserImpl) backToken(tok *Token) {
    p.peekTok = tok
}

// Helper function to peek the next token.
func (p *ParserImpl) peekToken() (*Token, error) {
    tok, err := p.nextToken()
    if err != nil {
        return nil, ErrParser
    }

    p.backToken(tok)

    return tok, nil
}

func (p *ParserImpl) parseTerm() (*Term, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}
	switch tok.typ {
	case tokenEOF:
		return nil, nil
	// comma only in functor(term, ...)
	case tokenComma:
		return nil, ErrParser
	case tokenNumber:
		return p.getTerm(&Term{Typ: TermNumber, Literal: tok.literal}), nil
	case tokenVariable:
		return p.getTerm(&Term{Typ: TermVariable, Literal: tok.literal}), nil
	case tokenAtom:
		atom := p.getTerm(&Term{Typ: TermAtom, Literal: tok.literal})
		nextTok, err := p.peekToken()
		if err != nil {
			return nil, ErrParser
		}
		if nextTok.typ != tokenLpar {
			return atom, nil
		}
		_, err = p.nextToken()
		if err != nil {
			return nil, ErrParser
		}
		// functor(term,...)
		// get first term
		firstArg, err := p.parseTerm()
		if err != nil {
			return nil, ErrParser
		}
		
		functor := &Term{Typ: TermCompound, Literal: "", Functor: atom, Args: []*Term{firstArg}}
		for {
			nextTok, err := p.nextToken()
			if err != nil {
				return nil, ErrParser
			}
			switch nextTok.typ {
			case tokenRpar:
				return p.getTerm(functor), nil
			case tokenComma:
				nextArg, err := p.parseTerm()
				if err != nil {
					return nil, ErrParser
				}
				functor.Args = append(functor.Args, p.getTerm(nextArg))
			default:
				return nil, ErrParser
			}
		}
	default:
		return nil, ErrParser
	}
}

func (p *ParserImpl) Parse(input string) (*Term, error) {
	p.lex = newLexer(input)

	term, err := p.parseTerm()
	if err != nil {
		return nil, ErrParser
	}

    if nextTok, err := p.nextToken(); err != nil || nextTok.typ != tokenEOF {
        return nil, ErrParser
    }                                                                                                                                                                                                                                        
    return term, nil
}
