package snet

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	type stu struct {
		Id   int
		Name string
	}
	s := stu{
		Id:   1,
		Name: "json",
	}
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	x := stu{}

	err = json.Unmarshal(b, &x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)
}
