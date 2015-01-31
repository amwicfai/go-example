/*结构体没有构造和析构函数，以普通函数代替*/

package main

import (
	"fmt"
)

type TestResult struct {
	Name string
}

func NewTestResult(name string) *TestResult {
	tr := TestResult{}
	tr.Name = name
	return &tr
}

func (tr *TestResult) Print() {
	fmt.Printf("Name:%s", tr.Name)
}

func main() {
	tr := NewTestResult("Insert")
	tr.Print()
}
