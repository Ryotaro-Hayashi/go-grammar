package main

import (
	"go-grammar/bypointer"
	"go-grammar/byte"
)

func main() {
	// ポインターについて
	bypointer.IntByValuePointer()
	bypointer.StructByValuePointer()

	// バイトスライスについて
	byte.ByteSurvey()
}