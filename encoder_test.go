package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Thomas",
		MiddleName: "Ardi",
		LastName:   "Ansah",
	}

	_ = encoder.Encode(customer)

	fmt.Println(customer)

}
