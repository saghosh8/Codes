package main
 
import "fmt"
 
func main(){

var x int
fmt.Print("Multiplication Table of:" )
fmt.Scan(&x)
	y:=1
for {
		if (y>10){
		break;
		}
	
fmt.Println(x, "X", y, "=", x*y)
y++
	}
}