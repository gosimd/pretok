package regexp_test

import (
	"errors"
	"fmt"

	regex "github.com/gosimd/pretok/internal/regexp"
)

func ExampleCompile() {
	re, err := regex.Compile(`\p{N}{1,3}`) // Exact P07 string.
	if err != nil {
		panic(err)
	}
	input := "x1234y"
	spans, err := re.Split(input, nil)
	if err != nil {
		panic(err)
	}
	for _, span := range spans {
		fmt.Printf("%q [%d,%d)\n", input[span.Start:span.End], span.Start, span.End)
	}
	// Output:
	// "x" [0,1)
	// "123" [1,4)
	// "4" [4,5)
	// "y" [5,6)
}

func ExampleRegexp_Cursor() {
	re, err := regex.Compile(`\p{N}{1,3}`)
	if err != nil {
		panic(err)
	}
	input := "x１２34y"
	cursor, err := re.Cursor(input)
	if err != nil {
		panic(err)
	}
	for span, ok := cursor.Next(); ok; span, ok = cursor.Next() {
		fmt.Printf("%q [%d,%d)\n", input[span.Start:span.End], span.Start, span.End)
	}
	// Output:
	// "１２3" [1,8)
	// "4" [8,9)
}

func ExampleRegexp_Split() {
	re, err := regex.Compile(`\p{N}{1,3}`)
	if err != nil {
		panic(err)
	}
	spans := make([]regex.Span, 0, 8)
	for _, input := range []string{"x1234y", "１２3"} {
		spans, err = re.Split(input, spans[:0])
		if err != nil {
			panic(err)
		}
		fmt.Println(spans)
	}
	// Output:
	// [{0 1} {1 4} {4 5} {5 6}]
	// [{0 7}]
}

func ExampleCompile_errors() {
	_, err := regex.Compile(`a+`) // Valid general regex, unregistered in regex.
	fmt.Println(errors.Is(err, regex.ErrUnsupportedPattern))

	re, err := regex.Compile(`\p{N}{1,3}`)
	if err != nil {
		panic(err)
	}
	dst := []regex.Span{{Start: 10, End: 11}}
	dst, err = re.Split("\xff", dst)
	fmt.Println(errors.Is(err, regex.ErrInvalidUTF8), dst)

	var zero regex.Regexp
	_, err = zero.Cursor("123")
	fmt.Println(errors.Is(err, regex.ErrUninitialized))
	// Output:
	// true
	// true [{10 11}]
	// true
}
