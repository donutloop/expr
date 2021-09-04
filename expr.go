package expr

import (
	"container/list"
)

const (
	NumberKind = iota
	PlusKind
	MinusKind
	MultiplyKind
)

type Node struct {
	Kind     int
	Value    int
	Priority int
}

const (
	Whitespace byte = 32
	Multiply   byte = 42
	Plus       byte = 43
	Minus      byte = 45
)

func Parse(expression []byte) *list.List {

	nodes := list.New()
	cleanedExpression := make([]byte, 0)
	for i := 0; i < len(expression); i++ {
		if expression[i] == Whitespace {
			continue
		}
		cleanedExpression = append(cleanedExpression, expression[i])
	}

	for i := len(cleanedExpression) - 1; i >= 0; i-- {
		if cleanedExpression[i] == Minus {
			nodes.PushBack(Node{Kind: MinusKind})
		} else if cleanedExpression[i] == Multiply {
			nodes.PushBack(Node{Kind: MultiplyKind, Priority: 1})
		} else if cleanedExpression[i] == Plus {
			nodes.PushBack(Node{Kind: PlusKind})
		} else if cleanedExpression[i] >= 48 && cleanedExpression[i] <= 57 {
			n := Node{Kind: NumberKind}
			var sum int
			sum = int(cleanedExpression[i]) - 48
			i--
			norm := 10
			for i >= 0 {
				if cleanedExpression[i] >= 48 && cleanedExpression[i] <= 57 {
					sum += (int(cleanedExpression[i]) - 48) * norm
					i--
					norm = norm * 10
				} else {
					i++
					break
				}
			}
			n.Value = sum
			nodes.PushBack(n)
		}
	}

	return nodes
}

func Eval(nodes *list.List) int {
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
			sum += node.Value
		} else if node.Kind == MinusKind {
			e = e.Prev()
			node = e.Value.(Node)
			sum -= node.Value
		}
	}
	return sum
}
