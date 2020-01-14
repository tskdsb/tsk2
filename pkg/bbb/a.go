package bbb

import (
	"fmt"

	"k8s.io/client-go/tools/pager"
)

func init() {
	fmt.Println("bbb")
}

func PrintTest() {
	fmt.Printf("%+v\n", pager.ListPager{})
}
