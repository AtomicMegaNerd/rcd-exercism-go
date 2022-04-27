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
func (std *SearchTreeData) MapString(f func(int) string) (result []string) {
	result = make([]string, 0)
	var helper func(*SearchTreeData, []string, func(int) string) []string
	helper = func(std *SearchTreeData, result []string, f func(int) string) []string {
		if std.left != nil {
			result = helper(std.left, result, f)
		}
		result = append(result, f(std.data))
		if std.right != nil {
			result = helper(std.right, result, f)
		}
		return result
	}
	return helper(std, result, f)
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(f func(int) int) (result []int) {
	result = make([]int, 0)
	var helper func(*SearchTreeData, []int, func(int) int) []int
	helper = func(std *SearchTreeData, result []int, f func(int) int) []int {
		if std.left != nil {
			result = helper(std.left, result, f)
		}
		result = append(result, f(std.data))
		if std.right != nil {
			result = helper(std.right, result, f)
		}
		return result
	}
	return helper(std, result, f)
}
