package main

import "fmt"

type hero struct {
	name, mainAttribute string
	skills []string
	items [5]string
	backPack [2]string
	hiddenItems [5]string
	farm, support, tank,carry bool
}
type heroBuilder struct{
	hero hero
}
func (value *heroBuilder) Name(name string) *heroBuilder{
	fmt.Println(name)

	value.hero.name= name
	return value
}
func (value *heroBuilder) MainAttribute(mainAttribute string) *heroBuilder{

	value.hero.mainAttribute= mainAttribute
	return value

}
func (value *heroBuilder) Skills(skills []string) *heroBuilder{
	value.hero.skills= skills
	return value

}
func (value *heroBuilder) Items(items [5]string) *heroBuilder{

	fmt.Println(items)
	value.hero.items= items
	return value

}
func (value *heroBuilder) HiddenItems(hiddenItems [5]string) *heroBuilder{

	value.hero.hiddenItems= hiddenItems
	return value

}
func (value *heroBuilder) BackPack(backPack [2]string) *heroBuilder{

	value.hero.backPack= backPack
	return value

}
func (value *heroBuilder) Farm(farm bool) *heroBuilder{

	value.hero.farm= farm
	return value

}
func (value *heroBuilder) Carry(carry bool) *heroBuilder{

	value.hero.carry= carry
	return value

}
func (value *heroBuilder) Support(support bool) *heroBuilder{

	value.hero.support = support
	return value

}
func (value *heroBuilder) Tank(tank bool) *heroBuilder{
	value.hero.tank = tank
	return value
}
func createHero(hero *hero) {

}
type build func(builder *heroBuilder)
func CreateHero(action build)  {

	builder :=heroBuilder{}
	fmt.Println(&builder.hero)
	action(&builder)
	createHero(&builder.hero)
}
func main(){
	CreateHero(func(value *heroBuilder) {
		value.Name("Axe").
			MainAttribute("Strength").
			Tank(true).
			Skills([]string{"Berserker's call", "Battle hunger", "Counter helix", "Culling blade"}).
			Items([5]string{"DESOLATOR", "Devine Rapier", "Devine Rapier", "Travel boots", "Devine Rapier"})

	})

}
