package main

import (
	"fmt"
)

func main() {

	// s := "haha hoho"
	// fmt.Println(strings.HasPrefix(s, "ha"))
	// fmt.Println(strings.Index(s, "a"))
	// fmt.Println(strings.ToUpper(s))
	// fmt.Println(strings.Contains(s, "ha"))
	// fmt.Println(strings.HasSuffix(s, "ha"))
	// fmt.Println(strings.Replace(s, "ha", "hp", 2))

	// fields := strings.Fields(s)
	// fmt.Println(fields[0])

	// s := []string{"hello", "word"}
	// fmt.Println(strings.Join(s, "-"))

	// s := "王杰哈哈哈"
	// fmt.Println(s[10])
	// fmt.Println(utf8.RuneCountInString(s))


	age := Age(25)
	age.String()



}

func sum(a, b int) func() int {
	i := 1
	return func() int {
		i++;
		return i
	}
}

// 这个类型是一个 int 的值，Age(里面是 int 值)
type Age uint

func (age Age) String() {
	fmt.Println("the age is", age)
}














