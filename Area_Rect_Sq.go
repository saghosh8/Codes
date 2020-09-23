package main
 
import "fmt"
 
 
func main(){

var l,b,a,s int
fmt.Print("Enter the Length:")
fmt.Scan(&l)
fmt.Print("Enter the Breadth:")
fmt.Scan(&b)
a = l*b
fmt.Print("Area of Rectangle:", a)

fmt.Print("Side Of the Square:")
fmt.Scan(&s)
a = s*s
fmt.Print("Area Of Square:", a)

}