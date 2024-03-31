package main
import (
	"fmt"
)
func main(){
	var n int
	fmt.Scanln(&n)
	arr := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	left,right := 0,len(arr)-1
	for left<=right {
		arr[left],arr[right] = arr[right], arr[left]
		left++
		right--
	}
	fmt.Println("Reversed array:", arr)
}