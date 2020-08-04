package spidercard

import (
	"fmt"

	"github.com/mohae/deepcopy"

	"github.com/tskdsb/tsk2/pkg/step"
)

type (
	Card struct {
		N int
		// r: red
		// g: green
		// b: black
		Color   byte
		IsZhong bool
		IsFa    bool
		IsBai   bool
		IsHua   bool
	}

	Value struct {
		place  [3]*Card
		flower *Card
		sorted [3][]*Card
		list   [8][]*Card
	}
)

func (c *Card) String() string {
	if c.IsZhong {
		return "ZZ"
	}
	if c.IsFa {
		return "FF"
	}
	if c.IsBai {
		return "BB"
	}
	if c.IsHua {
		return "HH"
	}
	return fmt.Sprintf("%c%d", c.Color, c.N)
}

func (v *Value) Finished() bool {
	v2 := v.deepCopy()
	for {
		hit := 0
	FOR1:
		for i1, cards1 := range v2.sorted {
			card := cards1[len(cards1)-1]
			for i2, cards2 := range v2.list {
				if len(cards2) == 0 {
					continue
				}
				card2 := cards2[len(cards2)-1]
				if card2.Color == card.Color &&
					card2.N == card.N+1 {
					v2.sorted[i1] = append(v2.sorted[i1], card2)
					v2.list[i2] = v2.list[i2][:len(cards2)-1]
					hit++
					continue FOR1
				}
			}
		}
		if hit == 0 {
			break
		}
	}

	for _, cards2 := range v2.list {
		if len(cards2) > 0 {
			return false
		}
	}

	return true
}

func (v *Value) NextStep() []step.Value {
// v2:=v.deepCopy()



	return nil
}

func (v *Value) Print() {
	fmt.Printf("%v %v %v\n", v.place, v.flower, v.sorted)
	fmt.Printf("%v\n\n", v.list)
}

func (v *Value) deepCopy() *Value {
	return deepcopy.Copy(v).(*Value)
}
