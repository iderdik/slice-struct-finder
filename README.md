# slice-struct-finder
##Find items in a slice of structs of any type

###Sample usage:

```
type Person struct {
	id        string
	firstName string
	lastName  string
}

func (p Person) GetKey() string {
	return p.id
}

people = make([]Person, 4, 4)
people[0] = Person{id: "101", firstName: "John", lastName: "Smith"}
people[1] = Person{id: "23", firstName: "Michael", lastName: "Jordan"}
people[2] = Person{id: "007", firstName: "James", lastName: "Bond"}
people[3] = Person{id: "008", firstName: "Jason", lastName: "Bourne"}

if person, err := FindByKey(&people, "007"); err== nil {
	fmt.Println(person.firstName, person.lastName)
}
```

Note: Make sure you provide a `GetKey() string` method for the struct you'd like to make searchable.
