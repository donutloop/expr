package expr

const (
	NumberKind = iota
	PlusKind
)

type Node struct {
	Kind  int
	Value int
}

const (
	Whitespace byte = 32
	Plus       byte = 43
)

func Parse(expression []byte) []Node {

	cleanedExpression := make([]byte, 0)
	for i := 0; i < len(expression); i++ {
		if expression[i] == Whitespace {
			continue
		}
		cleanedExpression = append(cleanedExpression, expression[i])
	}

	var nodes []Node
	for i := len(cleanedExpression) - 1; i >= 0; i-- {
		if cleanedExpression[i] == Plus {
			nodes = append(nodes, Node{Kind: PlusKind})
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
			nodes = append(nodes, n)
		}
	}

	return nodes
}

func Eval(nodes []Node) int {
	var sum int
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Kind == PlusKind {
			sum += nodes[i-1].Value + nodes[i+1].Value
			i++
		}
	}
	return sum
}
