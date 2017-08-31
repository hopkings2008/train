package test

import (
	//"fmt"
	"strings"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&TestSuite{})

type TestSuite struct {
}

func (cs *TestSuite) SetUpSuite(c *check.C) {
}

func (cs *TestSuite) TearDownSuite(c *check.C) {
}

func (cs *TestSuite) TestBasicq(c *check.C) {
	array := []int{4, 5, 2, 1, 3, 6}
	qsort(array)
	for i := 0; i < len(array); i++ {
		c.Logf("%d: %d", i, array[i])
	}
}

func (cs *TestSuite) TestBasicRevert(c *check.C) {
	str := "this is a test program."
	strs := strings.Split(str, " ")
	revert(&strs, c)
	c.Logf("strs: %v", strs)
}

func qsort(array []int) {
	start := -1
	size := len(array)
	if 0 == size {
		return
	}

	t := array[size-1]

	for i := 0; i < size; i++ {
		if array[i] < t {
			start += 1
			tmp := array[start]
			array[start] = array[i]
			array[i] = tmp
		}
	}
	start += 1
	tmp := array[start]
	array[start] = t
	array[size-1] = tmp

	if start-1 >= 0 {
		qsort(array[:start])
	}
	if start+1 < size {
		qsort(array[start+1:])
	}
}

func revert(strs *[]string, c *check.C) {
	num := len(*strs)
	if num <= 1 {
		return
	}
	var left []string
	var right []string
	left = append(left, (*strs)[:num/2]...)
	right = append(right, (*strs)[num/2:]...)
	revert(&left, c)
	revert(&right, c)
	//c.Logf("strs: %v, left: %v, right: %v", strs, left, right)
	*strs = right
	*strs = append(*strs, left...)
}
