package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	str := "jjxuhuade，徐华德"
	fmt.Println(bytes.Contains([]byte(str), []byte("jj1")))
	fmt.Println(bytes.ContainsAny([]byte(str), "g"))
	fmt.Println(bytes.ContainsRune([]byte(str), '徐'))
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each ru
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))

	fmt.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}))

	fmt.Println(bytes.HasPrefix([]byte(str), []byte("jjxuhuade")))

	fmt.Println(bytes.Index([]byte(str), []byte("xu")))
	fmt.Println(bytes.IndexByte([]byte(str), 'x'))

	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s, []byte(",")))

	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A'
		case r >= 'a' && r <= 'z':
			return 'a'
		}
		return r
	}
	fmt.Printf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))

	fmt.Println()
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("jj"), 3))

	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2))
	fmt.Println()
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2))
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	fmt.Printf("%s", bytes.Title([]byte("her royal highness")))

}
