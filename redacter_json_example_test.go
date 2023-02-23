package redacter_test

import (
	"fmt"

	"github.com/jcyamacho-go/redacter"
)

type Person struct {
	Name      string
	Age       int
	Password  string
	Addresses []Address
}

type Address struct {
	Country string
	City    string
}

func ExampleRedacter_JSON() {
	p := Person{
		Name:     "name",
		Age:      30,
		Password: "password",
		Addresses: []Address{
			{
				Country: "country-1",
				City:    "city-1",
			},
			{
				Country: "country-2",
				City:    "city-2",
			},
		},
	}

	res, err := redacter.New().
		WithFields("password", "country").
		WithStrict(false).
		WithRemove(true).
		JSON(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	// Output:
	// {"Addresses":[{"City":"city-1"},{"City":"city-2"}],"Age":30,"Name":"name"}
}

func ExampleRedacter_JSONRaw() {
	raw := []byte(`{
		"name": "name",
		"secrets": [
			{ "id": 1, "name": "secret-1" },
			{ "id": 2, "name": "secret-2" },
			{ "id": 3, "name": "secret-2", "shared": true }
		]
	}`)

	res, err := redacter.New().
		WithFields("name").
		JSONRaw(raw)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	// Output:
	// {"name":"[REDACTED]","secrets":[{"id":1,"name":"[REDACTED]"},{"id":2,"name":"[REDACTED]"},{"id":3,"name":"[REDACTED]","shared":true}]}
}
