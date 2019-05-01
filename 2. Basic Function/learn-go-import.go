package main

// format and math
// these are the standard packages
// that go provided
import "fmt"
import "math"

// package are reuseable except main package
// if we had some defined package then just
// import it to main or another package by
// eg.
// import "github.com/adithyavisnu/learn-go/words"

func main() {
	fmt.Println("Halo Adith!")
	fmt.Println(math.Abs(1))
}
