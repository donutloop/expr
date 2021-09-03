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
			expression: "9 - 9",
			value: 0,
		},
		{
			name: "multiple expression",
			expression: "9 - 9 - 9",
			value: -9,
		},
		{
			name: "simple expression",
			expression: "9 + 9",
			value: 18,
		},
		{
			name: "multiple expression",
			expression: "9 + 9 + 9",
			value: 27,
		},
		{
			name: "multiple expression",
			expression: "9 + 9 + 9 + 9",
			value: 36,
		},
		{
			name: "simple expression",
			expression: "50 + 50",
			value: 100,
		},
		{
			name: "simple expression",
			expression: "500 + 500",
			value: 1000,
		},
		{
			name: "simple expression",
			expression: "50000000000000 + 50000000000000",
			value: 100000000000000,
		},
		{
			name: "mixture expression",
			expression: "9 - 9 - 9 + 9 + 9 + 9",
			value: 18,
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
