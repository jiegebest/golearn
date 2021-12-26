package main

func main() {
	// var i int = 10
	// var f float64 = 12.3
	// var b bool = false
	// var s string = "aaa"
	// fmt.Println(i, f, b, s)
	// 10 12.3 false aaa


	// var i = 10
	// var f = 12.3
	// var b = false
	// var s = "bbb"
	// fmt.Println(i, f, b, s)
	// 10 12.3 false bbb

	// var (
	// 	i int = 10
	// 	f float64 = 12.3
	// 	b bool = false
	// 	s string = "bbb"
	// )
	// fmt.Println(i, f, b, s)

	// var (
	// 	i = 10
	// 	f = 12.3
	// 	b = false
	// 	s = "bbbb"
	// )
	// fmt.Println(i, f, b, s)

	// var f64 float64 = 10.33
	// var bf bool = true
	// var s1 string = "hello"
	// var s2 string = "world"
	// fmt.Println(f64, bf)
	// fmt.Println(s1 + s2)
	// s1 += s2
	// fmt.Println(s1)

	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Println(i, f, b, s)

	// i := 10
	// b := true
	// s := "hello"
	// fmt.Println(i, b, s)

	// i := 10
	// pi := &i
	// fmt.Println(pi)	// 0xc000018098 指针取地址
	// fmt.Println(*pi) // 10 值取值

	// i := 10
	// i = 20
	// fmt.Println(i)

	// const name string = "飞雪无情"
	// const name1 = "飞雪无情"
	// fmt.Println(name, name1)

	// fmt.Println(one, two, three, four)

	// i := 10
	// i2s := strconv.Itoa(i)
	// s2i, err := strconv.Atoi(i2s)
	// fmt.Println(i2s, s2i, err)

	// f := 1.2345678
	// // fmt 代表格式
	// // prec 代表精确度
	// // bitSize 代表f的来源
	// s := strconv.FormatFloat(f, 'f', 2, 64)
	// fmt.Println(s)

	// s := "12.3"
	// s2f, _ := strconv.ParseFloat(s, 64)
	// fmt.Println(s2f)

	// b := true
	// b2s := strconv.FormatBool(b)
	// s2b, err := strconv.ParseBool(b2s)
	// fmt.Println(s2b, err)

	// i := 10
	// i2f := float64(i)
	// f := 1.23
	// f2i := int(f)
	// fmt.Println(i2f, f2i)

	// s := "hello world"
	// fmt.Println(strings.HasPrefix(s, "he"))
	// fmt.Println(strings.HasSuffix(s, "ld"))
	// fmt.Println(strings.Index(s, "hh"))
	// fmt.Println(strings.ToUpper(s))
	// fmt.Println(strings.Contains(s, "ha"))
	// fmt.Println(strings.Replace(s, "he", "", 1))
	// fields := strings.Fields(s)
	// fmt.Println(cap(fields))
	// s1 := []string{"hello", "world"}
	// fmt.Println(strings.Join(s1, "-"))

	// c := complex(1, 3)
	// c1 := complex(2, 5)
	// fmt.Println(real(c), imag(c))
	// fmt.Println(c + c1) // 实部与实部相加，虚部与虚部相加
	// fmt.Println(c - c1) // 实部与实部相减，虚部与虚部相减
	// fmt.Println(c * c1) // 交错相乘作为虚部
	// fmt.Println(c / c1)

}

// const (
// 	one = 1
// 	two = 2
// 	three = 4
// 	four = 8
// )

const (
	one = 10 + iota
	two
	three
	four
)























