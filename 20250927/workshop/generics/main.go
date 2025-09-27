package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Slice[T any] []T

func (s Slice[T]) Filter(f func(T) bool) Slice[T] {
	var result Slice[T]
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[a, b any](s Slice[a], f func(a) b) Slice[b] {
	result := make(Slice[b], len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

type Container[T fmt.Stringer] struct {
	items []T
}

// Add メソッド（完成済み）
func (c *Container[T]) Add(item T) {
	c.items = append(c.items, item)
}

// PrintAll メソッド（完成済み）
func (c *Container[T]) PrintAll() {
	for _, item := range c.items {
		fmt.Println(item.String())
	}
}

// Person型
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}

// Product型
type Product struct {
	Name  string
	Price float64
}

func (p Product) String() string {
	return fmt.Sprintf("%s: $%.2f", p.Name, p.Price)
}

type User struct {
	Name string
	Age  int
}

func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	return json.Unmarshal(data, &aux)
}

func Unmarshal[T any, PT Unmershaller[T]](data []byte) (T, error) {
	var v T
	err := PT(&v).UnmarshalJSON(data)
	return v, err
}

type Unmershaller[T any] interface {
	*T
	json.Unmarshaler
}

func main() {
	/*
		v1
	*/
	ints := Slice[int]{1, 2, 3, 4, 5}
	evens := ints.Filter(func(i int) bool { return (i)%2 == 0 })
	fmt.Printf("%#v\n", evens) // Output: main.Slice{2, 4}

	v := Map(ints, func(v int) string { return strconv.Itoa(v) })
	fmt.Printf("%#v\n", v) // Output: main.Slice{"1", "2", "3", "4", "5"}

	/*
		v2
	*/
	// Person用のContainer
	// people := Container[any]{} // TODO: any を適切な制約に変更
	people := Container[Person]{}
	people.Add(Person{"Alice", 30})
	people.Add(Person{"Bob", 25})
	// people.Add(Product{"Bob", 25})
	fmt.Println("People:")
	people.PrintAll()

	// Product用のContainer
	products := Container[Product]{}
	products.Add(Product{"Laptop", 999.99})
	products.Add(Product{"Mouse", 25.50})
	fmt.Println("\nProducts:")
	products.PrintAll()

	/*
		v3
	*/
	data := []byte(`{"Name": "Alice", "Age": 30}`)
	user, err := Unmarshal[User](data)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Name, user.Age) // Output: Alice 30

}
