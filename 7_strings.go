package main
import("fmt")
func main(){
	book := "The colour of magic"
	fmt.Println(book)
	
	fmt.Println(len(book))
	fmt.Printf("book[0] = %v, (type of %T)\n", book[0],book[0])
	fmt.Println(book[4:11]) //string slicing
	fmt.Println(book[4:])
	fmt.Println(book[:4])
}