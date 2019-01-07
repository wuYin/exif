package exif

import (
	"fmt"
	"os"
	"testing"
)

// ok
func TestRead(t *testing.T) {
	f, _ := os.Open("cup.jpg")
	r := NewExifReader()
	if err := r.ReadContent(f); err != nil {
		t.Fatalf("test exif read failed: %v", err)
	}

	// bingo
	for _, tag := range r.MainTags {
		fmt.Printf("%v\t%v\n", tag.Desc, tag.Value)
	}
	fmt.Println()
	for _, tag := range r.GPSTags {
		fmt.Printf("%v\t%v\n", tag.Desc, tag.Value)
	}
}
