package main

import (
	"bufio"
	"fmt"
	"os"
)
import "practice/sub" // subは'package'で指定したパッケージ名。.modファイルで決めたmodule名の絶対パスで指定
import (
	"practice/array"
	"practice/slice"
)

func main() {
	fmt.Print("いずれかのコマンドを入力してください。 (hello/array/slice)>")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()
		fmt.Println("in: ", in)
		switch in {
		case "hello":
			fmt.Printf("hello, world\n")
			sub.Hello()
			goto L
		case "array":
			array.Output()
			goto L
		case "slice":
			slice.Output()
			goto L
		case "EXIT":
			fmt.Println("exitしました。")
		default:
			fmt.Println("コマンドが不正です。入力を終えたいときは、EXITと入力してください。")
			continue
		}
	}
L:
}
