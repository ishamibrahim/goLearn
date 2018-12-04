package main

import "fmt"

const s string = "I'm Isham"

func typeChange() {
	const s string = "Im Isham inside change"
	fmt.Println(s)
}

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

func arrayCheck() {
	var a [1]int
	fmt.Println("emp:", a)
}

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

func addThree(a, b, c int) int {
	return a + b + c
}
func funcCheck() {
	var a, b, c int = 3, 5, 7
	fmt.Println(addThree(a, b, c))
}

func printPars(nums ...int) {
	for _, r := range nums {
		fmt.Println(r)
	}
}
func variardicCheck() {
	pars := []int{3, 5, 7, 4, 9}
	printPars(pars...)
}

func main() {
	variardicCheck()
}
