package bypointer

import "fmt"

type Person struct {
	Name string
	Age  int
}

func StructByValue(person Person) {
	fmt.Println("-----値渡しを実行-----")
	person.Name = "ハヤシ"

	// やはり引数に渡されるとコピーされる
	fmt.Println("引数に渡された後名前を変更された構造体", person)
	fmt.Printf("引数に渡され後名前を変更された構造体のポインタは%p\n", &person)
}

func StructByPointer(personPointer *Person) {
	fmt.Println("-----ポインタ渡しを実行-----")
	personPointer.Name = "遼太朗"

	fmt.Println("引数に渡された後名前を変更された構造体", *personPointer)
	fmt.Printf("引数に渡され後名前を変更された構造体のポインタは%p\n", personPointer)
}

func StructByValuePointer() {

	fmt.Print("-----Structでのポインタの練習-----\n\n")

	person := Person{
		Name: "りょうたろう",
		Age:  22,
	}

	fmt.Println("Structの初期値は", person)

	// 「&変数」で変数のポインタを取得
	fmt.Println("構造体は「&変数」のみではポインタを出力してくれない→", &person)
	// 構造体のポインタは「&変数」だけでなく出力の際もフォーマット指定しないと出力されない
	fmt.Printf("構造体は「&変数」に加えてフォーマット指定して出力した初期値のポインタは%p\n", &person)

	StructByValue(person) // 値渡し

	// 初期値
	fmt.Println("関数内のものはコピーなので渡す前の構造体は変わらない→", person)
	fmt.Printf("コピーと渡す前の変数のポインタは異なる→%p\n", &person)

	StructByPointer(&person)

	// 初期値
	fmt.Println("ポインタ渡しで関数を実行したので値はもちろん変わる→", person)
	fmt.Printf("ポインタは渡す前の変数のポインタ→%p\n", &person)

}
