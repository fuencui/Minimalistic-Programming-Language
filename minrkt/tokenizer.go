package minrkt

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
global variables
HW5 Update: two more token Regex
*/
var tokenRegexList = []string{
	`^(\()`,
	`^(\))`,
	`^((?:\-)*[0]|[1-9][0-9]*(?:\.[0-9]*)?)`,
	`^(\+)`,
	`^(\-)`,
	`^(\*)`,
	`^(\/)`,
	`^(\s+)`,
	`^(\=)`,
	`^(\>=)`,
	`^(\>)`,
	`^(\<=)`,
	`^(\<)`,
	`^(true|#t)`,
	`^(false|#f)`,
	`^(if)`,
	`^(and)`,
	`^(or)`,
	`^(not)`,
	`^(define$|define )`,
	`^([a-zA-Z][a-zA-Z0-9_-]*)`,
}

var re = regexp.MustCompile(strings.Join(tokenRegexList, "|"))

/*
TokenType struct refactor
HW5 Update: new var DefineName (string)
*/
type Token struct {
	TokenType  TokenType
	NumValue   float64
	DefineName string
}

type TokenType int

// TokenType enum
// HW5 Update: TOK_DEFINE, TOK_DEF_VAR
const (
	TOK_INVALID TokenType = iota
	TOK_LPAREN
	TOK_RPAREN
	TOK_NUM
	TOK_ADD
	TOK_SUB
	TOK_MUL
	TOK_DIV
	TOK_WSP
	TOK_COMP_EQ
	TOK_COMP_GE
	TOK_COMP_G
	TOK_COMP_SE
	TOK_COMP_S
	TOK_BOOL_T
	TOK_BOOL_F
	TOK_COND_IF
	TOK_COND_AND
	TOK_COND_OR
	TOK_COND_NOT
	TOK_DEFINE
	TOK_DEF_VAR
)

/*
A helper function to generate a token
HW5 Update: Identify DefineName
*/
func GenerateToken(index int, s string) (Token, error) {
	if index == 3 {
		floatRes, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return Token{}, err
		}
		return Token{TokenType: TokenType(index), NumValue: floatRes}, nil
	} else if index == 21 {
		return Token{TokenType: TokenType(index), DefineName: s}, nil
	} else {
		return Token{TokenType: TokenType(index)}, nil
	}
}

/*
Tokenizer enter function
*/
func Tokenizer(line string) ([]Token, error) {
	remainder := line
	var tokens []Token
	for {
		token, newRemainder, err := NextToken(remainder)
		if err != nil {
			return nil, fmt.Errorf("invalid token: %v", err)
		}
		if token.TokenType != TOK_WSP {
			tokens = append(tokens, token)
		}
		if len(newRemainder) == 0 {
			//fmt.Printf("[RESULT] Tokenizer final output:  %v\n", tokens)
			return tokens, nil
		}
		remainder = newRemainder
	}
}

/*
Token logic function
HW5 Update: Identify define function name or var name
*/
func NextToken(remainder string) (Token, string, error) {
	outputString := re.FindStringSubmatch(remainder)
	var regexIndex int
	var rawString string
	for i, temp := range outputString {
		if temp != "" {
			regexIndex = i
			rawString = temp
		}
	}

	var newRemainder string
	if rawString == "(" || rawString == ")" {
		newRemainder = remainder[1:]
	} else if rawString == "define" || rawString == "define " {
		newRemainder = remainder[6:]
	} else {
		newRemainder = strings.TrimLeft(remainder, rawString)
	}

	ResultToken, err := GenerateToken(regexIndex, rawString)
	if err != nil {
		return Token{}, "", err
	}
	if ResultToken.TokenType == TOK_INVALID {
		return Token{}, "", fmt.Errorf("invalid token")
	}
	return ResultToken, newRemainder, nil
}

// print nicely only for debug
var TokenStringArray = [...]string{
	"TOK_INVALID",
	"TOK_LPAREN",
	"TOK_RPAREN",
	"TOK_NUM",
	"TOK_ADD",
	"TOK_SUB",
	"TOK_MUL",
	"TOK_DIV",
	"TOK_WSP",
	"TOK_COMP_EQ",
	"TOK_COMP_GE",
	"TOK_COMP_G",
	"TOK_COMP_SE",
	"TOK_COMP_S",
	"TOK_BOOL_T",
	"TOK_BOOL_F",
	"TOK_COND_IF",
	"TOK_COND_AND",
	"TOK_COND_OR",
	"TOK_COND_NOT",
	"TOK_DEFINE",
	"TOK_DEF_VAR",
}

// print nicely only for debug
func PrintTokenType(index int) string {
	return TokenStringArray[index]
}
