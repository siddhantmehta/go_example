package main
import("fmt")
func main(){
	a := []string{"one","two","three"}
	fmt.Printf("a = %v (type %T)\n",a,a)
	//length
	fmt.Println(len(a))
	// 0 indexing
	fmt.Println(a[1])
	//slices
	fmt.Println(a[1:])
	//Single range
	for i := range a{
		fmt.Println(i)
	}
	//Double range
	for i,name := range a{
		fmt.Printf("name = %s at %v\n",name,i)
	}
	//Double range with ignore index
	for _,name := range a{
		fmt.Println(name)
	}
		//append
	b := append(a,"four")
	fmt.Println(b)

}