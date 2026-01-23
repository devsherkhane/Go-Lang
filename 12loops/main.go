package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	//fmt.Println("Days :", days)

		// for i:=0 ; i<len(days) ; i++ {
		// 	fmt.Println("Day :",days[i])
		// }

	   	// for i := range(days){
		// 	fmt.Println("Day :",days[i])
		// }


		for i , val := range(days){
			fmt.Printf("The index is %v , & the day is %v\n ",i ,val)
		}
}
