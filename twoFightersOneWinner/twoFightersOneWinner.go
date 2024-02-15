package main

import "fmt"

// The Fighter type represents a character with a name, health, and damage per attack attributes.
// @property {string} Name - The Name property is a string that represents the name of the fighter.
// @property {int} Health - The "Health" property represents the current health points of the fighter.
// It is an integer value that indicates how much damage the fighter can withstand before being
// defeated.
// @property {int} DamagePerAttack - The "DamagePerAttack" property represents the amount of damage a
// fighter can inflict with each attack.
type Fighter struct {
    Name            string
    Health          int
    DamagePerAttack int
}

// The function "DeclareWinner" determines the winner of a fight between two fighters based on their
// health and damage per attack.
func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	for {
		if firstAttacker == fighter1.Name {
			fighter2.Health -= fighter1.DamagePerAttack
			firstAttacker = fighter2.Name
		} else {
			fighter1.Health -= fighter2.DamagePerAttack
			firstAttacker = fighter1.Name
		}
			
		if fighter1.Health <= 0 {
			return fighter2.Name
		}
		
		if fighter2.Health <= 0 {
			return fighter1.Name
		}
	}
}

func main() {
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew"))
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"))
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry"))
}
