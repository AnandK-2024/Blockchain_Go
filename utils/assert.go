package util

import (
	"log"
	"reflect"
)


// As DeepEqual traverses the data values it may find a cycle. The
// second and subsequent times that DeepEqual compares two pointer
// values that have been compared before, it treats the values as
// equal rather than examining the values to which they point.
// This ensures that DeepEqual terminates.
func AssertEqual(a, b any) {
	if !reflect.DeepEqual(a, b) {
		log.Fatalf("ASSERTION: %+v != %+v", a, b)
	}
}
