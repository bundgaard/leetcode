package roman_number

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		Input    *ListNode
		Expected bool
	}{
		{
			// 1 1 2 1
			Input:    &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
			Expected: false,
		},
		{
			Input:    &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: nil}}},
			Expected: false,
		},
		{
			Input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			Expected: false,
		},
		{
			Input:    &ListNode{Val: 0, Next: &ListNode{Val: 0}},
			Expected: true,
		},
		{
			Input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			Expected: true,
		},
		{
			Input:    &ListNode{Val: 1},
			Expected: true,
		},
	}

	for _, test := range tests {
		if got := isPalindrome(test.Input); got != test.Expected {
			t.Errorf("expected %t. got %t; %v\n", test.Expected, got, test.Input)
		}
	}
}
