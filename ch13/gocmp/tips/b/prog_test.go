package a

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// test-start
func TestX(t *testing.T) {
	type X struct {
		numUnExport int
		NumExport   int
	}

	num1 := X{100, -1}
	num2 := X{999, -1}

	opt := cmp.AllowUnexported(X{})

	if diff := cmp.Diff(num1, num2, opt); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}

// test-end
