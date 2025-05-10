// package main
// import "fmt"
// const NMAX int = 100
// func main(){
// 	var A[NMAX]int
// 	var n, i int
// 	fmt.Scan(&n)
// 	for i = 0; i < n; i++{
// 		fmt.Scan(&A[i])
// 	}
// 	hasil := avaregValue(A[:n], n)
// 	fmt.Print(hasil)
// }
// func avaregValue(nums [] int, n int)int{
// 	var i, jum, c int
// 	for i = 0; i < n; i++{
// 		if nums[i]%2 == 0 && nums[i]%3 == 0{
// 			jum+= nums[i]
// 			c++
// 		}
// 	}
// 	if c == 0{
// 		return 0
// 	}else{
// 		return jum/c
// 	}
// }



// package main
// import "fmt"

// func subtractProductAndSum(n int) int {
//     x := n
//     kali := 1
//     tambah := 0
//     for n > 0 {
//         hasil := n % 10
//         kali = kali * hasil
//         n = n / 10
//     }
//     for x > 0 {
//         hasil := x % 10
//         tambah = tambah + hasil
//         x = x / 10
//     }

//     return kali - tambah
// }
// func main() {
// 	var n int
// 	fmt.Scan(&n)
//     fmt.Println(subtractProductAndSum(n))
// }



// package main
// import "fmt"
// const NMAX int = 100
// func main(){
// 	var A[NMAX]int
// 	var n, i int
// 	fmt.Scan(&n)
// 	for i = 0; i <n; i++{
// 		fmt.Scan(&A[i])
// 	}
// 	hasil := findNumber(n, A[:n])
// 	fmt.Print(hasil)
// }
// func findNumber(n int, A[] int)int{
// 	var i, genap int
// 	for i = 0; i<n; i++{
// 		if A[i]%2 == 0{
// 			genap++
// 		}
// 	}
// 	return genap
// }




// package main
// import "fmt"
// const NMAX int = 100
// func main(){
// 	var A[NMAX]int
// 	var B[NMAX]int
// 	var C[NMAX]int
// 	var n, n1, i, j, k int
// 	fmt.Scan(&n)
// 	for i = 0; i < n; i++{
// 		fmt.Scan(&A[i])
// 	}
// 	fmt.Scan(&n1)
// 	for i = 0; i < n; i++{
// 		fmt.Scan(&B[i])
// 	}
// 	k = 0
// 	for i = 0; i < n; i++{
// 		for j = 0; j < n1; j++{
// 			if A[i] == B[j]{
// 				C[k] = A[i]
// 				k++
// 				break
// 			}
// 		}
// 	}
// 	for i = 0; i < k; i++{
// 		fmt.Print(C[i], " ")
// 	}
// }

