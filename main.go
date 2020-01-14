/*
 * Copyright 2017 caicloud authors. All rights reserved.
 */

package main

import (
	"fmt"

	"github.com/tskdsb/tsk2/pkg/aaa"
	"github.com/tskdsb/tsk2/pkg/bbb"
	"github.com/tskdsb/tsk2/pkg/ccc"
	"k8s.io/client-go/tools/pager"
)

func main() {
	aaa.PrintTest()
	bbb.PrintTest()
	fmt.Printf("%+v\n", pager.ListPager{})
	fmt.Printf("%+v\n", ccc.Ccc("xxx"))
}
