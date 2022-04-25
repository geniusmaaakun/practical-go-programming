package main

import (
	"encoding/csv"
	"log"
	"os"
)

// write-sample-start
func main() {
	records := [][]string{
		{"書籍名", "出版年", "ページ数"},
		{"Go言語によるWebアプリケーション開発", "2016", "280"},
		{"Go言語による並行処理", "2018", "256"},
		{"Go言語でつくるインタプリタ", "2018", "316"},
	}

	f, err := os.OpenFile("oreilly.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush() // バッファにあるデータを書き込み

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// write-sample-end
