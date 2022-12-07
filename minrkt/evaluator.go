package minrkt

import (
	"fmt"
	"reflect"
)

// global variables Hashmap_minrkt stack
var stack Stack

func Evaluator(exp Exp, Stack Stack) Exp {
	stack = Stack
	return exp
}

/*
Print Result nicely
*/
func PrintResult(exp Exp) (interface{}, error) {

	res, err := exp.eval()
	if err != nil {
		fmt.Println("[FINAL RESULT]: " + fmt.Sprintf("%s", err))
		return 0, err
	}
	if res == nil {
		return nil, nil
	}

	if t := reflect.TypeOf(res); t.String() == "float64" {
		res = reflect.ValueOf(res).Interface().(float64)
		fmt.Println("[FINAL RESULT]: " + fmt.Sprintf("%f", res))
	}
	if t := reflect.TypeOf(res); t.String() == "bool" {
		res = reflect.ValueOf(res).Interface().(bool)
		if res == true {
			fmt.Println("[FINAL RESULT]: #t")
		} else {
			fmt.Println("[FINAL RESULT]: #f")
		}

	}
	return res, nil
}

/*
A Exp interface
*/
type Exp interface {
	eval() (interface{}, error)
	PrintExp() string
	getVarName() string
	getOperands() []Exp
}

type expNumber struct {
	val float64
}

type expOperator struct {
	opType   TokenType
	operands []Exp
}

// HW5 Update: A new exp for var of function
type expVar struct {
	opType  TokenType
	varName string
}

func (e *expVar) getVarName() string {
	return e.varName
}

func (e *expNumber) getVarName() string {

	return ""
}

func (e *expOperator) getVarName() string {
	return ""
}

func (e *expVar) getOperands() []Exp {
	return nil
}

func (e *expNumber) getOperands() []Exp {
	return nil
}

func (e *expOperator) getOperands() []Exp {
	return e.operands
}

/*
HW5 Update:

	eval() for expVar
	return eval() from map's value of Hashmap_minrkt
*/
func (e *expVar) eval() (interface{}, error) {
	if value, ok := stack.StackHM[len(stack.StackHM)-1].TemporaryMap[e.varName]; ok {
		return value.eval()
	}
	if _, ok := stack.StackHM[len(stack.StackHM)-1].DefineNameMap[e.varName]; ok {
		for _, i := range stack.StackHM[len(stack.StackHM)-1].DefineFuncParser[e.varName].getOperands() {
			if i.getVarName() != "" && i.getVarName() == e.varName {
				return nil, fmt.Errorf("cannot reference an identifier before its definition")
			}
		}
		return stack.StackHM[len(stack.StackHM)-1].DefineFuncParser[e.varName].eval()
	}
	return nil, fmt.Errorf("missing a variable error")
}

func (e *expNumber) eval() (interface{}, error) {
	return e.val, nil
}

