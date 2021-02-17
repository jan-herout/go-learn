package main

import "fmt"

// Animal represents anything, which has legs, and can walk.
type Animal struct {
	Legs   int    // number of legs this animal has
	kind   string // what kind of the animal this is?
	atLat  int    // latitude where the animal is
	atLong int    // longitude where the animal is
}

// WalkTo moves the animal to latitude and longitude.
func (a *Animal) WalkTo(lat, lng int) {
	a.atLat, a.atLong = lat, lng
}

// Dog is an Animal, which can be named, and can also bark!
type Dog struct {
	Name        string // everybody names their favourite puppy...
	IsPureBreed bool   // is the dog a pure blood?
	Animal             // COMPOSITION: Dog is also an Animal!
}

// Bark is a Dog method. Pure breed's bark is slightly more affectionate.
func (d *Dog) Bark() string {
	if d.IsPureBreed {
		return "Haf!"
	}
	return "Štěk!"
}

// NewDog returns an Animal which is also a Dog.
func NewDog(name string, isPure bool) Dog {
	return Dog{
		Name:        name,
		IsPureBreed: isPure,
		Animal: Animal{
			Legs: 4,
			kind: "dog",
		},
	}
}

// String implements a fmt.Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("lat=%d, long=%d", a.atLat, a.atLong)
}

func main() {
	d := NewDog("Fido", false)
	d.WalkTo(1, -1)
	fmt.Println(d, ":", d.Bark())
}
