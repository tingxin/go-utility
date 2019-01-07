package test

import (
	"testing"

	"github.com/tingxin/go-utility/array"
)

type Person struct {
	Name string
	Age  int
}

var (
	stringTestCases []string
	intTestCases    []int
	pTestCases      []*Person
)

func init() {
	stringTestCases = make([]string, 0)
	stringTestCases = append(stringTestCases, "hello")
	stringTestCases = append(stringTestCases, "barry")
	stringTestCases = append(stringTestCases, "nio")
	stringTestCases = append(stringTestCases, "it")

	intTestCases = make([]int, 0)
	intTestCases = append(intTestCases, 122)
	intTestCases = append(intTestCases, 3234)
	intTestCases = append(intTestCases, 199)
	intTestCases = append(intTestCases, 234)

	pTestCases = make([]*Person, 0)
	pTestCases = append(pTestCases, &Person{Name: "huangye", Age: 33})
	pTestCases = append(pTestCases, &Person{Name: "barry", Age: 31})
	pTestCases = append(pTestCases, &Person{Name: "dandan", Age: 27})
	pTestCases = append(pTestCases, &Person{Name: "xiping", Age: 31})

}

func TestContain(t *testing.T) {

	objArray := array.StringsToInterfaces(stringTestCases)

	contain := array.Contain(objArray, "barry")
	assertEqual(t, contain, true)

	contain = array.Contain(objArray, "xxxx")
	assertEqual(t, contain, false)

	objArray = array.IntsToInterfaces(intTestCases)

	contain = array.Contain(objArray, 199)
	assertEqual(t, contain, true)

	contain = array.Contain(objArray, 1212)
	assertEqual(t, contain, true)

}

func TestFirst(t *testing.T) {

	objArray := array.StringsToInterfaces(stringTestCases)

	compare1 := func(a interface{}) bool {
		if a == "nio" {
			return true
		}
		return false
	}

	first := array.First(objArray, compare1)
	assertEqual(t, first, "nio")

	objArray = array.IntsToInterfaces(intTestCases)

	compare2 := func(a interface{}) bool {
		if a == 3234 {
			return true
		}
		return false
	}

	first = array.First(objArray, compare2)
	assertEqual(t, first, 3234)

	b := make([]interface{}, len(pTestCases), len(pTestCases))
	for i := range pTestCases {
		b[i] = pTestCases[i]
	}
	compare3 := func(a interface{}) bool {
		p := a.(*Person)
		if p.Name == "barry" && p.Age == 31 {
			return true
		}
		return false
	}
	first = array.First(b, compare3)
	assertEqual(t, first, pTestCases[1])

}

func assertEqual(t *testing.T, actually interface{}, expected interface{}) {
	if actually != expected {
		t.Fatalf("Expected: %v ===> Actually: %v\n", expected, actually)
	}
}