/*
HW5 Update:
*/
func (e *expOperator) eval() (interface{}, error) {
	if len(e.operands) == 0 && e.opType == 0 {
		return 0, nil
	}
	if e.opType == TOK_INVALID {
		return 0, fmt.Errorf("TOK_INVALID")
	}
	var res float64
	var v float64
	if e.opType == TOK_DEFINE {
		if len(e.operands) != 2 {
			return 0, fmt.Errorf("2 number operators only")
		}
		funcName := e.operands[0].getVarName()
		if funcName != "" {
			stack.StackHM[len(stack.StackHM)-1].DefineNameMap[funcName] = e.operands[0]
			stack.StackHM[len(stack.StackHM)-1].DefineFuncParser[funcName] = e.operands[1]
		} else {
			stack.StackHM[len(stack.StackHM)-1].DefineNameMap[e.operands[0].getOperands()[0].getVarName()] = e.operands[0]
			stack.StackHM[len(stack.StackHM)-1].DefineFuncParser[e.operands[0].getOperands()[0].getVarName()] = e.operands[1]
		}

		return nil, nil
	} else if e.opType == TOK_DEF_VAR {
		if storedEXP, ok := stack.StackHM[len(stack.StackHM)-1].DefineNameMap[e.operands[0].getVarName()]; ok {
			if len(storedEXP.getOperands()) != len(e.operands) {
				return nil, fmt.Errorf("variable number error")
			}
			// If found a function into recursion
			if _, ok := stack.StackHM[len(stack.StackHM)-1].TemporaryMap[e.operands[0].getVarName()]; ok {
				TemporaryNameMap := CopyMap(stack.StackHM[len(stack.StackHM)-1].DefineNameMap)
				TemporaryFuncParser := CopyMap(stack.StackHM[len(stack.StackHM)-1].DefineFuncParser)
				TemporaryNewMap := CopyMap(stack.StackHM[len(stack.StackHM)-1].TemporaryMap)
				for i := 1; i < len(e.operands); i++ {
					temp, _ := e.operands[i].eval()
					num := expNumber{val: reflect.ValueOf(temp).Interface().(float64)}
					TemporaryNewMap[TemporaryNameMap[e.operands[0].getVarName()].getOperands()[i].getVarName()] = &num
				}

				TemporaryStorage := Hashmap_minrkt{DefineNameMap: TemporaryNameMap, DefineFuncParser: TemporaryFuncParser, TemporaryMap: TemporaryNewMap}
				stack.StackHM = append(stack.StackHM, TemporaryStorage)
				TempExp := Evaluator(TemporaryFuncParser[e.operands[0].getVarName()], stack)
				defer func() {
					stack.StackHM = stack.StackHM[:len(stack.StackHM)-1]
				}()
				return TempExp.eval()
			}
			// not found a function normal loop
			for i := 0; i < len(e.operands); i++ {
				if e.operands[i].getVarName() == "" {
					stack.StackHM[len(stack.StackHM)-1].TemporaryMap[storedEXP.getOperands()[i].getVarName()] = e.operands[i]
					continue
				} else {
					if value, ok := stack.StackHM[len(stack.StackHM)-1].DefineNameMap[e.operands[i].getVarName()]; ok {
						stack.StackHM[len(stack.StackHM)-1].TemporaryMap[storedEXP.getOperands()[i].getVarName()] = value
						continue
					}
					return nil, fmt.Errorf("unknow variable number error")
				}
			}
			return stack.StackHM[len(stack.StackHM)-1].DefineFuncParser[e.getOperands()[0].getVarName()].eval()
		}
		return nil, fmt.Errorf("cannot reference an identifier before its definition")
	} else if e.opType == TOK_BOOL_T || e.opType == TOK_BOOL_F {
		if len(e.operands) != 0 {
			return 0, fmt.Errorf("invalid expression")
		}
		if e.opType == TOK_BOOL_T {
			return true, nil
		} else {
			return false, nil
		}
	} else if e.opType == TOK_COMP_EQ {
		if len(e.operands) == 0 {
			return 0, fmt.Errorf("invalid expression")
		}
		if len(e.operands) == 1 {
			return e.operands[0].eval()
		}
		if len(e.operands) > 2 {
			return 0, fmt.Errorf("2 number operators only")
		}
		temp, err := e.operands[0].eval()
		temp2, err2 := e.operands[1].eval()
		if err != nil || err2 != nil {
			return 0, fmt.Errorf("invalid expression")
		}
		if reflect.TypeOf(temp) == reflect.TypeOf(temp2) {
			return reflect.ValueOf(temp).Interface() ==
				reflect.ValueOf(temp2).Interface(), nil
		}
		return 0, fmt.Errorf("invalid expression")
	} else if e.opType == TOK_COMP_GE || e.opType == TOK_COMP_G || e.opType == TOK_COMP_S || e.opType == TOK_COMP_SE {
		if len(e.operands) != 2 {
			return 0, fmt.Errorf("2 number operators only")
		}
		temp, err := e.operands[0].eval()
		t := reflect.TypeOf(temp)
		if err != nil || t.String() != "float64" {
			return 0, fmt.Errorf("invalid expression")
		}
		temp2, err2 := e.operands[1].eval()
		t2 := reflect.TypeOf(temp2)
		if err2 != nil || t2.String() != "float64" {
			return 0, fmt.Errorf("invalid expression")
		}
		if e.opType == TOK_COMP_GE {
			return reflect.ValueOf(temp).Interface().(float64) >=
				reflect.ValueOf(temp2).Interface().(float64), nil
		} else if e.opType == TOK_COMP_G {
			return reflect.ValueOf(temp).Interface().(float64) >
				reflect.ValueOf(temp2).Interface().(float64), nil
		} else if e.opType == TOK_COMP_SE {
			return reflect.ValueOf(temp).Interface().(float64) <=
				reflect.ValueOf(temp2).Interface().(float64), nil
		} else if e.opType == TOK_COMP_S {
			return reflect.ValueOf(temp).Interface().(float64) <
				reflect.ValueOf(temp2).Interface().(float64), nil
		}
		return 0, fmt.Errorf("invalid expression")
	} else if e.opType == TOK_COND_AND || e.opType == TOK_COND_OR {
		if len(e.operands) != 2 {
			return 0, fmt.Errorf("invalid expression")
		}
		temp, err := e.operands[0].eval()
		t := reflect.TypeOf(temp)
		if err != nil || t.String() != "bool" {
			return 0, fmt.Errorf("invalid expression")
		}
		if reflect.ValueOf(temp).Interface().(bool) {
			if e.opType == TOK_COND_OR {
				return true, nil
			}
		} else {
			if e.opType == TOK_COND_AND {
				return false, nil
			}
		}
		temp2, err2 := e.operands[1].eval()
		t2 := reflect.TypeOf(temp2)
		if err2 != nil || t2.String() != "bool" {
			return 0, fmt.Errorf("invalid expression")
		}
		if e.opType == TOK_COND_AND {
			return reflect.ValueOf(temp2).Interface().(bool), nil
		} else if e.opType == TOK_COND_OR {
			return reflect.ValueOf(temp2).Interface().(bool), nil
		}
		return 0, fmt.Errorf("invalid expression")
	} else if e.opType == TOK_COND_NOT {
		if len(e.operands) != 1 {
			return 0, fmt.Errorf("invalid expression")
		}
		temp, err := e.operands[0].eval()
		t := reflect.TypeOf(temp)
		if err != nil || t.String() != "bool" {
			return 0, fmt.Errorf("invalid expression")
		}
		return !reflect.ValueOf(temp).Interface().(bool), nil
	} else if e.opType == TOK_COND_IF {
		if len(e.operands) != 3 {
			return 0, fmt.Errorf("invalid expression")
		}
		temp, err := e.operands[0].eval()
		t := reflect.TypeOf(temp)
		if err != nil || t.String() != "bool" {
			return 0, fmt.Errorf("invalid expression")
		}
		if reflect.ValueOf(temp).Interface().(bool) {

			return e.operands[1].eval()
		} else {

			return e.operands[2].eval()
		}
	} else if e.opType == TOK_ADD {
		for _, i := range e.operands {
			temp, err := i.eval()
			if err != nil {
				return 0, err
			}
			if t := reflect.TypeOf(temp).String() == "float64"; t {
				v = reflect.ValueOf(temp).Interface().(float64)
			}

			res += v
		}
	} else if e.opType == TOK_SUB {
		for _, i := range e.operands {
			temp, err := i.eval()
			if err != nil {
				return 0, err
			}
			if t := reflect.TypeOf(temp).String() == "float64"; t {
				v = reflect.ValueOf(temp).Interface().(float64)
			}
			// assign the value to result if result is declared but not defined
			// This step is necessary for SUB MUL and DIV
			if res == 0 {
				res = v
				continue
			}
			res -= v
		}
	} else if e.opType == TOK_MUL {
		for _, i := range e.operands {
			temp, err := i.eval()
			if err != nil {
				return 0, err
			}
			if t := reflect.TypeOf(temp).String() == "float64"; t {
				v = reflect.ValueOf(temp).Interface().(float64)
			}
			// assign the value to result if result is declared but not defined
			// This step is necessary for SUB MUL and DIV
			if res == 0 {
				res = v
				continue
			}
			res *= v
		}
	} else if e.opType == TOK_DIV {
		for _, i := range e.operands {
			temp, err := i.eval()
			if err != nil {
				return 0, err
			}
			if t := reflect.TypeOf(temp).String() == "float64"; t {
				v = reflect.ValueOf(temp).Interface().(float64)
			}
			// assign the value to result if result is declared but not defined
			// This step is necessary for SUB MUL and DIV
			if res == 0 {
				res = v
				continue
			}
			if v == 0 {
				return 0, fmt.Errorf("cannot be divided by 0")
			}
			res /= v
		}
	}
	return res, nil
}

/*
printExp helps print exp nicely
*/
func (e *expNumber) PrintExp() string {
	return fmt.Sprintf("%f ", e.val)
}

func (e *expOperator) PrintExp() string {
	res := "NodeTokenType: (" + PrintTokenType(int(e.opType)) + " : "
	for _, i := range e.operands {
		res += i.PrintExp()
	}
	res += ") "
	return res
}

func (e *expVar) PrintExp() string {
	return fmt.Sprintf("%s ", e.varName)
}

func CopyMap(m map[string]Exp) map[string]Exp {
	cp := make(map[string]Exp)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}
