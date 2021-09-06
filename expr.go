package expr

import (
	"bufio"
	"bytes"
	"container/list"
	"errors"
)

const (
	NumberKind = iota
	PlusKind
	MinusKind
	MultiplyKind
	ConstKind
	IdentifierKind
)

type Node struct {
	Identifier byte
	Kind       int
	Value      int
	Priority   int
}

const (
	Whitespace byte = 32
	Multiply   byte = 42
	Plus       byte = 43
	Minus      byte = 45
	C          byte = 99
	O          byte = 111
	N          byte = 110
	S          byte = 115
	T          byte = 116
	EQUAL      byte = 61
)

var constWord = []byte{C, O, N, S, T}

func Parse(text []byte) *list.List {

	scanner := bufio.NewScanner(bytes.NewReader(text))

	nodes := list.New()
outerLoop:
	for scanner.Scan() {
		expression := scanner.Bytes()

		for i := 0; i < len(expression); i++ {
			if expression[i] == C {
				n, err := parseConst(expression)
				if err != nil {
					panic(err)
				}

				Env.Consts[n.Identifier] = n
				continue outerLoop
			}
		}

		for i := len(expression) - 1; i >= 0; i-- {
			if expression[i] == Minus {
				nodes.PushBack(Node{Kind: MinusKind})
			} else if expression[i] == Multiply {
				nodes.PushBack(Node{Kind: MultiplyKind, Priority: 1})
			} else if expression[i] == Plus {
				nodes.PushBack(Node{Kind: PlusKind})
			} else if expression[i] >= 48 && expression[i] <= 57 {
				n := Node{Kind: NumberKind}
				var sum int
				sum = int(expression[i]) - 48
				i--
				norm := 10
				for i >= 0 {
					if expression[i] >= 48 && expression[i] <= 57 {
						sum += (int(expression[i]) - 48) * norm
						i--
						norm = norm * 10
					} else {
						i++
						break
					}
				}
				n.Value = sum
				nodes.PushBack(n)
			} else if expression[i] >= 97 && expression[i] <= 122 {
				nodes.PushBack(Node{Kind: IdentifierKind, Identifier: expression[i]})
			}
		}
	}

	return nodes
}

func parseConst(value []byte) (Node, error) {
	i := 0
	for ; i < len(constWord); i++ {
		if !peakNextChar(value, i, constWord[i]) {
			return Node{}, errors.New("const is bad (CONST)")
		}
	}
	i++
	identifier := value[i]
	i++
	i++
	if !peakNextChar(value, i, EQUAL) {
		return Node{}, errors.New("const is bad (EQUAL)")
	}
	i++
	i++

	sum := int(value[len(value)-1]) - 48

	norm := 10
	for j := len(value) - 2; j >= i; j-- {
		if value[j] >= 48 && value[j] <= 57 {
			sum += (int(value[j]) - 48) * norm
			j--
			norm = norm * 10
		}
	}

	return Node{
		Kind:       ConstKind,
		Value:      sum,
		Identifier: identifier,
	}, nil
}

func peakNextChar(value []byte, i int, char byte) bool {
	if value[i] == char {
		return true
	}
	return false
}

type Environment struct {
	Consts map[byte]Node
}

var Env Environment = Environment{
	Consts: map[byte]Node{},
}

func Eval(nodes *list.List) int {
	if nodes.Len() == 0 {
		return 0
	}

	for e := nodes.Back(); e != nil; e = e.Prev() {
		node := e.Value.(Node)
		if node.Priority == 1 {
			if node.Kind == MultiplyKind {
				eA := e.Prev()
				nodeA := eA.Value.(Node)
				eb := e.Next()
				nodeB := eb.Value.(Node)
				sum := nodeA.Value * nodeB.Value

				nodes.Remove(eb)
				nodes.Remove(eA)
				newNode := nodes.InsertBefore(Node{Kind: NumberKind, Value: sum}, e)
				nodes.Remove(e)
				e = newNode
			}
		}
	}

	value := nodes.Back()
	sum := value.Value.(Node).Value
	value = value.Prev()
	for e := value; e != nil; e = e.Prev() {
		node := e.Value.(Node)
		if node.Kind == PlusKind {
			e = e.Prev()
			node = e.Value.(Node)

			if node.Kind == IdentifierKind {
				valueNode, ok := Env.Consts[node.Identifier]
				if !ok {
					panic("value is not set")
				}
				sum += valueNode.Value
			} else {
				sum += node.Value
			}

		} else if node.Kind == MinusKind {
			e = e.Prev()
			node = e.Value.(Node)
			if node.Kind == IdentifierKind {
				valueNode, ok := Env.Consts[node.Identifier]
				if !ok {
					panic("value is not set")
				}
				sum -= valueNode.Value
			} else {
				sum -= node.Value
			}
		}
	}
	return sum
}
