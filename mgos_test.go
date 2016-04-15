package mgos

import (
	"fmt"
	"net/url"
	"testing"
)

type Gender uint8

const (
	Male Gender = iota + 1
	Female
)

type City string

const (
	Tokyo   City = "tokyo"
	NewYork City = "new york"
)

type Person struct {
	FirstName  string `mgos:"first_name"`
	LastName   string `mgos:"last_name"`
	Age        int    `mgos:"age"`
	BirthMonth int    `mgos:"birth_month"`
	Gender     Gender `mgos:"gender"`
	City       City   `mgos:"city"`
	IsAlive    bool   `mgos:"is_alive"`
	Hobby      string
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

		if dest.BirthMonth != expect.BirthMonth {
			t.Errorf("BirthMonth expected (%d) but (%d)", expect.BirthMonth, dest.BirthMonth)
		}

		if dest.Gender != expect.Gender {
			t.Errorf("Gender expected (%d) but (%d)", expect.Gender, dest.Gender)
		}

		if dest.City != expect.City {
			t.Errorf("City expected (%s) but (%s)", expect.City, dest.City)
		}
	}

	test("first_name=Tanaka&last_name=Satoshi&age=18&gender=1&city=tokyo&is_alive=1", Person{
		FirstName:  "Tanaka",
		LastName:   "Satoshi",
		Age:        18,
		BirthMonth: 0,
		Gender:     Male,
		City:       Tokyo,
		IsAlive:    true,
	})

	test("first_name=Inoue&last_name=Shingo&age=19&city=mexico&&is_alive=0", Person{
		FirstName:  "Inoue",
		LastName:   "Shingo",
		Age:        19,
		BirthMonth: 0,
		Hobby:      "tenis",
		City:       "mexico",
		IsAlive:    false,
	})

	test("first_name=John&last_name=Handen&age=1", Person{
		FirstName:  "John",
		LastName:   "Handen",
		Age:        1,
		BirthMonth: 0,
		Hobby:      "video games",
		IsAlive:    false,
	})

}

func ExampleFromGetter() {
	v, _ := url.ParseQuery("first_name=Tanaka&last_name=Satoshi&age=18&birth_month=January&is_alive=1")
	dest := Person{}

	FromGetter(v, &dest)

	fmt.Println(dest.FirstName)
	fmt.Println(dest.LastName)
	fmt.Println(dest.Age)
	fmt.Println(dest.BirthMonth)
	fmt.Println(dest.IsAlive)
	// Output:
	// Tanaka
	// Satoshi
	// 18
	// 0
	// true
}
