package finder_test

import (
	. "github.com/iderdik/slice-struct-finder/finder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Person struct {
	id        string
	firstName string
	lastName  string
}
type Dog struct {
	name string
}

func (p Person) GetKey() string {
	return p.id
}

var _ = Describe("Finder", func() {
	var (
		people []Person
		dogs   []Dog
	)

	BeforeSuite(func() {
		people = make([]Person, 4, 4)
		people[0] = Person{id: "101", firstName: "John", lastName: "Smith"}
		people[1] = Person{id: "23", firstName: "Michael", lastName: "Jordan"}
		people[2] = Person{id: "007", firstName: "James", lastName: "Bond"}
		people[3] = Person{id: "008", firstName: "Jason", lastName: "Bourne"}

		dogs = make([]Dog, 2, 2)
		dogs[0] = Dog{name: "Lassie"}
		dogs[1] = Dog{name: "Old Yeller"}

	})

	Describe("Searching a valid slice", func() {

		It("should find Persons that exist", func() {
			person, err := FindByKey(&people, "007")

			Expect(person.(Person).firstName).To(Equal("James"))
			Expect(err).NotTo(HaveOccurred())
		})

		It("should not find Persons that don't exist", func() {
			_, err := FindByKey(&people, "42")

			Expect(err).Should(MatchError("Key not found"))
		})
	})

	Describe("Searching an invalid slice", func() {

		It("should find Persons that exist", func() {
			_, err := FindByKey(&dogs, "Lassie")

			Expect(err).To(HaveOccurred())
		})
	})

})
