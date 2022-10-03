package doc_test

import "fmt"

// PrintDoc() 함수에 대한 예제입니다.
func ExamplePrintDoc() {
	fmt.Println("This is package level example")
}

// TextDoc 의 PrintDoc() 메서드에 대한 예제입니다.
func ExampleTextDoc_PrintDoc() {
	fmt.Println("This is PRintDoc() example")
}

// TextDoc 에 대한 예제입니다.
func ExampleTextDoc_lines() {
	fmt.Println("This is lines() example")
}
