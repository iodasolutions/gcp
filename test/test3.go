package main

import (
	"fmt"
	"github.com/iodasolutions/xbee-common/newfs"
	"log"
)

type XbeeElement[T any] struct {
	Provider map[string]interface{} `json:"provider,omitempty"`
	Element  T                      `json:"element,omitempty"`
}

type Volume struct {
	Name string
	Size int64
}

func main() {
	a := XbeeElement[Volume]{
		Provider: map[string]interface {
		}{
			"hello": "Eric",
		},
		Element: Volume{
			Name: "v1",
			Size: int64(100),
		},
	}

	f := newfs.TmpDir().RandomFile()
	f.Save(a)

	aa, err := newfs.Unmarshal[XbeeElement[Volume]](f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(aa)
}
