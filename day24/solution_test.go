package day24

import (
	"testing"
)

func TestSolution(t *testing.T) {
	root := NewLockingBinaryTree(
		NewLockingBinaryTree(
			NewLockingBinaryTree(nil, nil), nil),
		NewLockingBinaryTree(
			nil, NewLockingBinaryTree(nil, nil)))

	var wasLocked bool
	if wasLocked = root.Left.Lock(); !wasLocked {
		t.Errorf("In newly created tree it should be possible to lock a node")
	}
	if wasLocked = root.Lock(); wasLocked {
		t.Errorf("Root of the tree with locked node should not be locked")
	}
	if wasLocked = root.Left.Left.Lock(); wasLocked {
		t.Errorf("Other nodes of the tree with locked node should not be locked")
	}

	if wasUnlocked := root.Left.Unlock(); !wasUnlocked {
		t.Errorf("Locked node of the tree should be unlocked")
	}
	if wasLocked = root.Lock(); !wasLocked {
		t.Errorf("Root of the tree with no locked node should be locked")
	}
	if wasLocked = root.Left.Left.Lock(); wasLocked {
		t.Errorf("Other nodes of the tree with locked node should not be locked")
	}
}
