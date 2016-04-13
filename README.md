mgos
=====

**M**ap **GO** to **S**truct



Installation
-------------

```
go get github.com/ushios/mgos
```

Documentation
-------------

[![GoDoc](https://godoc.org/github.com/ushios/mgos?status.svg)](https://godoc.org/github.com/ushios/mgos)

Example
========

### Using url.Values 

```go
type Person struct {
	FirstName  string `mgos:"first_name"`
	LastName   string `mgos:"last_name"`
	Age        int    `mgos:"age"`
	BirthMonth int    `mgos:"birth_month"`
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
```
