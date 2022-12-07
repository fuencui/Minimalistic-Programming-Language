package minrkt

import (
	"fmt"
)

/*
Parser enter function
HW5 Update:

	EXP,Eval and other function are integrated into evaluator
*/
func Parser(tokens []Token) Exp {
	if err := lparen_checker(tokens); err != nil {
		fmt.Println("invalid expression")
		return nil
	}

	// HW4 Update: new parser function return the remainder of unconsumed tokens mentioned in fuen's #PR
	root2, remainToken, err := CompleteParser(tokens, expOperator{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(remainToken) != 0 {
		fmt.Println("invalid expression")
		return nil
	}
	return root2
}

/*
new completeParser function
Take advice you mentioned in PR
builder a parser tree recursively
instead of update i, return unconsumed tokens
it will be recognized invalid expression
HW5 Update:

	Identify Define function Name
	Build Parser with new token type of define or define var
*/
func CompleteParser(tokens []Token, root expOperator) (Exp, []Token, error) {
	if len(tokens) == 0 {
		return nil, nil, nil
	}
	if len(tokens) == 1 && tokens[0].TokenType == TOK_NUM {
		return &expNumber{val: tokens[0].NumValue}, nil, nil
	} else if len(tokens) == 1 && tokens[0].TokenType == TOK_BOOL_F {
		return &expOperator{opType: TOK_BOOL_F}, nil, nil
	} else if len(tokens) == 1 && tokens[0].TokenType == TOK_BOOL_T {
		return &expOperator{opType: TOK_BOOL_T}, nil, nil
	} else if len(tokens) == 1 && tokens[0].TokenType == TOK_DEF_VAR {
		return &expVar{opType: tokens[0].TokenType, varName: tokens[0].DefineName}, nil, nil
	}

	for len(tokens) != 0 {
		if tokens[0].TokenType == TOK_LPAREN {
			if root.opType == 0 {
				tokens = tokens[1:]
				continue
			} else {
				node, updateToken, err := CompleteParser(tokens[1:], expOperator{})
				if err != nil {
					return nil, nil, err
				}
				root.operands = append(root.operands, node)
				tokens = updateToken
			}
		} else if tokens[0].TokenType == TOK_NUM {
			if root.opType == TOK_NUM {
				return nil, nil, fmt.Errorf("invalid expression")
			} else if root.opType == 0 {
				return nil, nil, fmt.Errorf("invalid expression")
			} else {
				root.operands = append(root.operands, &expNumber{val: tokens[0].NumValue})
				tokens = tokens[1:]
				continue
			}
		} else if tokens[0].TokenType != TOK_RPAREN {
			if root.opType == 0 {
				root.opType = tokens[0].TokenType
				if tokens[0].TokenType == TOK_DEF_VAR {
					root.operands = append(root.operands, &expVar{opType: tokens[0].TokenType, varName: tokens[0].DefineName})
				}
				tokens = tokens[1:]
				continue
			} else if tokens[0].TokenType == TOK_DEF_VAR {
				root.operands = append(root.operands, &expVar{opType: tokens[0].TokenType, varName: tokens[0].DefineName})
				tokens = tokens[1:]
				continue
			} else {
				root.operands = append(root.operands, &expOperator{opType: tokens[0].TokenType})
				tokens = tokens[1:]
				continue
			}
		} else if tokens[0].TokenType == TOK_RPAREN {
			return &root, tokens[1:], nil
		}
	}
	return nil, nil, fmt.Errorf("invalid expression")
}

/*
HW4 Update: a Parser checker
*/
func lparen_checker(tokens []Token) error {
	var lparen_checker int
	for _, i := range tokens {
		if i.TokenType == TOK_LPAREN {
			lparen_checker++
		} else if i.TokenType == TOK_RPAREN {
			lparen_checker--
		}
	}
	if lparen_checker != 0 {
		return fmt.Errorf("invalid expression")
	}
	return nil
}

//legacy code from HW3

// func inCompleteParser(tokens []Token, pointer int) (Exp, int, error) {
// 	// edge cases
// 	if len(tokens) == 0 {
// 		return nil, 0, nil
// 	}
// 	if tokens[0].TokenType == TOK_NUM {
// 		return &expNumber{val: tokens[0].NumValue}, 0, nil
// 	} else if len(tokens) == 1 && (tokens[0].TokenType == TOK_BOOL_F || tokens[0].TokenType == TOK_BOOL_T) {
// 		return &expOperator{opType: tokens[0].TokenType}, 0, nil
// 	} else if tokens[0].TokenType != TOK_LPAREN {
// 		return nil, 0, fmt.Errorf("invalid expression1")
// 	}
// 	// main logic loop
// 	root := expOperator{}
// 	for i := pointer; i < len(tokens); i = i + 1 {
// 		if i == 0 {
// 			continue
// 		}
// 		if tokens[i].TokenType == TOK_NUM {
// 			root.operands = append(root.operands, &expNumber{val: tokens[i].NumValue})
// 		} else if tokens[i].TokenType == TOK_LPAREN {
// 			// recursion call
// 			newNode, updatePointer, err := inCompleteParser(tokens, i+1)
// 			if err != nil {
// 				return nil, i, err
// 			}
// 			root.operands = append(root.operands, newNode)
// 			if updatePointer != len(tokens)-1 {
// 				i = updatePointer
// 			} else {
// 				return nil, 0, fmt.Errorf("invalid expression2")
// 			}
// 		} else if tokens[i].TokenType == TOK_RPAREN {
// 			return &root, i, nil
// 		} else {
// 			//youwenti
// 			if tokens[i-1].TokenType == TOK_LPAREN ||
// 				tokens[i-1].TokenType == TOK_BOOL_F ||
// 				tokens[i-1].TokenType == TOK_BOOL_T ||
// 				tokens[i-1].TokenType == TOK_COND_IF ||
// 				tokens[i-1].TokenType == TOK_COND_NOT ||
// 				tokens[i-1].TokenType == TOK_COND_OR {
// 				root.opType = tokens[i].TokenType
// 			} else {
// 				return nil, 0, fmt.Errorf("invalid expression3")
// 			}
// 		}
// 	}

// 	// check invalid expression by the end of recursion call
// 	if tokens[len(tokens)-1].TokenType != TOK_RPAREN {
// 		return nil, 0, fmt.Errorf("invalid expression4")
// 	}

// 	return &root, 0, nil
// }
