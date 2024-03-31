package main
import (
	"fmt"
	"math"
)
func main(){
	var n int
	fmt.Scanln(&n)
	arr := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	maxInt := math.MaxInt
	minInt := math.MinInt
	for i:=0;i<n;i++{
		if arr[i]>minInt{
			minInt = arr[i]
		}
		if arr[i]<maxInt{
			maxInt = arr[i]
		}

	}

	fmt.Println("Maximum Integer:",minInt)
	fmt.Println("Minimum Integer:",maxInt)
}