package solution1

import (
	"testing"

	. "github.com/KevinBaiSg/myleecode/common"
)

func TestIsScramble(t *testing.T)  {
	bstToGst(MakeTree("[3,9,20,null,null,15,7]"))
}

func BenchmarkIsScramble(b *testing.B) {
	bstToGst()
}