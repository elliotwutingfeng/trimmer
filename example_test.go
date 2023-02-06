package trimmer

import (
	"fmt"
	"strings"

	"github.com/karlseguin/intset"
)

func Example_basic() {
	const charsToTrim string = "@👍🏽新 "
	var cutset *intset.Rune = MakeRuneSet(charsToTrim)

	fmt.Println(fastTrim("", cutset, TrimBoth))
	fmt.Println(strings.Trim("", "@👍🏽新 "))

	fmt.Println(fastTrim(" ", cutset, TrimBoth))
	fmt.Println(strings.Trim(" ", "@👍🏽新 "))

	fmt.Println(fastTrim("@b👍🏽新", cutset, TrimBoth))
	fmt.Println(strings.Trim("@b👍🏽新", "@👍🏽新 "))

	fmt.Println(fastTrim("@b👍🏽新", cutset, TrimLeft))
	fmt.Println(strings.TrimLeft("@b👍🏽新", "@👍🏽新 "))

	fmt.Println(fastTrim("@b👍🏽新", cutset, TrimRight))
	fmt.Println(strings.TrimRight("@b👍🏽新", "@👍🏽新 "))

	fmt.Println(fastTrim("@b👍新", cutset, TrimRight))
	fmt.Println(strings.TrimRight("@b👍新", "@👍🏽新 "))
	//Output:
	//
	//
	//b
	//b
	//b👍🏽新
	//b👍🏽新
	//@b
	//@b
	//@b
	//@b
}
