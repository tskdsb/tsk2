package aaa

import (
  "fmt"

  "k8s.io/client-go/tools/pager"
)

func init() {
  fmt.Println("aaa")
}

func PrintTest() {
  fmt.Printf("%+v\n", pager.ListPager{})
}
