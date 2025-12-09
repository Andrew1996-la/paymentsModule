package spellsystem

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Spell interface {
	Name() string
	Cast() int
	Type() string
}

type Mage struct {
	Spells map[string]Spell
}

func NewMage() *Mage {
	return &Mage{
		Spells: map[string]Spell{},
	}
}

func (m *Mage) LearnSpell(spell Spell) {
	m.Spells[spell.Name()] = spell
}

func (m Mage) ShowMySpells() {
	pp.Println(m.Spells)
	for _, value := range m.Spells {
		fmt.Printf("Spell называется %s\n", value.Name())
		fmt.Printf("Spell имеет тип %s\n", value.Type())
		fmt.Printf("Spell наносит урон/лечит %d\n", value.Cast())
		fmt.Printf("")
	}
}

func (m *Mage) CastSpell(name string) {
	spell, ok := m.Spells[name]
	if !ok {
		fmt.Println("spell not found")
		return
	}
	
	switch spell.Type() {
	case "damage":
		fmt.Printf("Mage выполнил %s и нанес %d ед урона\n", spell.Name(), spell.Cast())
	case "healing":
		fmt.Printf("Mage выполнил %s и вылечил %d ед здоровья\n", spell.Name(), spell.Cast())
	default:
		fmt.Println("Неизвестный тип заклинания")
	}
}
