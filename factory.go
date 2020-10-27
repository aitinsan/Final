package main

import "fmt"

type Creep struct{
	NameArmy,Group string
	People int
	OnWar bool

}
func NewCreepsFactory(name string, ready bool) func(group string,people int) *Creep{
	return func(group string,people int)*Creep{
		return &Creep{name,group,people,ready}
	}
}
func main(){

	direCreepsFactory:= NewCreepsFactory("Dire",true)
	direCreepsBlyshniki:=direCreepsFactory("ближний",3)
	direCreepsDalniki:=direCreepsFactory("дальний",1)
	fmt.Println(direCreepsBlyshniki)
	fmt.Println(direCreepsDalniki)

	radiantCreepsFactory:= NewCreepsFactory("Radiant",true)
	radiantCreepsBlyshniki:=radiantCreepsFactory("ближний",3)
	radiantCreepsDalniki:=radiantCreepsFactory("дальний",1)
	fmt.Println(radiantCreepsBlyshniki)
	fmt.Println(radiantCreepsDalniki)

	neutralCreepsFactory1:= NewCreepsFactory("Волки",true)
	neutralCreepsBlyshniki:=neutralCreepsFactory1("ближний",3)
	fmt.Println(neutralCreepsBlyshniki)



}