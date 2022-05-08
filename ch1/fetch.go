// URLの内容を表示する
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) // エラーがなければレスポンス構造体respで結果を返す。respのBodyフィールドは読み込み可能なストリームとしてサーバからのレスポンスを含む
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // 標準エラー出力にエラー内容を出力する
			os.Exit(1)                                 // プロセスをステータスコード1で終了させる
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
	}
}
