package byte

import "fmt"

type User struct {
	Id 	 int
	Name string
}

func ByteSurvey() {
	str := "123abc"
	byte := []byte(str)

	fmt.Println("byteは", byte)
	// 出力→ byteは [49 50 51 97 98 99]
	// 49⇆1

	str = string(byte)
	fmt.Println("strは", str)

}