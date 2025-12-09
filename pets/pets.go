package pets

import "fmt"

type Pet interface {
	Name() string
	Action() string
}

type Dog struct {
	name string
}

func NewDog(name string) *Dog {
	return &Dog{name: name}
}

func (d Dog) Name() string {
	return d.name
}
func (d Dog) Action() string {
	return "Лаять и кусать"
}

type Eagle struct {
	name string
}

func NewEagle(name string) *Eagle {
	return &Eagle{name: name}
}

func (e Eagle) Name() string {
	return e.name
}
func (e Eagle) Action() string {
	return "Летать и охотиться"
}

type Hero struct {
	name string
	pets map[string]Pet
}

func NewHero(name string) *Hero {
	return &Hero{
		name: name,
		pets: make(map[string]Pet),
	}
}

func (h Hero) AddPet(pet Pet) {
	h.pets[pet.Name()] = pet
}
func (h Hero) ShowName() {
	fmt.Println("Мое имя", h.name)
}
func (h Hero) GetPetComand(petName string) {
	currentPet, ok := h.pets[petName]

	if !ok {
		fmt.Println("животное не наидено")
		return
	}

	fmt.Println(currentPet.Name(), "выполняет: ", currentPet.Action())
}
func (h Hero) ShowPets() {
	if len(h.pets) == 0 {
		fmt.Println("Еще нет животных")
	}

	for _, pet := range h.pets {
		fmt.Printf("Животное - %s, оно умеет %s\n", pet.Name(), pet.Action())
	}
}
