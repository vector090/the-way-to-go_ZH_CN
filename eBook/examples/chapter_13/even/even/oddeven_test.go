// oddeven_test.go
package even

import "testing"
import "fmt"

func TestEven(t *testing.T) {

	// Comparnig ways to print during testing.
	fmt.Println("This still get printed")
	t.Log("Another way to print when testing")
	println("This prints too, but quite untraceable")

	if !Even(10) {
		t.Log(" 10 must be even!")
		t.Fail()
	}
	if Even(7) {
		t.Log(" 7 is not even!")
		t.Fail()
	}

}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log(" 11 must be odd!")
		t.Fail()
	}
	if Odd(10) {
		t.Log(" 10 is not odd!")
		t.Fail()
	}
}
