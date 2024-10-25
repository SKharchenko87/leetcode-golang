package stucture

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const NULL = math.MinInt
const null = math.MinInt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Функция для представления дерева в матрице
//func (tn *TreeNode) ToMatrix() [][]int {
//	maxLevel := maxDepth(tn, 0)
//	matrix := make([][]int, maxLevel)
//	for level := 0; level < maxLevel; level++ {
//		matrix[level] = make([]int, 1<<level)
//	}
//	recursiveLoadToMatrix(&matrix, tn, 0, 0)
//	return matrix
//}

// Рекурсивная функция прохода по дереву и записи значений в матрицу
func recursiveLoadToMatrix(matrix *[][]int, root *TreeNode, level int, i int) {
	if root == nil || level >= len(*matrix) || root.Val == 0 {
		return
	}
	(*matrix)[level][i] = root.Val
	recursiveLoadToMatrix(matrix, root.Left, level+1, i*2)
	recursiveLoadToMatrix(matrix, root.Right, level+1, i*2+1)
}

// Печать дерева в древовдином виде
func (tn *TreeNode) String() string {
	maxVal := MaxValue(tn)
	minVal := MinValue(tn)
	maxValStr := strconv.Itoa(maxVal)
	minValStr := strconv.Itoa(minVal)
	maxLenElemen := max(len(maxValStr), len(minValStr))
	spaceString := strings.Repeat(" ", maxLenElemen)
	maxLevel := maxDepth(tn, 0)

	maxElements := 1 << (maxLevel - 1)
	sb := strings.Builder{}
	sb.WriteString("Max depth: " + strconv.Itoa(maxLevel) + "\n")
	matrix := ToMatrix(tn)
	//fmt.Println(matrix)
	for level := 0; level < maxLevel; level++ {
		cntElenentsOfLevel := 1 << level

		for i := 0; i < cntElenentsOfLevel; i++ {
			for j := 0; j < maxElements/cntElenentsOfLevel-1; j++ {
				sb.WriteString(spaceString)
			}
			d := strconv.Itoa(matrix[level][i])
			ld := len(d)
			sb.WriteString(strings.Repeat(" ", maxLenElemen-ld))
			sb.WriteString(d)
			for j := 0; j < maxElements/cntElenentsOfLevel; j++ {
				sb.WriteString(spaceString)
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func ToMatrix(tn *TreeNode) [][]int {
	maxLevel := maxDepth(tn, 0)
	matrix := make([][]int, maxLevel)
	for level := 0; level < maxLevel; level++ {
		matrix[level] = make([]int, 1<<level)
	}
	recursiveLoadToMatrix(&matrix, tn, 0, 0)
	return matrix
}

func MaxValue(root *TreeNode) int {
	if root == nil {
		return math.MinInt
	}
	return max(root.Val, MaxValue(root.Left), MaxValue(root.Right))
}

func MinValue(root *TreeNode) int {
	if root == nil {
		return math.MaxInt
	}
	return min(root.Val, MinValue(root.Left), MinValue(root.Right))
}

func maxDepth(root *TreeNode, null int) int {
	if root == nil || root.Val == null {
		return 0
	}
	return 1 + max(maxDepth(root.Left, null), maxDepth(root.Right, null))
}

func addTreeNode(i int, slice *[]int, null int) *TreeNode {
	if i >= len(*slice) {
		return nil
	}
	if (*slice)[i] == null {
		return nil
	}
	leftIdx, rightIdx := getIndex(i)
	return &TreeNode{Val: (*slice)[i], Left: addTreeNode(leftIdx, slice, null), Right: addTreeNode(rightIdx, slice, null)}
}

func SliceToTreeNode(slice []int, full bool, null int) *TreeNode {
	if full {
		return addTreeNode(0, &slice, null)
	} else {
		var fullSlice []int
		lengthSlice := len(slice)
		if lengthSlice == 0 {
			return nil
		}
		currPositionSlice := 0 // Текущая позиция в срезе
		for level := 0; currPositionSlice < lengthSlice; level++ {
			for i := 0; i < 1<<level; i++ {
				if parentIsNil(fullSlice, null) {
					fullSlice = append(fullSlice, null)
				} else {
					if currPositionSlice < lengthSlice {
						fullSlice = append(fullSlice, slice[currPositionSlice])
						currPositionSlice++
					} else {
						fullSlice = append(fullSlice, null)
					}
				}
			}
		}
		fmt.Println(fullSlice)
		return addTreeNode(0, &fullSlice, null)
	}
}

func SliceToTreeNodeFullMinInt(slice []int) *TreeNode {
	return SliceToTreeNode(slice, false, NULL)
}

func getIndex(x int) (left, right int) {
	if x == 0 {
		return 1, 2
	}
	var x_level int
	for x>>x_level > 0 {
		x_level++
	}
	pos_in_level := x - (1 << (x_level - 1)) + 1
	child_start_of_level := (1 << (x_level)) - 1
	left = child_start_of_level + pos_in_level*2
	right = child_start_of_level + pos_in_level*2 + 1
	return left, right
}

func parentIsNil(slice []int, null int) bool {
	l := len(slice)
	if l <= 1 {
		return false
	}

	return slice[l-(l+2)/2] == null
}

func TreeNodeToSlice(root *TreeNode) []int {

	start := 0
	arr := []*TreeNode{root}

	for start < len(arr) {
		if v := arr[start]; v != nil {
			arr = append(arr, v.Left)
			arr = append(arr, v.Right)
		}
		start++
	}

	result := make([]int, len(arr))
	for i, node := range arr {
		if node == nil {
			result[i] = NULL
		} else {
			result[i] = node.Val
		}
	}
	cntLastNull := 0
	for i := len(result) - 1; i >= 0 && result[i] == NULL; i-- {
		cntLastNull++
	}
	return result[:len(result)-cntLastNull]
}
