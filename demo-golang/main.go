package main
import "fmt"
func main(){
	fmt.Println("Hello world")
	//Learning about variable and function
	//Control statement: if:, for
	var a,b,c,e int
	a = 100
	b = 1000
	c = a + b
	fmt.Println(c)
	d:= c*2
	fmt.Println(d)
	e = sum(100,102)
	fmt.Println(e)
	//Slice
	var myslice []int
	myslice= []int{24, 2, 4, 15, 90}
	fmt.Println(len(myslice)) 
	result := contain( 22, myslice)
	fmt.Println(result)
	chaythu := contain(24, myslice)
	fmt.Println(chaythu)
	printSlice([]int{2, 34, 5, 12, 6})
}
func sum(a int, b int) int{
	var result int
	result = a+b
	return result
}
/*func contain(x int, y []int) bool{
	var myslice []int
	myslice= 
	for k := range myslice {
		if x == 3{
			return true
		}
	}

	return false
}*/
func contain(x int, y []int) bool {
	for _, v := range y {
		if v == x {
			return true
		}
	}
	return false
}
//Bai tap ve nha
func printSlice(h []int) {
	for _, gtri := range h {
		fmt.Println("Giá trị:", gtri)
	}
}