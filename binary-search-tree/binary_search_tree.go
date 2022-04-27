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
	return []string{}
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(f func(int) int) (result []int) {
	if std.left == nil {
		result = append(result, std.data)
	} else {
		result = std.left.MapInt(f)
	}
	if std.right == nil {
		result = append(result, std.data)
	} else {
		result = std.right.MapInt(f)
	}
	return
}
