package main
import("fmt")
func main(){
	a:=1
	for a < 21 {
		if a % 3 == 0 && a % 5 == 0{
			fmt.Println("fizzbuzz")
		} else if a % 5 == 0{
			fmt.Println("buzz")
		} else if a % 3 == 0{
			fmt.Println("fizz")
		} else{
			fmt.Println(a)
		}
		a++
	}
}