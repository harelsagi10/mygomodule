package arithmetic
import "fmt"

func Factorial(n int) int {
    var f int = 1
    for i := 2; i <= n; i++ {
f *= i }
return f }


func SayHello(){
fmt.Println("hello Go")
}
