package easy

import "testing"

func TestHasCycle(t *testing.T) {
	circle := &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}
	circle.Next.Next.Next = circle
	if v := hasCycle(&ListNode{Val: 3, Next: circle}); !v {
		t.Fatal(v)
	}
	circle = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: circle}}
	if v := hasCycle(circle); !v {
		t.Fatal(v)
	}
	if v := hasCycle(&ListNode{Val: 1}); v {
		t.Fatal(v)
	}
}
