package main

import (
	"errors"
	"fmt"
	"math"
)

const s string = "I'm Isham"

func typeChange() {
	const s string = "Im Isham inside change"
	fmt.Println(s)
}

//////////////////////////////////////
func typeCheck() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(3000)
}

//////////////////////////////////////
func arrayCheck() {
	var a [1]int
	fmt.Println("emp:", a)
}

//////////////////////////////////////
func sliceCheck() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "Zero"
	s[1] = "One"
	s[2] = "Two"
	s = append(s, "three")
	fmt.Println("s:", s)

	c := make([]string, len(s))
	copy(c, s)
	c = append(c, "Five")
	fmt.Println("s:", s)
	fmt.Println("c:", c)

}

//////////////////////////////////////
func mapCheck() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	fmt.Println("map:", m)
	delete(m, "two")
	apr, prs := m["two"]
	fmt.Println("key:", apr)
	fmt.Println("prs:", prs)
}

//////////////////////////////////////
func rangeCheck() {
	nums := []string{"first", "second", "third", "fourth", "fifth"}
	for i, num := range nums {
		fmt.Println("number : ", num)
		fmt.Println("index : ", i)
	}

	for i, c := range "ABCDEFGH" {
		fmt.Println(i, c)
	}
}

//////////////////////////////////////
func addThree(a, b, c int) int {
	return a + b + c
}
func funcCheck() {
	var a, b, c int = 3, 5, 7
	fmt.Println(addThree(a, b, c))
}

//////////////////////////////////////
func printPars(nums ...int) {
	for _, r := range nums {
		fmt.Println(r)
	}
}
func variardicCheck() {
	pars := []int{3, 5, 7, 4, 9}
	printPars(pars...)
}

//////////////////////////////////////
func getSeq() func() int {
	n := 5
	return func() int {
		n++
		return n
	}

}
func closureCheck() {
	nextNum := getSeq()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())
}

//////////////////////////////////////
func facto(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println("Recursive call No. ", n)
	return n * facto(n-1)
}
func recursionCheck() {
	fmt.Println(facto(8))
}

//////////////////////////////////////
func changeVal(i int) {
	i = 0
}
func changePtr(i *int) {
	*i = 0
}
func pointerCheck() {
	var b = 345
	fmt.Println("Value of ib: ", b)
	changeVal(b)
	fmt.Println("Value after changing b : ", b)
	changePtr(&b)
	fmt.Println("Value after changing addresss of b : ", b)
	fmt.Println("Value of the address of b  ", &b)

}

//////////////////////////////////////
type rectangle struct {
	length  int
	breadth int
}

func structCheck() {
	f := rectangle{length: 30, breadth: 20}
	fmt.Println("rectangle details : ", f)
	fb := &f
	fb.breadth = 25
	fmt.Println("rectangle details : ", f)
}

//////////////////////////////////////
func (r rectangle) area() float64 {
	return float64(r.length * r.breadth)
}
func (r rectangle) perimeter() float64 {
	return float64(2*r.length + 2*r.breadth)
}

func (r *rectangle) perimeterptr() float64 {
	return float64(2*r.length + 2*r.breadth)
}
func structMethodCheck() {
	didi := rectangle{length: 20, breadth: 30}
	fmt.Println("Rectangle details : ", didi)

	fmt.Println("Rectangle area : ", didi.area())
	didir := &didi
	fmt.Println("Rectangle perimeter : ", didir.perimeterptr())
}

//////////////////////////////////////
func (r rectangle) diagonal() float64 {
	diag := math.Sqrt(float64((r.length * r.length) + (r.breadth * r.breadth)))
	return diag
}

type geometry interface {
	area() float64
	perimeter() float64
}

func measurements(g geometry) {
	fmt.Println("Geometry : ", g)
	fmt.Println("Area : ", g.area())
	fmt.Println("Perimeter : ", g.perimeter())
}
func interfaceCheck() {
	red := rectangle{breadth: 33, length: 24}
	measurements(red)
	fmt.Println("Diagonal : ", red.diagonal())
}

//////////////////////////////////////

func errorCheck(b int) (float64, error) {
	if b == 0 {
		return -1, errors.New("Anything other than 0")
	}
	return float64(b / 10), nil
}

func main() {
	fmt.Println(errorCheck(0))
}
