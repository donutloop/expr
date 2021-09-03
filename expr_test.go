package expr

import "testing"

func TestParse(t *testing.T) {
	value := "9 + 9"

	nodes := Parse([]byte(value))

	t.Logf("%#v", nodes)

}
