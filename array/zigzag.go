package main
import "fmt"

func fillzigzagmatrix(n int)[][]int{
	matrix:= make([][]int,n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1
	for layer:=0; layer<(n+1)/2; layer++{
		for i:=layer; i<n-layer;i++{
			matrix[layer][i] = num
			num++
		}
		for i:=layer+1; i<n-layer;i++{
			matrix[i][n-layer-1] = num
			num++
		}
		for i:=n-layer-2;i>=layer;i--{
			matrix[n-layer-1][i] = num
			num++
		}
		for i:=n-layer-2;i>layer;i--{
			matrix[i][layer] = num
			num++
		}
	}
	return matrix
}



func main(){
	var n int
	fmt.Scanln(&n)
	matrix := fillzigzagmatrix(n)
	for _,row := range matrix{
		fmt.Println(row)
	}
	
}