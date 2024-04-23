package main
import (
	"fmt"
)
func main(){
	var n int
	fmt.Scanln(&n)
	arr:= make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	var k int
	fmt.Scanln(&k)
	

}