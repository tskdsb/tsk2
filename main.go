/*
 * Copyright 2017 caicloud authors. All rights reserved.
 */

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tskdsb/tsk2/example/sudoku"
	"github.com/tskdsb/tsk2/pkg/step"
)

func main1() {

	a, b, err := sudoku.GetPair()
	if err != nil {
		log.Fatalf("get pair err: %s\n", err)
	}

	node := step.New(a, nil)
	node.Value.Print()

	now := time.Now()
	node.Run()
	fmt.Printf("%s\n\n", time.Now().Sub(now))

	node.Show(false)
	b.Print()
}

type Tsk struct {
	A int
	B string
}

func Defer() (t *Tsk) {
	// defer func() {
	// 	fmt.Println(t)
	// }()
	return &Tsk{
		A: 1,
		B: "2",
	}
}

func main() {
	a:=map[string]*Tsk{}
a["aa"]=&Tsk{
	A: 1,
	B: "aa",
}

	for _, tsk := range a {
		tsk.A=10
	}
	fmt.Println(a["aa"].A)

	fmt.Printf("%s\n", "OK")
}
