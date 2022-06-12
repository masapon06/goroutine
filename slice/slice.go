package slice

import "fmt"

// https://go-tour-jp.appspot.com/moretypes/7
// 基本的なスライス
func fundamental() {
	fmt.Println("- Slices -")
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println()
}

// https://go-tour-jp.appspot.com/moretypes/8
// 配列とスライスの関係
func likeReferenceArray() {
	fmt.Println("- Slices are like references to arrays -")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// スライスの値を変更
	b[0] = "XXX"
	fmt.Println(a, b)
	// namesの値も変更される. => sliceは配列への参照のようなもの
	fmt.Println(names)
	fmt.Println()
}

// https://go-tour-jp.appspot.com/moretypes/9
// スライスのリテラルは、配列のリテラルを参照する
func sliceLiterals() {
	fmt.Println("- Slice literals -")
	// [6]int{2, 3, 5, 7, 11, 13} という配列を参照
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// [6]bool{true, false, true, true, false, true} という配列を参照
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	fmt.Println()
}

// Slice length(長さ) and capacity（容量）
// // https://go-tour-jp.appspot.com/moretypes/11
func lengthAndCapacity() {
	fmt.Println("- Slice length and capacity -")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// 以下、「簡易スライス式」
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []
	// => 「容量」 =「元となる配列の要素数」

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	// 下限を指定すると、「容量」 - 「下限」が新たな容量になる
	s = s[2:]     // cap = 6 - 2 = 4
	printSlice(s) // len=2 cap=4 [5 7]
	s = s[1:]     // cap = 4 - 1 = 3
	printSlice(s) // len=1 cap=3 [7]
	fmt.Println()
}

func nilSlices() {
	fmt.Println("- Nil slices -")
	var s []int
	printSlice(s)
	if s == nil {
		fmt.Println("nil!")
	}
	fmt.Println()
}

func createSliceWithMake() {
	fmt.Println("- Creating a slice with make -")
	a := make([]int, 5)
	fmt.Print("a: ")
	printSlice(a)

	b := make([]int, 0, 5) // 3番目の引数 => capacity
	fmt.Print("b: ")
	printSlice(b)

	c := b[:2]
	fmt.Print("c: ")
	printSlice(c)

	d := c[2:5]
	fmt.Print("d: ")
	printSlice(d)
	fmt.Println()
}

func appendToSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// クロージャ
func closures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// fmt.Println("- XXXX -")
// print
func Output() {
	fmt.Println("=== Slices ===")
	fundamental()
	likeReferenceArray()
	sliceLiterals()
	lengthAndCapacity()
	nilSlices()
	createSliceWithMake()
	closures()
}
