package game

import (
	"fmt"
	"math/rand"
)

type Attacker interface {
	Attack() int  // сколько урона наносит
	Name() string // имя существа или оружия
}

// Типы воинов

type Warrior struct {
	name string
}

func NewWarrior(name string) *Warrior {
	return &Warrior{name: name}
}

func (w Warrior) Attack() int {
	return rand.Intn(11) + 10
}
func (w Warrior) Name() string {
	return w.name
}

type Archer struct {
	name string
}

func NewArcher(name string) *Archer {
	return &Archer{name: name}
}

func (a Archer) Attack() int {
	return rand.Intn(6) + 25
}
func (a Archer) Name() string {
	return a.name
}

type Mage struct {
	name string
}

func NewMage(name string) *Mage {
	return &Mage{name: name}
}

func (m Mage) Attack() int {
	return rand.Intn(10) + 18
}
func (m Mage) Name() string {
	return m.name
}

// типы оружия

type Sword struct {
	name string
}

func NewSword(name string)*Sword {
	return &Sword{name: name}
}

func (s Sword) Attack() int {
	return rand.Intn(10) + 15
}
func (s Sword) Name() string {
	return s.name
}

type Bow struct {
	name string
}

func NewBow(name string) *Bow {
	return &Bow{name: name}
}

func (b Bow) Attack() int {
	return rand.Intn(8) + 18
}
func (b Bow) Name() string {
	return b.name
}

type BattleModule struct {
	firstWarrior  Attacker
	secondWarrior Attacker
}

func (b BattleModule) Fight() string {
	damageFirst := b.firstWarrior.Attack()
	damageSecond := b.secondWarrior.Attack()

	fmt.Printf("%s ударил - %d\n", b.firstWarrior.Name(), damageFirst)
	fmt.Printf("%s ударил - %d\n", b.secondWarrior.Name(), damageSecond)
	if damageFirst > damageSecond {
		return fmt.Sprintf("Побеждает %s", b.firstWarrior.Name())
	} else {
		return fmt.Sprintf("Побеждает %s", b.secondWarrior.Name())
	}
}

func NewBattleModule(a, b Attacker) *BattleModule {
	return &BattleModule{
		firstWarrior:  a,
		secondWarrior: b,
	}
}
