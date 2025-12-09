package spellsystem

type Fireball struct{}

func NewFireball() *Fireball {
	return &Fireball{}
}

func (f Fireball) Name() string {
	return "Fireball"
}
func (f Fireball) Cast() int {
	return 20
}
func (f Fireball) Type() string {
	return "damage"
}

// ----------------------------------

type Healing struct{}

func NewHealing() *Healing {
	return &Healing{}
}

func (h Healing) Name() string {
	return "Healing"
}
func (h Healing) Cast() int {
	return 10
}
func (h Healing) Type() string {
	return "healing"
}
