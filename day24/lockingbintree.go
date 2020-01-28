package day24

// LockingBinaryTree is a lockable binary tree implementation
// A binary tree node can be locked or unlocked only if all of its descendants or ancestors are not locked
type LockingBinaryTree struct {
	parent         *LockingBinaryTree
	Left           *LockingBinaryTree
	Right          *LockingBinaryTree
	isLocked       bool
	isOnLockedPath bool
}

// NewLockingBinaryTree builds binary tree node instance with the pointer to parent node
func NewLockingBinaryTree(left *LockingBinaryTree, right *LockingBinaryTree) *LockingBinaryTree {
	root := &LockingBinaryTree{nil, nil, nil, false, false}

	root.Left = left
	root.Right = right

	if left != nil {
		left.parent = root
	}
	if right != nil {
		right.parent = root
	}

	return root
}

// IsLocked returns whether the node is locked
func (t *LockingBinaryTree) IsLocked() bool {
	return t.isLocked
}

// Lock attempts to lock the node. If it cannot be locked, then it returns false. Otherwise, it locks it and returns true
func (t *LockingBinaryTree) Lock() bool {
	if t.isLocked || t.isLockedPathBottomUp() {
		return false
	}

	target := t.parent
	for {
		if target == nil {
			break
		}
		target.isOnLockedPath = true
		target = target.parent
	}
	t.isLocked = true

	return true
}

// Unlock unlocks the node. If it cannot be unlocked, then it returns false. Otherwise, it unlocks it and returns true
func (t *LockingBinaryTree) Unlock() bool {
	if !t.isLocked {
		return false
	}

	target := t.parent
	for {
		if target == nil {
			break
		}
		target.isOnLockedPath = false
		target = target.parent
	}
	t.isLocked = false

	return true
}

func (t *LockingBinaryTree) isLockedPathBottomUp() bool {
	if t.isOnLockedPath {
		return true
	}

	target := t.parent
	for {
		if target == nil {
			break
		}
		if target.isOnLockedPath || target.isLocked {
			return true
		}
		target = target.parent
	}

	return false
}
