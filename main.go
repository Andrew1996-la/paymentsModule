package main

import (
	spellsystem "study/spellSystem"
)

func main() {
	fireball := spellsystem.NewFireball()
	healing := spellsystem.NewHealing()
	mage := spellsystem.NewMage()
	mage.LearnSpell(fireball)
	mage.LearnSpell(healing)
	// mage.ShowMySpells()
	mage.CastSpell("Fireball")
	mage.CastSpell("Healing")
}
