package main
 
import "fmt"

type Student struct {
	name string
	math, english float64
}
 
func main() {
	// fmt.Printf("Hello World!")
	
	//https://www.youtube.com/watch?v=kPXfMFJ0oIE
	//1:10:00あたりが構造体の説明
	//構造体Studentをsに代入　構造体の初期化

	var s Student
	s.name ="sato"
	s.math = 80
	s.english = 70
	fmt.Println(s)
}