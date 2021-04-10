package main
import("fmt")
func main(){
	n :=42
	s := fmt.Sprintf("%d",n)
	fmt.Printf("s = %q (type %T)\n",s,s)
	fmt.Printf("s = %s (type %T)\n",s,s)
	count := 0
	for a := 1000; a<=9999; a++{
		for b := a; b<=9999; b++{
			n := a * b
			s = fmt.Sprintf("%d", n)
			if s[0] == s[len(s) - 1]{
				count++
			}
		}
	} 
	fmt.Printf("Total count = %v",count)
}