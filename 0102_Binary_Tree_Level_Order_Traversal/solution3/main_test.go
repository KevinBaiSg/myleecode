package solution3

import "testing"

func TestIsScramble(t *testing.T)  {
	levelOrder(&TreeNode{
		3,
		&TreeNode {
			9, nil, nil,
		},
		&TreeNode {
			20,
			&TreeNode {
				15, nil, nil,
			},
			&TreeNode {
				7, nil, nil,
			},
		},
	})
}