package regex

import (
	"testing"
)

func Test(t *testing.T) {
	Comp(`a (test)( [0-9]|)`).RepFunc([]byte("this is a test and a test 2"), func(b func(int) []byte) []byte {
		return []byte{}
	})
}
