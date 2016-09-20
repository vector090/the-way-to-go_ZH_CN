// uc_test.go
package uc

import "testing"

type ucTest struct {
	in, out string
}

var ucTests = []ucTest{
	ucTest{"abc", "ABC"},
	ucTest{"cvo-az", "CVO-AZ"},
	ucTest{"Antwerp", "ANTWERP"},
}

func TestUC(t *testing.T) {
	for _, ut := range ucTests {
		uc := UpperCase(ut.in)
		t.Log("uc ", uc) // will only be printed when "-v" is specified for "go test".
		if uc != ut.out {
			t.Errorf("UpperCase(%s) = %s, must be %s.", ut.in, uc, ut.out)
		}
	}
}

func TestUC2(t *testing.T) {
	//dummy
}

func BenchmarkReverse(b *testing.B) {
	//dummy
}
