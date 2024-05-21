package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number with one digit after the decimal.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox returns a string describing the NumberBox, ensuring type safety.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber returns the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		value, _ := strconv.Atoi(fnb.Value())
		return value
	default:
		return 0
	}
// 	if fancyNum, ok := fnb.(FancyNumber); ok {
//        	if output, err := strconv.Atoi(fancyNum.Value()); err == nil {
//             return output
//         }
//     }
//    return 0
}

// DescribeFancyNumberBox returns a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %d.0", ExtractFancyNumber(fnb))
}

// DescribeAnything returns a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch value := i.(type) {
	case float64:
		return DescribeNumber(value)
	case int:
		return DescribeNumber(float64(value))
	case NumberBox:
		return DescribeNumberBox(value)
	case FancyNumberBox:
		return DescribeFancyNumberBox(value)
	default:
		return "Return to sender"
	}
}
