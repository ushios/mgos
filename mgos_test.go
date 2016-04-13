package mgos

import (
	"fmt"
	"net/url"
	"testing"
)

type Person struct {
	FirstName  string `mgos:"first_name"`
	LastName   string `mgos:"last_name"`
	Age        int    `mgos:"age"`
	BirthMonth int    `mgos:"birth_month"`
}

func TestFromGetterUsingURLValues(t *testing.T) {
	test := func(q string, expect Person) {
		v, err := url.ParseQuery(q)
		if err != nil {
			t.Errorf("ParseQuery got error: %s", err)
		}

		dest := Person{}
		FromGetter(v, &dest)

		if dest.FirstName != expect.FirstName {
			t.Errorf("FirstName expected (%s) but (%s)", expect.FirstName, dest.FirstName)
		}

		if dest.LastName != expect.LastName {
			t.Errorf("LastName expected (%s) but (%s)", expect.LastName, dest.LastName)
		}

		if dest.Age != expect.Age {
			t.Errorf("Age expected (%d) but (%d)", expect.Age, dest.Age)
		}
	}

	test("first_name=Tanaka&last_name=Satoshi&age=18", Person{
		FirstName: "Tanaka",
		LastName:  "Satoshi",
		Age:       18,
	})
}

func ExampleFromGetter() {
	v, _ := url.ParseQuery("first_name=Tanaka&last_name=Satoshi&age=18&birth_month=January")
	dest := Person{}

	FromGetter(v, &dest)

	fmt.Println(dest.FirstName)
	fmt.Println(dest.LastName)
	fmt.Println(dest.Age)
	fmt.Println(dest.BirthMonth)
	// Output:
	// Tanaka
	// Satoshi
	// 18
	// 0
}
