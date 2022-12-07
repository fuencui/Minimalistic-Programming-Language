package minrkt

/*
HW5 Update:

	A new struct for HW5 stored in "environment"
	DefineNameMap map:
		key: function define name
		value: function variables
	DefineFuncParser：
		key: function define name
		value: function parser tree
	TemporaryMap:
		key: function define name
		value: temporary variables value will be delete each loop
*/
type Hashmap_minrkt struct {
	DefineNameMap    map[string]Exp
	DefineFuncParser map[string]Exp
	TemporaryMap     map[string]Exp
}

/*
HW5 Update:

	A Hashmap_minrkt stack can be extende for more other features
	For now just for recursive tracking
*/
type Stack struct {
	StackHM []Hashmap_minrkt
}
