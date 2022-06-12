// 配列操作
package array

import "fmt"

type StringArr = [2]string

func handleStringArray(arr StringArr) StringArr {
	arr[1] = "masapon06"

	return arr
}

func Output() {
	fmt.Println("=== array ===")
	var a StringArr
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := handleStringArray(a)
	fmt.Println(b)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println()
}
