package main

import "study/pets"

func main() {
	dog := pets.NewDog("Рекс")
	eagle := pets.NewEagle("Акела")

	hero := pets.NewHero("Робинзон Крузо")
	hero.AddPet(dog)
	hero.AddPet(eagle)
	hero.ShowName()
	// hero.ShowPets()
	hero.GetPetComand(eagle.Name())
	hero.GetPetComand(dog.Name())

}
