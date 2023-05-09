package main

import(
	"fmt"
)

func Max(vals ...int) int{
	temp := -1111111
	for _, val := range vals{
		if val > temp{
			temp = val
		}
	}
	return temp
}

func main(){
	fmt.Println(Max(1,2,3,4,5,6,7))
	fmt.Println(Max(1,2,3,4,5,6,75555))
	fmt.Println(Max(1))
	fmt.Println(Max())
}