package main

import "fmt"

func main() {

	//// 变量声明
	//var power1 int = 1000
	//power2 := 1000
	//power3, power4 := 2000, 3000
	//var power5, power6 int  = 4000, 5000
	//
	//fmt.Println(power1)
	//fmt.Println(power2)
	//fmt.Println(power3)
	//fmt.Println(power4)
	//fmt.Println(power5)
	//fmt.Println(power6)
	//
	//// 函数使用
	//c, _ := add3(100, 200)
	//fmt.Println(c)
	//
	//person1 := new(Person)
	//// or
	//person2 := &Person{}
	//
	//girl := &Girl {
	//	Person: &Person {
	//		Sex: 1,
	//	},
	//}
	//
	//girl.showSex()
	//// or
	//girl.Person.showSex()


	scores := make([]int, 9, 10)
	////scores = append(scores, 10)
	scores = append(scores, 10)
	scores = append(scores, 10)
	scores = append(scores, 10)
	//scores = append(scores, 10)
	//scores = append(scores, 10)
	//scores = append(scores, 10)
	//scores = append(scores, 10)
	//scores = append(scores, 10)
	//scores = append(scores, 10)
	//scores = append(scores, 10)
	//scores = append(scores, 10)
	//scores = append(scores, 10)
	//scores = scores[0:8]
	scores[7] = 9033
	//scores[0] = 9033
	//fmt.Println(scores)

	//scores := make([]int, 5)
	//scores = append(scores, 9332)
	fmt.Println(len(scores))

	fmt.Println(scores)
}

func log(message string)  {
	
}

func add(a, b int) int {
	return a + b
}

func add1(a, b int) (c int)  {
	c = a + b
	return
}

func add2(a, b int) (c int, compare bool)  {
	c = a + b
	compare = a > b

	return
}

func add3(a, b int) (int, bool)  {
	c := a + b
	compare := a > b

	return c, compare
}

const (
	Female = iota
	Male
)
type Person struct {
	Name string
	Age int
	Sex int
}

func (p *Person) showSex() {
	fmt.Println(p.Sex)
}

type Girl struct {
	*Person
}
