# slice-struct-finder
##Find items in a slice of structs of any type

###Sample usage:

```
people = make([]Person, 4, 4)
people[0] = Person{id: "101", firstName: "John", lastName: "Smith"}
people[1] = Person{id: "23", firstName: "Michael", lastName: "Jordan"}
people[2] = Person{id: "007", firstName: "James", lastName: "Bond"}
people[3] = Person{id: "008", firstName: "Jason", lastName: "Bourne"}

if person, err := FindByID(&people, "007"); err== nil {
	fmt.Println(person.firstName, person.lastName)
}
```
