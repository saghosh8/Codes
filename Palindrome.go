package main
import "fmt"

func main() {
	var num, remainder, temp int
	var rev_num int =0

	fmt.Print("Enter the number:")
	fmt.Scan(&num)

	temp=num


	for{
		remainder = num%10
		rev_num = rev_num*10 + remainder
		num /= 10
		
		if(num==0) {
			break 
		}
	}
		
	if(temp==rev_num){
		fmt.Println("The number is Palindrome")
		}else{
		fmt.Printf("The number is not Palindrome")
	}

}
