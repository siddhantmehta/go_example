package main
import("fmt")
func main(){
	x:= 15
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is smaller than or equal to 5")
	}
	a,b := 10.0,2.0
	if frac := a/b; frac > 0.5{
		fmt.Println("a is more than half of b")
	}
}