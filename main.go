package main

import (
	"study/game"

	"github.com/k0kubun/pp"
)

func main() {
	warrior := game.NewWarrior("Aragorn")
	mage := game.NewMage("Gendalf")

	battleModuel := game.NewBattleModule(warrior, mage)

	result := battleModuel.Fight()
	pp.Println(result)
}
