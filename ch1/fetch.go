// URLの内容を表示する
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url) // エラーがなければレスポンス構造体respで結果を返す。respのBodyフィールドは読み込み可能なストリームとしてサーバからのレスポンスを含む
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // 標準エラー出力にエラー内容を出力する
			os.Exit(1)                                 // プロセスをステータスコード1で終了させる
		}
		fmt.Printf("status code: %v\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
	}
}
