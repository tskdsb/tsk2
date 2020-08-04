package sudoku

import (
	"fmt"

	"github.com/tskdsb/tsk2/pkg/step"
)

type Value [9][9]int

func New() *Value {
	return &Value{}
}

func (v *Value) deepCopy() *Value {
	v2 := &Value{}
	for x := range v {
		for y := range v[x] {
			if v[x][y] != 0 {
				v2[x][y] = v[x][y]
			}
		}
	}

	return v2
}

func (v *Value) NextStep() []step.Value {
	var x, y int

A:
	for x = range v {
		for y = range v[x] {
			if v[x][y] == 0 {
				break A
			}
		}
	}

	m := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
		9: {},
	}

	for i := 0; i < 9; i++ {
		if i == y {
			continue
		}
		if _, ok := m[v[x][i]]; ok {
			delete(m, v[x][i])
		}
	}

	for i := 0; i < 9; i++ {
		if i == x {
			continue
		}
		if _, ok := m[v[i][y]]; ok {
			delete(m, v[i][y])
		}
	}

	x0 := x / 3 * 3
	y0 := y / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (x0+i) == x && (y0+j) == y {
				continue
			}
			if _, ok := m[v[x0+i][y0+j]]; ok {
				delete(m, v[x0+i][y0+j])
			}
		}
	}

	result := make([]step.Value, 0)
	for s := range m {
		v2 := v.deepCopy()
		v2[x][y] = s
		result = append(result, v2)
	}

	return result
}

func (v *Value) Finished() bool {
	for _, a := range v {
		for _, b := range a {
			if b == 0 {
				return false
			}
		}
	}

	return true
}
func (v *Value) Print() {
	for _, a := range v {
		fmt.Println(a)
	}
	fmt.Println()
}
