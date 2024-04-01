package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestCode(t *testing.T) {
	LogJson("Thomas")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"Thomas", "ardiansah", "uhuy"})
}
