package solution1

import (
	"testing"

	"github.com/KevinBaiSg/myleecode/common/makeTree"
)

func TestIsScramble(t *testing.T)  {
	bstToGst(makeTree.MakeTree("[3,9,20,null,null,15,7]"))
}

func BenchmarkIsScramble(b *testing.B) {
	bstToGst()
}