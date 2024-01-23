package main

import "fmt"

type house struct {
	name string
}

type animal struct {
	name    string
	candies []int
	house1  *house
}

func main() {
	//list := []int{0, 1, 2, 3}
	//change(list)
	//fmt.Println(list)

	//an1 := &animal{name: "aaa"}
	//an2 := &animal{name: "aaa"}
	//an3 := &animal{name: "aaa"}
	//animals := []*animal{an1, an2, an3}
	//changeAnimal(animals)
	//for i := range animals {
	//	fmt.Println(*animals[i])
	//}

	an := animal{name: "lllaa", candies: []int{0, 1, 2, 3}, house1: &house{name: "111222"}}
	newAn := animal{}
	newAn = an
	fmt.Println(newAn)
	fmt.Println(*((newAn).house1))
	fmt.Printf("%p    %p", &an, &newAn)
}

func change(list []int) {
	list = append(list, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9)
}

func changeAnimal(list []*animal) {
	list[0].name = "888"
}
