package step

import (
	"fmt"
)

const (
	NodeStateCreated  = "Created"
	NodeStateRunning  = "Running"
	NodeStateWaiting  = "Waiting"
	NodeStateFinished = "Finished"
	NodeStateFailed   = "Failed"
)

type (
	NodeState string
	Response  struct {
		Index     int
		Succeeded bool
	}

	Node struct {
		UUID      string
		Index     int
		Value     Value
		Children  []*Node
		Father    *Node
		End       bool
		Signal    chan *Response
		State     NodeState
		All       int
		Succeeded int
		Failed    int
	}
)

func New(value Value, father *Node) *Node {
	n := &Node{
		UUID:   "",
		Value:  value,
		Father: father,
		Signal: make(chan *Response, 1),
		State:  NodeStateCreated,
	}

	if father != nil {
		father.Children = append(father.Children, n)
		father.All++
		n.Index = len(father.Children) - 1
	}

	return n
}

func read(result []*Node, n *Node) {
	result = append(result, n)
	if n.End {
		results = append(results, result)
	} else {
		result2 := make([]*Node, len(result))
		copy(result2, result)
		for i := range n.Children {
			if n.Children[i] != nil {
				read(result2, n.Children[i])
			}
		}
	}
}

var results = make([][]*Node, 0)

func (n *Node) Show(step bool) {
	// results := make([][]*Node, 0)
	result := make([]*Node, 0)
	read(result, n)
	for _, nodes := range results {
		nodes[len(nodes)-1].Value.Print()
		if step {
			for _, node := range nodes {
				node.Value.Print()
			}
		}
		fmt.Printf("\n\n")
	}
}

func (n *Node) Run() {
	n.State = NodeStateRunning

	if n.Value.Finished() {
		n.End = true
		n.finish()
		return
	}

	values := n.Value.NextStep()
	if len(values) == 0 {
		n.boom()
		return
	}

	for _, value := range values {
		node := New(value, n)
		go node.Run()
	}

	n.wait()
	if n.Failed == len(n.Children) {
		n.boom()
		return
	}
	n.finish()
}

func (n *Node) wait() {
	n.State = NodeStateWaiting

	for i := 0; i < len(n.Children); i++ {
		resp := <-n.Signal
		if resp.Succeeded {
			n.Succeeded++
		} else {
			n.Failed++
			n.Children[resp.Index] = nil
		}
	}
}

func (n *Node) boom() {
	n.State = NodeStateFailed

	if n.Father != nil {
		n.Father.Signal <- &Response{
			Succeeded: false,
			Index:     n.Index,
		}
	}
}

func (n *Node) finish() {
	n.State = NodeStateFinished

	if n.Father != nil {
		n.Father.Signal <- &Response{
			Succeeded: true,
			Index:     n.Index,
		}
	}
}
