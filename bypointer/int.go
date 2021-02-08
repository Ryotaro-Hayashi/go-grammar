package bypointer

import "fmt"

// 値渡し
func IntByValue(value int) {
	fmt.Println("-----値渡しを実行-----")
	value += 10
	fmt.Println("初期値10に10足して→", value) // 値
	fmt.Println("関数の引数に渡されるとコピーされるので元の引数のポインタとは異なるポインタ→", &value) // ポインタ
}

// ポインタ渡し
func IntByPointer(valuePointer *int) {
	fmt.Println("-----ポインタ渡しを実行-----")
	*valuePointer += 20
	fmt.Println("初期値10に20足して→",*valuePointer) // 値
	fmt.Println("ポインタで渡したので元の引数のポインタと同じポインタ→", valuePointer) // ポインタ
}

func IntByValuePointer() {

	fmt.Print("-----Intでのポインタの練習-----\n\n")

	var integer = 10
	fmt.Println("1.Intの初期値は", integer)

	// 「&変数」で変数のポインタを取得
	var integerPointer *int = &integer // 明示的に「*int」を付ける必要はない
	fmt.Println("1.「&変数」で取得した初期値のポインタは", integerPointer)

	// 「*ポインタ型変数」で中身へアクセス
	fmt.Println("「*ポインタ型変数」でポインタの中身へアクセスすると", *integerPointer) // ポインタ変数から中身へのアクセス

	v := 10
	fmt.Println("2.Intの初期値は", v)
	fmt.Println("2.「&変数」で取得した初期値のポインタは", &v)

	IntByValue(v) // 値渡し

	// 初期値
	fmt.Println("関数内のものはコピーなので渡す前の変数の値は変わらない→", v)
	fmt.Println("コピーと渡す前の変数のポインタは異なる→", &v)

	IntByPointer(&v) // ポインタ渡し

	// 初期値
	fmt.Println("ポインタ渡しで関数を実行したので値はもちろん変わる→", v)
	fmt.Println("ポインタは渡す前の変数のポインタ→", &v)
}