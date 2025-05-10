package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var B[NMAX]int
	var C[NMAX]int
	var n, n1, i, j, k int
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	fmt.Scan(&n1)
	for i = 0; i < n; i++{
		fmt.Scan(&B[i])
	}
	k = 0
	for i = 0; i < n; i++{
		for j = 0; j < n1; j++{
			if A[i] == B[j]{
				C[k] = A[i]
				k++
				break
			}
		}
	}
	for i = 0; i < k; i++{
		fmt.Print(C[i], " ")
	}
}

