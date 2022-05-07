package main

import (
	"fmt"
	"os" // NOTE: プラットフォームから独立した形式でオペレーティングシステムを扱うための関数と値を提供する
)

func main() {
	var s, sep string // NOTE: 暗黙的にゼロ値として初期化される
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
