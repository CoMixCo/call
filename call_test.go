package call

import (
	"fmt"
	"testing"
)

//define demo struct
type Show struct{}

func (s *Show) GetCode(code string) string {
	return code
}

func TestRun(t *testing.T) {

	//step1: init
	c := NewCall()

	//step2: register handle
	c.Register("bb", func() any {
		return &Show{}
	})

	//step3: call struct by string
	result, err := c.Invok("bb", "GetCode", "7ELFHG3HIOIPPOR4UV4YPLHXMPLG6CSM")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, rs := range result {
		fmt.Println(rs.Interface())
	}

}
