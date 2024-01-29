package bst

import (
	"testing"
)

type insertTest[T any] struct {
	key      int
	data     T
	expected error
}

var testCase = []insertTest[string]{
	{425, "sdfaefaf", nil},
	{542, "kjklj", nil},
	{453, "jlkhkjbjkbjb", nil},
	{532, "joooinn", nil},
	{213, "oxjcdjo", nil},
	{12, "nkn nkl", nil},
	{899, "jjjooooo", nil},
	{313, "oooooooooo", nil},
	{345, "jljop", nil},
	{298, "iouiui", nil},
	{666, "iopiouji", nil},
	{987, "jijlik", nil},
	{136, "jnnnn", nil},
	{700, "[qqqq]", nil},
	{803, "qwefjkj ihihqefhj", nil},
	{58, "0xheyyyyyyyyyyy", nil},
	{901, "0q ceee", nil},
	{110, "no", nil},
	{1245, "dddd", nil},
	{679, "qw", nil},
	{404, "htttttttttttttttttttttttttttt", nil},
	{505, "tterert", nil},
	{887, "sfsedf", nil},
	{323, "qqqqqqqqqqqqqqqqqqqqqqqq", nil},
}

func TestInsert(t *testing.T) {
	tr := New[string]()

	for _, v := range testCase {
		if output := tr.Insert(uint64(v.key), v.data); output != v.expected {
			t.Errorf("Output %q not equal to expected %q", output, v.expected)
		}

	}
}
