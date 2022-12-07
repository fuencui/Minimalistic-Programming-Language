package minrkt

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

/*
minrkt Test integration multiple test files
include HW5 and HW4 TEST
open coverage.htm to see coverage
coverage: 90%+ of statements
*/

var test_defineNameMap = make(map[string]Exp)
var defineFuncParser = make(map[string]Exp)
var temporaryMap = make(map[string]Exp)
var HW5_DefineStorage = Hashmap_minrkt{DefineNameMap: test_defineNameMap, DefineFuncParser: defineFuncParser, TemporaryMap: temporaryMap}
var HW5_stack = Stack{StackHM: []Hashmap_minrkt{HW5_DefineStorage}}

func TestMinrkt(t *testing.T) {
	// =============================================================================================//
	// HW5 Test below
	// HW5 TEST 1
	test_hw5_1 := []Token{{1, 0, ""}, {20, 0, ""}, {21, 0, "x"}, {1, 0, ""}, {4, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(define x (+ 1 2))"); err != nil || !cmp.Equal(got, test_hw5_1) {
		t.Errorf("token not match")
	}
	exp_hw5_1 := Parser(test_hw5_1)
	if exp_hw5_1.PrintExp() != "NodeTokenType: (TOK_DEFINE : x NodeTokenType: (TOK_ADD : 1.000000 2.000000 ) ) " {
		t.Errorf("parser tree not match")
	}
	PrintResult(Evaluator(exp_hw5_1, HW5_stack))
	var test_hw5_1_2 Exp
	if v, err := Tokenizer("x"); err == nil {
		test_hw5_1_2 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_1_2, HW5_stack)); err != nil || evalResult.(float64) != 3 {
		t.Errorf("result not match")
	}
	// ---------------------------------------------------------------------------------------------//
	// HW5 TEST 2
	test_hw5_2 := []Token{{1, 0, ""}, {20, 0, ""}, {21, 0, "y"}, {1, 0, ""}, {6, 0, ""}, {21, 0, "x"}, {21, 0, "x"}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(define y (* x x))"); err != nil || !cmp.Equal(got, test_hw5_2) {
		t.Errorf("token not match")
	}
	exp_hw5_2 := Parser(test_hw5_2)
	if exp_hw5_2.PrintExp() != "NodeTokenType: (TOK_DEFINE : y NodeTokenType: (TOK_MUL : x x ) ) " {
		t.Errorf("parser tree not match")
	}
	PrintResult(Evaluator(exp_hw5_2, HW5_stack))
	var test_hw5_2_2 Exp
	if v, err := Tokenizer("y"); err == nil {
		test_hw5_2_2 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_2_2, HW5_stack)); err != nil || evalResult.(float64) != 9 {
		t.Errorf("result not match")
	}
	// ---------------------------------------------------------------------------------------------//
	// HW5 TEST 3
	test_hw5_3 := []Token{{1, 0, ""}, {20, 0, ""}, {21, 0, "z"}, {1, 0, ""}, {4, 0, ""}, {21, 0, "z"}, {3, 1, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(define z (+ z 1))"); err != nil || !cmp.Equal(got, test_hw5_3) {
		t.Errorf("token not match")
	}
	exp_hw5_3 := Parser(test_hw5_3)
	if exp_hw5_3.PrintExp() != "NodeTokenType: (TOK_DEFINE : z NodeTokenType: (TOK_ADD : z 1.000000 ) ) " {
		t.Errorf("parser tree not match")
	}
	PrintResult(Evaluator(exp_hw5_3, HW5_stack))
	var test_hw5_3_1 Exp
	if v, err := Tokenizer("z"); err == nil {
		test_hw5_3_1 = Parser(v)
	}
	Error_HW5_1 := errors.New("cannot reference an identifier before its definition")
	if _, err := PrintResult(Evaluator(test_hw5_3_1, HW5_stack)); Error_HW5_1.Error() != err.Error() {
		t.Errorf(err.Error())
	}
	// ---------------------------------------------------------------------------------------------//
	//HW5 TEST 4
	test_hw5_4 := []Token{{1, 0, ""}, {21, 0, "fib"}, {3, 5, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(define z (+ z 1))"); err != nil || !cmp.Equal(got, test_hw5_3) {
		t.Errorf("token not match")
	}
	exp_hw5_4 := Parser(test_hw5_4)
	if exp_hw5_4.PrintExp() != "NodeTokenType: (TOK_DEF_VAR : fib 5.000000 ) " {
		t.Errorf("parser tree not match")
	}
	PrintResult(Evaluator(exp_hw5_4, HW5_stack))
	var test_hw5_4_1 Exp
	if v, err := Tokenizer("z"); err == nil {
		test_hw5_4_1 = Parser(v)
	}
	Error_HW5_2 := errors.New("cannot reference an identifier before its definition")
	if _, err := PrintResult(Evaluator(test_hw5_4_1, HW5_stack)); Error_HW5_2.Error() != err.Error() {
		t.Errorf(err.Error())
	}

	// ---------------------------------------------------------------------------------------------//
	//HW5 TEST 5
	test_hw5_5 := []Token{{1, 0, ""}, {20, 0, ""}, {1, 0, ""}, {21, 0, "fib"}, {21, 0, "n"},
		{2, 0, ""}, {1, 0, ""}, {16, 0, ""}, {1, 0, ""}, {13, 0, ""}, {21, 0, "n"},
		{3, 2, ""}, {2, 0, ""}, {21, 0, "n"}, {1, 0, ""}, {4, 0, ""}, {1, 0, ""}, {21, 0, "fib"},
		{1, 0, ""}, {5, 0, ""}, {21, 0, "n"}, {3, 1, ""}, {2, 0, ""}, {2, 0, ""}, {1, 0, ""},
		{21, 0, "fib"}, {1, 0, ""}, {5, 0, ""}, {21, 0, "n"}, {3, 2, ""}, {2, 0, ""},
		{2, 0, ""}, {2, 0, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(define (fib n) (if (< n 2) n (+ (fib (- n 1)) (fib (- n 2)))))"); err != nil || !cmp.Equal(got, test_hw5_5) {
		t.Errorf("token not match")
	}
	exp_hw5_5 := Parser(test_hw5_5)
	if exp_hw5_5.PrintExp() != "NodeTokenType: (TOK_DEFINE : NodeTokenType: (TOK_DEF_VAR : fib n ) "+
		"NodeTokenType: (TOK_COND_IF : NodeTokenType: (TOK_COMP_S : n 2.000000 ) "+
		"n NodeTokenType: (TOK_ADD : NodeTokenType: (TOK_DEF_VAR : fib NodeTokenType: (TOK_SUB : n 1.000000 ) ) "+
		"NodeTokenType: (TOK_DEF_VAR : fib NodeTokenType: (TOK_SUB : n 2.000000 ) ) ) ) ) " {
		t.Errorf("parser tree not match")
	}
	PrintResult(Evaluator(exp_hw5_5, HW5_stack))
	var test_hw5_5_1 Exp
	if v, err := Tokenizer("(fib 0)"); err == nil {
		test_hw5_5_1 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_5_1, HW5_stack)); err != nil || evalResult.(float64) != 0 {
		t.Errorf("result not match")
	}

	var test_hw5_5_2 Exp
	if v, err := Tokenizer("(fib 1)"); err == nil {
		test_hw5_5_2 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_5_2, HW5_stack)); err != nil || evalResult.(float64) != 1 {
		t.Errorf("result not match")
	}

	var test_hw5_5_3 Exp
	if v, err := Tokenizer("(fib 2)"); err == nil {
		test_hw5_5_3 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_5_3, HW5_stack)); err != nil || evalResult.(float64) != 1 {
		t.Errorf("result not match")
	}

	var test_hw5_5_4 Exp
	if v, err := Tokenizer("(fib 3)"); err == nil {
		test_hw5_5_4 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_5_4, HW5_stack)); err != nil || evalResult.(float64) != 2 {
		t.Errorf("result not match")
	}

	var test_hw5_5_5 Exp
	if v, err := Tokenizer("(fib 4)"); err == nil {
		test_hw5_5_5 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_5_5, HW5_stack)); err != nil || evalResult.(float64) != 3 {
		t.Errorf("result not match")
	}

	var test_hw5_5_6 Exp
	if v, err := Tokenizer("(fib 12)"); err == nil {
		test_hw5_5_6 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_5_6, HW5_stack)); err != nil || evalResult.(float64) != 144 {
		t.Errorf("result not match")
	}

	var test_hw5_5_7 Exp
	if v, err := Tokenizer("(fib (+ x 1))"); err == nil {
		test_hw5_5_7 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_5_7, HW5_stack)); err != nil || evalResult.(float64) != 3 {
		t.Errorf("result not match")
	}

	var test_hw5_5_8 Exp
	if v, err := Tokenizer("(fib (+ (* 2 x) y))"); err == nil {
		test_hw5_5_8 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_5_8, HW5_stack)); err != nil || evalResult.(float64) != 610 {
		t.Errorf("result not match")
	}

	var test_hw5_5_9 Exp
	if v, err := Tokenizer("(fib 1 2)"); err == nil {
		test_hw5_5_9 = Parser(v)
	}

	Error_HW5_3 := errors.New("variable number error")
	if _, err := PrintResult(Evaluator(test_hw5_5_9, HW5_stack)); Error_HW5_3.Error() != err.Error() {
		t.Errorf(err.Error())
	}

	// ---------------------------------------------------------------------------------------------//
	//HW5 TEST 6

	test_hw5_6 := []Token{{1, 0, ""}, {20, 0, ""}, {1, 0, ""}, {21, 0, "addx"}, {21, 0, "a"},
		{2, 0, ""}, {1, 0, ""}, {4, 0, ""}, {21, 0, "a"}, {21, 0, "x"}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(define (addx a) (+ a x))"); err != nil || !cmp.Equal(got, test_hw5_6) {
		t.Errorf("token not match")
	}
	exp_hw5_6 := Parser(test_hw5_6)
	if exp_hw5_6.PrintExp() != "NodeTokenType: (TOK_DEFINE : NodeTokenType: (TOK_DEF_VAR : addx a ) NodeTokenType: (TOK_ADD : a x ) ) " {
		t.Errorf("parser tree not match")
	}
	PrintResult(Evaluator(exp_hw5_6, HW5_stack))
	var test_hw5_6_1 Exp
	if v, err := Tokenizer("(addx 12)"); err == nil {
		test_hw5_6_1 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_6_1, HW5_stack)); err != nil || evalResult.(float64) != 15 {
		t.Errorf("result not match")
	}
	var test_hw5_6_2 Exp
	if v, err := Tokenizer("(define x 5)"); err == nil {
		test_hw5_6_2 = Parser(v)
	}
	if _, err := PrintResult(Evaluator(test_hw5_6_2, HW5_stack)); err != nil {
		t.Errorf("result not match")
	}
	var test_hw5_6_3 Exp
	if v, err := Tokenizer("(addx 12)"); err == nil {
		test_hw5_6_3 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_6_3, HW5_stack)); err != nil || evalResult.(float64) != 17 {
		t.Errorf("result not match")
	}

	// ---------------------------------------------------------------------------------------------//
	//HW5 TEST 7
	test_hw5_7 := []Token{{1, 0, ""}, {20, 0, ""}, {1, 0, ""}, {21, 0, "xplusy"}, {2, 0, ""},
		{1, 0, ""}, {4, 0, ""}, {21, 0, "x"}, {21, 0, "y"}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(define (xplusy) (+ x y))"); err != nil || !cmp.Equal(got, test_hw5_7) {
		t.Errorf("token not match")
	}
	exp_hw5_7 := Parser(test_hw5_7)
	if exp_hw5_7.PrintExp() != "NodeTokenType: (TOK_DEFINE : NodeTokenType: (TOK_DEF_VAR : xplusy ) NodeTokenType: (TOK_ADD : x y ) ) " {
		t.Errorf("parser tree not match")
	}
	PrintResult(Evaluator(exp_hw5_7, HW5_stack))
	var test_hw5_7_1 Exp
	if v, err := Tokenizer("(xplusy)"); err == nil {
		test_hw5_7_1 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_7_1, HW5_stack)); err != nil || evalResult.(float64) != 30 {

		t.Errorf("result not match")
	}
	var test_hw5_7_2 Exp
	if v, err := Tokenizer("(define x 11)"); err == nil {
		test_hw5_7_2 = Parser(v)
	}
	if _, err := PrintResult(Evaluator(test_hw5_7_2, HW5_stack)); err != nil {
		t.Errorf("result not match")
	}
	var test_hw5_7_3 Exp
	if v, err := Tokenizer("(define y 22)"); err == nil {
		test_hw5_7_3 = Parser(v)
	}
	if _, err := PrintResult(Evaluator(test_hw5_7_3, HW5_stack)); err != nil {
		t.Errorf("result not match")
	}

	var test_hw5_7_4 Exp
	if v, err := Tokenizer("(xplusy)"); err == nil {
		test_hw5_7_4 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_7_4, HW5_stack)); err != nil || evalResult.(float64) != 33 {
		t.Errorf("result not match")
	}

	// ---------------------------------------------------------------------------------------------//
	//HW5 TEST 8
	test_hw5_8 := []Token{{1, 0, ""}, {20, 0, ""}, {1, 0, ""}, {21, 0, "add1"}, {21, 0, "x"},
		{2, 0, ""}, {1, 0, ""}, {4, 0, ""}, {21, 0, "x"}, {3, 1, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(define (add1 x) (+ x 1))"); err != nil || !cmp.Equal(got, test_hw5_8) {
		t.Errorf("token not match")
	}
	exp_hw5_8 := Parser(test_hw5_8)
	if exp_hw5_8.PrintExp() != "NodeTokenType: (TOK_DEFINE : NodeTokenType: (TOK_DEF_VAR : add1 x ) NodeTokenType: (TOK_ADD : x 1.000000 ) ) " {
		t.Errorf("parser tree not match")
	}
	PrintResult(Evaluator(exp_hw5_8, HW5_stack))
	test_hw5_8_1 := []Token{{1, 0, ""}, {20, 0, ""}, {1, 0, ""}, {21, 0, "add2"}, {21, 0, "x"},
		{2, 0, ""}, {1, 0, ""}, {21, 0, "add1"}, {1, 0, ""}, {21, 0, "add1"}, {21, 0, "x"},
		{2, 0, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(define (add2 x) (add1 (add1 x)))"); err != nil || !cmp.Equal(got, test_hw5_8_1) {
		t.Errorf("token not match")
	}
	exp_hw5_8_1 := Parser(test_hw5_8_1)
	if exp_hw5_8.PrintExp() != "NodeTokenType: (TOK_DEFINE : NodeTokenType: (TOK_DEF_VAR : add1 x ) NodeTokenType: (TOK_ADD : x 1.000000 ) ) " {
		t.Errorf("parser tree not match")
	}
	PrintResult(Evaluator(exp_hw5_8_1, HW5_stack))
	var test_hw5_8_1_1 Exp
	if v, err := Tokenizer("(add1 12)"); err == nil {
		test_hw5_8_1_1 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_8_1_1, HW5_stack)); err != nil || evalResult.(float64) != 13 {
		t.Errorf("result not match")
	}
	var test_hw5_8_1_2 Exp
	if v, err := Tokenizer("(add2 12)"); err == nil {
		test_hw5_8_1_2 = Parser(v)
	}
	if evalResult, err := PrintResult(Evaluator(test_hw5_8_1_2, HW5_stack)); err != nil || evalResult.(float64) != 14 {

		t.Errorf("result not match")
	}
	// =============================================================================================//
	// =============================================================================================//
	// =============================================================================================//
	// HW4 Test below
	DefineStorage := Hashmap_minrkt{}
	stack := Stack{StackHM: []Hashmap_minrkt{DefineStorage}}
	test1 := []Token{{1, 0, ""}, {4, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(+ 1 2)"); err != nil || !cmp.Equal(got, test1) {
		t.Errorf("token not match")
	}
	exp1 := Parser(test1)
	if exp1.PrintExp() != "NodeTokenType: (TOK_ADD : 1.000000 2.000000 ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp1, stack)); err != nil || evalResult.(float64) != 3 {
		t.Errorf("result not match")
	}

	test2 := []Token{{1, 0, ""}, {4, 0, ""}, {3, 1.1, ""}, {3, 1.2, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(+ 1.1 1.2)"); err != nil || !cmp.Equal(got, test2) {
		t.Errorf("token not match")
	}
	exp2 := Parser(test2)
	if exp2.PrintExp() != "NodeTokenType: (TOK_ADD : 1.100000 1.200000 ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp2, stack)); err != nil || evalResult.(float64) != 2.3 {
		t.Errorf("result not match")
	}

	test3 := []Token{{1, 0, ""}, {6, 0, ""}, {1, 0, ""}, {4, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {3, 3, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(* (+ 1 2) 3)"); err != nil || !cmp.Equal(got, test3) {
		t.Errorf("token not match")
	}
	exp3 := Parser(test3)
	if exp3.PrintExp() != "NodeTokenType: (TOK_MUL : NodeTokenType: (TOK_ADD : 1.000000 2.000000 ) 3.000000 ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp3, stack)); err != nil || evalResult.(float64) != 9 {
		t.Errorf("result not match")
	}

	test4 := []Token{{1, 0, ""}, {7, 0, ""}, {3, 5, ""}, {3, 4, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(/5 4)"); err != nil || !cmp.Equal(got, test4) {
		t.Errorf("token not match")
	}
	exp4 := Parser(test4)
	if exp4.PrintExp() != "NodeTokenType: (TOK_DIV : 5.000000 4.000000 ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp4, stack)); err != nil || evalResult.(float64) != 1.25 {
		t.Errorf("result not match")
	}

	test5 := []Token{{1, 0, ""}, {5, 0, ""}, {3, 1.8, ""}, {3, 12.2, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(- 1.8 12.2)"); err != nil || !cmp.Equal(got, test5) {
		t.Errorf("token not match")
	}
	exp5 := Parser(test5)
	if exp5.PrintExp() != "NodeTokenType: (TOK_SUB : 1.800000 12.200000 ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp5, stack)); err != nil || evalResult.(float64) != -10.399999999999999 {
		t.Errorf("result not match")
	}

	ErrInvalidToken := errors.New("invalid token: invalid token")
	if _, err := Tokenizer("(ASD 0.1 1.2)"); err.Error() != ErrInvalidToken.Error() {
		t.Errorf("token not match")
	}
	if _, err := Tokenizer(""); err.Error() != ErrInvalidToken.Error() {
		t.Errorf("token not match")
	}
	// new test
	test10 := []Token{{14, 0, ""}}
	if got, err := Tokenizer("true"); err != nil || !cmp.Equal(got, test10) {
		t.Errorf("token not match")
	}
	exp10 := Parser(test10)
	if exp10.PrintExp() != "NodeTokenType: (TOK_BOOL_T : ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp10, stack)); err != nil || evalResult.(bool) != true {
		t.Errorf("result not match")
	}
	//new test
	test11 := []Token{{15, 0, ""}}
	if got, err := Tokenizer("false"); err != nil || !cmp.Equal(got, test11) {
		t.Errorf("token not match")
	}
	exp11 := Parser(test11)
	if exp11.PrintExp() != "NodeTokenType: (TOK_BOOL_F : ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp11, stack)); err != nil || evalResult.(bool) != false {
		t.Errorf("result not match")
	}

	// new test
	test12 := []Token{{1, 0, ""}, {9, 0, ""}, {1, 0, ""}, {4, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {3, 3, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(= (+ 1 2) 3)"); err != nil || !cmp.Equal(got, test12) {
		t.Errorf("token not match")
	}
	exp12 := Parser(test12)
	if exp12.PrintExp() != "NodeTokenType: (TOK_COMP_EQ : NodeTokenType: (TOK_ADD : 1.000000 2.000000 ) 3.000000 ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp12, stack)); err != nil || evalResult.(bool) != true {
		t.Errorf("result not match")
	}
	// new test
	test13 := []Token{{1, 0, ""}, {9, 0, ""}, {1, 0, ""}, {4, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {1, 0, ""}, {5, 0, ""},
		{3, 4, ""}, {3, 3, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(= (+ 1 2) (- 4 3))"); err != nil || !cmp.Equal(got, test13) {
		t.Errorf("token not match")
	}
	exp13 := Parser(test13)
	if exp13.PrintExp() != "NodeTokenType: (TOK_COMP_EQ : NodeTokenType: (TOK_ADD : 1.000000 2.000000 ) NodeTokenType: (TOK_SUB : 4.000000 3.000000 ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp13, stack)); err != nil || evalResult.(bool) != false {
		t.Errorf("result not match")
	}
	// new test
	test14 := []Token{{1, 0, ""}, {10, 0, ""}, {1, 0, ""}, {4, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {1, 0, ""}, {5, 0, ""},
		{3, 3, ""}, {1, 0, ""}, {6, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(>= (+ 1 2) (- 3 (* 1 2)))"); err != nil || !cmp.Equal(got, test14) {
		t.Errorf("token not match")
	}
	exp14 := Parser(test14)
	if exp14.PrintExp() != "NodeTokenType: (TOK_COMP_GE : NodeTokenType: (TOK_ADD : 1.000000 2.000000 ) "+
		"NodeTokenType: (TOK_SUB : 3.000000 NodeTokenType: (TOK_MUL : 1.000000 2.000000 ) ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp14, stack)); err != nil || evalResult.(bool) != true {
		t.Errorf("result not match")
	}
	// new test
	test15 := []Token{{1, 0, ""}, {13, 0, ""}, {1, 0, ""}, {4, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {1, 0, ""}, {5, 0, ""},
		{3, 3, ""}, {1, 0, ""}, {6, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(< (+ 1 2) (- 3 (* 1 2)))"); err != nil || !cmp.Equal(got, test15) {
		t.Errorf("token not match")

	}
	exp15 := Parser(test15)
	if exp15.PrintExp() != "NodeTokenType: (TOK_COMP_S : NodeTokenType: (TOK_ADD : 1.000000 2.000000 ) "+
		"NodeTokenType: (TOK_SUB : 3.000000 NodeTokenType: (TOK_MUL : 1.000000 2.000000 ) ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp15, stack)); err != nil || evalResult.(bool) != false {
		t.Errorf("result not match")
	}
	// new test
	test16 := []Token{{1, 0, ""}, {17, 0, ""}, {1, 0, ""}, {9, 0, ""}, {3, 1, ""}, {3, 1, ""}, {2, 0, ""}, {1, 0, ""}, {11, 0, ""},
		{1, 0, ""}, {7, 0, ""}, {3, 1, ""}, {3, 0, ""}, {2, 0, ""}, {3, 3, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(and (= 1 1) (> (/ 1 0) 3))"); err != nil || !cmp.Equal(got, test16) {
		t.Errorf("token not match")
	}
	exp16 := Parser(test16)
	if exp16.PrintExp() != "NodeTokenType: (TOK_COND_AND : NodeTokenType: (TOK_COMP_EQ : 1.000000 1.000000 ) "+
		"NodeTokenType: (TOK_COMP_G : NodeTokenType: (TOK_DIV : 1.000000 0.000000 ) 3.000000 ) ) " {
		t.Errorf("parser tree not match")
	}
	ErrInvalidToken2 := errors.New("invalid expression")
	if _, err := PrintResult(Evaluator(exp16, stack)); err.Error() != ErrInvalidToken2.Error() {
		t.Errorf("token not match")
	}

	// new test
	test17 := []Token{{1, 0, ""}, {17, 0, ""}, {1, 0, ""}, {9, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {1, 0, ""}, {11, 0, ""},
		{1, 0, ""}, {7, 0, ""}, {3, 1, ""}, {3, 0, ""}, {2, 0, ""}, {3, 3, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(and (= 1 2) (> (/ 1 0) 3))"); err != nil || !cmp.Equal(got, test17) {
		t.Errorf("token not match")
	}
	exp17 := Parser(test17)
	if exp17.PrintExp() != "NodeTokenType: (TOK_COND_AND : NodeTokenType: (TOK_COMP_EQ : 1.000000 2.000000 "+
		") NodeTokenType: (TOK_COMP_G : NodeTokenType: (TOK_DIV : 1.000000 0.000000 ) 3.000000 ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp17, stack)); err != nil || evalResult.(bool) != false {
		t.Errorf("result not match")
	}
	// new test
	test18 := []Token{{1, 0, ""}, {17, 0, ""}, {1, 0, ""}, {19, 0, ""}, {1, 0, ""}, {9, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {2, 0, ""},
		{1, 0, ""}, {18, 0, ""}, {1, 0, ""}, {11, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {1, 0, ""}, {9, 0, ""},
		{1, 0, ""}, {4, 0, ""}, {3, 3, ""}, {3, 4, ""}, {2, 0, ""}, {3, 7, ""}, {2, 0, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(and (not (= 1 2)) (or (> 1 2) (= (+ 3 4) 7)))"); err != nil || !cmp.Equal(got, test18) {
		t.Errorf("token not match")
	}
	exp18 := Parser(test18)
	if exp18.PrintExp() != "NodeTokenType: (TOK_COND_AND : NodeTokenType: (TOK_COND_NOT : NodeTokenType: "+
		"(TOK_COMP_EQ : 1.000000 2.000000 ) ) NodeTokenType: (TOK_COND_OR : NodeTokenType: "+
		"(TOK_COMP_G : 1.000000 2.000000 ) NodeTokenType: (TOK_COMP_EQ : "+
		"NodeTokenType: (TOK_ADD : 3.000000 4.000000 ) 7.000000 ) ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp18, stack)); err != nil || evalResult.(bool) != true {
		t.Errorf("result not match")
	}

	// new test
	test19 := []Token{{1, 0, ""}, {18, 0, ""}, {14, 0, ""}, {15, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(or true false)"); err != nil || !cmp.Equal(got, test19) {
		t.Errorf("token not match")
	}
	exp19 := Parser(test19)
	if exp19.PrintExp() != "NodeTokenType: (TOK_COND_OR : NodeTokenType: (TOK_BOOL_T : ) NodeTokenType: (TOK_BOOL_F : ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp19, stack)); err != nil || evalResult.(bool) != true {
		t.Errorf("result not match")
	}

	// new test
	test20 := []Token{{1, 0, ""}, {17, 0, ""}, {14, 0, ""}, {1, 0, ""}, {9, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(and true (= 1 2))"); err != nil || !cmp.Equal(got, test20) {
		t.Errorf("token not match")
	}
	exp20 := Parser(test20)
	if exp20.PrintExp() != "NodeTokenType: (TOK_COND_AND : NodeTokenType: (TOK_BOOL_T : ) NodeTokenType: (TOK_COMP_EQ : 1.000000 2.000000 ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp20, stack)); err != nil || evalResult.(bool) != false {
		t.Errorf("result not match")
	}
	// new test
	test21 := []Token{{1, 0, ""}, {18, 0, ""}, {1, 0, ""}, {9, 0, ""}, {3, 2, ""}, {1, 0, ""}, {4, 0, ""}, {3, 1, ""}, {3, 1, ""}, {2, 0, ""},
		{2, 0, ""}, {15, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(or (= 2 (+ 1 1)) false)"); err != nil || !cmp.Equal(got, test21) {
		t.Errorf("token not match")
	}
	exp21 := Parser(test21)
	if exp21.PrintExp() != "NodeTokenType: (TOK_COND_OR : NodeTokenType: (TOK_COMP_EQ : 2.000000 NodeTokenType: (TOK_ADD : 1.000000 1.000000 ) ) NodeTokenType: (TOK_BOOL_F : ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp21, stack)); err != nil || evalResult.(bool) != true {
		t.Errorf("result not match")
	}
	// new test
	test22 := []Token{{1, 0, ""}, {19, 0, ""}, {14, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(not true)"); err != nil || !cmp.Equal(got, test22) {
		t.Errorf("token not match")
	}
	exp22 := Parser(test22)
	if exp22.PrintExp() != "NodeTokenType: (TOK_COND_NOT : NodeTokenType: (TOK_BOOL_T : ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp22, stack)); err != nil || evalResult.(bool) != false {
		t.Errorf("result not match")
	}
	// new test
	test23 := []Token{{1, 0, ""}, {19, 0, ""}, {15, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(not false)"); err != nil || !cmp.Equal(got, test23) {
		t.Errorf("token not match")
	}
	exp23 := Parser(test23)
	if exp23.PrintExp() != "NodeTokenType: (TOK_COND_NOT : NodeTokenType: (TOK_BOOL_F : ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp23, stack)); err != nil || evalResult.(bool) != true {
		t.Errorf("result not match")
	}
	// new test
	test24 := []Token{{1, 0, ""}, {16, 0, ""}, {1, 0, ""}, {11, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {1, 0, ""}, {4, 0, ""},
		{3, 3, ""}, {3, 4, ""}, {2, 0, ""}, {1, 0, ""}, {6, 0, ""}, {3, 5, ""}, {3, 6, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(if (> 1 2) (+ 3 4) (* 5 6))"); err != nil || !cmp.Equal(got, test24) {
		t.Errorf("token not match")
	}
	exp24 := Parser(test24)
	if exp24.PrintExp() != "NodeTokenType: (TOK_COND_IF : NodeTokenType: (TOK_COMP_G : 1.000000 2.000000 )"+
		" NodeTokenType: (TOK_ADD : 3.000000 4.000000 ) NodeTokenType: (TOK_MUL : 5.000000 6.000000 ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp24, stack)); err != nil || evalResult.(float64) != 30 {
		t.Errorf("result not match")
	}
	// new test
	test25 := []Token{{1, 0, ""}, {16, 0, ""}, {1, 0, ""}, {17, 0, ""}, {1, 0, ""}, {11, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {1, 0, ""}, {9, 0, ""},
		{3, 3, ""}, {3, 4, ""}, {2, 0, ""}, {2, 0, ""}, {1, 0, ""}, {7, 0, ""}, {3, 1, ""}, {3, 0, ""}, {2, 0, ""}, {1, 0, ""}, {18, 0, ""}, {14, 0, ""}, {15, 0, ""},
		{2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(if (and (> 1 2) (= 3 4)) (/ 1 0) (or true false))"); err != nil || !cmp.Equal(got, test25) {
		t.Errorf("token not match")
	}
	exp25 := Parser(test25)
	if exp25.PrintExp() != "NodeTokenType: (TOK_COND_IF : NodeTokenType: (TOK_COND_AND : NodeTokenType: "+
		"(TOK_COMP_G : 1.000000 2.000000 ) NodeTokenType: (TOK_COMP_EQ : 3.000000 4.000000 ) ) NodeTokenType:"+
		" (TOK_DIV : 1.000000 0.000000 ) NodeTokenType: (TOK_COND_OR : NodeTokenType: (TOK_BOOL_T : ) "+
		"NodeTokenType: (TOK_BOOL_F : ) ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp25, stack)); err != nil || evalResult.(bool) != true {
		t.Errorf("result not match")
	}
	// new test
	test26 := []Token{{1, 0, ""}, {16, 0, ""}, {1, 0, ""}, {17, 0, ""}, {1, 0, ""}, {11, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {1, 0, ""}, {9, 0, ""},
		{3, 3, ""}, {3, 4, ""}, {2, 0, ""}, {2, 0, ""}, {1, 0, ""}, {7, 0, ""}, {3, 1, ""}, {3, 0, ""}, {2, 0, ""}, {1, 0, ""}, {18, 0, ""}, {15, 0, ""},
		{1, 0, ""}, {16, 0, ""}, {14, 0, ""}, {15, 0, ""}, {14, 0, ""}, {2, 0, ""}, {2, 0, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(if (and (> 1 2) (= 3 4)) (/ 1 0) (or false (if true false true)))"); err != nil || !cmp.Equal(got, test26) {
		t.Errorf("token not match")
	}
	exp26 := Parser(test26)
	if exp26.PrintExp() != "NodeTokenType: (TOK_COND_IF : NodeTokenType: (TOK_COND_AND : NodeTokenType: "+
		"(TOK_COMP_G : 1.000000 2.000000 ) NodeTokenType: (TOK_COMP_EQ : 3.000000 4.000000 ) ) NodeTokenType: "+
		"(TOK_DIV : 1.000000 0.000000 ) NodeTokenType: (TOK_COND_OR : NodeTokenType: (TOK_BOOL_F : ) "+
		"NodeTokenType: (TOK_COND_IF : NodeTokenType: (TOK_BOOL_T : ) NodeTokenType: (TOK_BOOL_F : )"+
		" NodeTokenType: (TOK_BOOL_T : ) ) ) ) " {
		t.Errorf("parser tree not match")
	}
	if evalResult, err := PrintResult(Evaluator(exp26, stack)); err != nil || evalResult.(bool) != false {
		t.Errorf("result not match")
	}
	// new test
	test27 := []Token{{1, 0, ""}, {11, 0, ""}, {1, 0, ""}, {9, 0, ""}, {3, 1, ""}, {3, 2, ""}, {2, 0, ""}, {3, 3, ""}, {2, 0, ""}}
	if got, err := Tokenizer("(> (= 1 2) 3)"); err != nil || !cmp.Equal(got, test27) {
		t.Errorf("token not match")
	}
	exp27 := Parser(test27)
	if exp27.PrintExp() != "NodeTokenType: (TOK_COMP_G : NodeTokenType: (TOK_COMP_EQ : 1.000000 2.000000 ) 3.000000 ) " {
		t.Errorf("parser tree not match")
	}

	ErrInvalidToken3 := errors.New("invalid expression")
	if _, err := PrintResult(Evaluator(exp27, stack)); err.Error() != ErrInvalidToken3.Error() {
		t.Errorf("token not match")
	}

}
