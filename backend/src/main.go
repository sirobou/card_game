package main

import (
	"fmt"
	"net/http"
)

func main() {
	// HTTPハンドラ関数を定義
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Casino!")
	})

	// サーバーを指定のポートで起動
	port := ":8080" // ポート番号を必要に応じて変更
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
