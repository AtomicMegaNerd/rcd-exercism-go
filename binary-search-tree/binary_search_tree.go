package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func NewBst(i int) SearchTreeData {
	return SearchTreeData{nil, i, nil}
}

func (std *SearchTreeData) Insert(i int) {
	if i <= std.data {
		if std.left == nil {
			std.left = &SearchTreeData{nil, i, nil}
		} else {
			std.left.Insert(i)
		}
	} else {
		if std.right == nil {
			std.right = &SearchTreeData{nil, i, nil}
		} else {
			std.right.Insert(i)
		}
	}
}

// MapString returns the ordered contents of SearchTreeData as a []string.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []string ["1", "3", "5", "7"].
func (std *SearchTreeData) MapString(func(int) string) (result []string) {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	return []string{}
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(func(int) int) (result []int) {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	return []int{}
}