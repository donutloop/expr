package expr

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		expression string
		value int
	}{
		{
			name: "simple expression",
			expression: "9 + 9",
			value: 18,
		},
		{
			name: "simple expression",
			expression: "50 + 50",
			value: 100,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nodes := Parse([]byte(test.expression))
			sum := Eval(nodes)
			t.Logf("sum: %#v", sum)
			if test.value != sum {
				t.Fatalf("sum is bad, got:%v, want:%v", sum, test.value)
			}
		})
	}
}
