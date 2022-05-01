package main

import "log"

//明示的に型を指定する必要がなければ、短縮形にする。

func main() {
	// two-types
	// ゼロ値
	var empty string

	// 値を設定かつ右辺から型がわかる場合
	one := 1
	// two-types

	// no-use-var
	// 短縮記法の方が良い
	var two = 2
	// no-use-var

	// both-ok
	// varを使っても使わなくてもいい
	var three int8 = 3
	four := int16(4)
	// both-ok

	log.Println(empty, one, two, three, four)
}
