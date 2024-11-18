package textbasics

import (
	"fmt"
	"unicode/utf8"
)

func Dummy() {
	rr := 'a'
	bb := utf8.RuneLen(rr)
	msg := fmt.Sprintf("You need %d byte(s) to encode the Rune:  `%c`", bb, rr)
	fmt.Println(msg)
	fmt.Println(len("emeka😈")) // this give u the number of bytes!
	fmt.Println(utf8.RuneCountInString("emeka😈"))
}
