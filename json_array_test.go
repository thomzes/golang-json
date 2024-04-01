package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Thomas",
		MiddleName: "Ardi",
		LastName:   "Ansah",
		Hobbies:    []string{"Gaming", "Traveling", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Thomas","MiddleName":"Ardi","LastName":"Ansah","Age":0,"Married":false,"Hobbies":["Gaming","Traveling","Coding"]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)

}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName: "Thomas",
		Addresses: []Address{
			{
				Street:     "Jl.Panjang",
				Country:    "Indonesia",
				PostalCode: "111",
			},
			{
				Street:     "Jl.Pendek",
				Country:    "Singapore",
				PostalCode: "222",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Thomas","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jl.Panjang","Country":"Indonesia","PostalCode":"111"},{"Street":"Jl.Pendek","Country":"Singapore","PostalCode":"222"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jl.Panjang","Country":"Indonesia","PostalCode":"111"},{"Street":"Jl.Pendek","Country":"Singapore","PostalCode":"222"}]`

	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {

	addresses := []Address{
		{
			Street:     "Jl.Panjang",
			Country:    "Indonesia",
			PostalCode: "111",
		},
		{
			Street:     "Jl.Pendek",
			Country:    "Singapore",
			PostalCode: "222",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))

}
