package main

import "fmt"

type heroface interface {
	attribute()
	setAttackType(attackType)
}
type agility struct {
	attackType attackType
}

func (a *agility) attribute() {
	fmt.Println("agility")
	a.attackType.knowAttackType()
}

func (a *agility) setAttackType(t attackType) {
	a.attackType = t
}
type strength struct {
	attackType attackType
}

func (s *strength) attribute() {
	fmt.Println("strength")
	s.attackType.knowAttackType()
}

func (s *strength) setAttackType(t attackType) {
	s.attackType = t
}
type intellect struct {
	attackType attackType
}

func (i *intellect) attribute() {
	fmt.Println("intellect")
	i.attackType.knowAttackType()
}

func (i *intellect) setAttackType(t attackType) {
	i.attackType = t
}
type attackType interface {
	knowAttackType()
}


type dalnik struct {
}

func (d *dalnik) knowAttackType() {
	fmt.Println("dalnik")
}
type blizhnik struct {
}

func (b *blizhnik) knowAttackType() {
	fmt.Println("blizhnik")
}
func main() {
	terrorblade:=&agility{}
	axe :=&strength{}
	tinker :=&intellect{}

	dalnik :=&dalnik{}
	blizhnik :=&blizhnik{}

	fmt.Println("Terrorblade")
	terrorblade.setAttackType(blizhnik)
	terrorblade.setAttackType(dalnik)
	terrorblade.attribute()
	fmt.Println()
	fmt.Println("Axe")
	axe.setAttackType(blizhnik)
	axe.attribute()
	fmt.Println()
	fmt.Println("Tinker")
	tinker.setAttackType(dalnik)
	tinker.attribute()
	fmt.Println()
}
