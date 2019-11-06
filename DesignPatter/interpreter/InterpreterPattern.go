package interpreter

import (
	"fmt"
	"strings"
)

type Expression interface {
	Interpret(context string) bool
}

type TerminalExpression struct {
	Word string
}

func (te *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, te.Word) {
		return true
	}
	return false
}

type OrExpression struct {
	expre1 Expression
	expre2 Expression
}

func (oe *OrExpression) Interpret(context string) bool {
	return oe.expre1.Interpret(context) || oe.expre2.Interpret(context)
}

type AndExpression struct {
	expre1 Expression
	expre2 Expression
}

func (a *AndExpression) Interpret(context string) bool {
	return a.expre1.Interpret(context) && a.expre2.Interpret(context)
}

func InterpretTest() {
	isMale := &OrExpression{&TerminalExpression{"Robort"}, &TerminalExpression{"John"}}
	isMarriedWoman := &AndExpression{&TerminalExpression{"Juile"}, &TerminalExpression{"Married"}}
	fmt.Println("John is male?", isMale.Interpret("John"))
	fmt.Println("Julie is a married women", isMarriedWoman.Interpret("Married Juile"))
}
