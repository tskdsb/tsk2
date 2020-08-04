package ccc

import (
	"fmt"

	v1 "github.com/tskdsb/tsk2/pkg/pkg1/v1"
)

type Ccc string



type Ct1 v1.T1

func (receiver *Ct1) Name() {
	fmt.Println(receiver.A)
}

