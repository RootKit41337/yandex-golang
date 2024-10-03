package main

import (
	"fmt"
)


type Animal interface{
	MakeSound()
}

type Dog struct{
	
}

type Cat struct{
	
}

func (d Dog) MakeSound() {
	sound_dog := "Гав!"
	fmt.Printf("%s\n", sound_dog)
}

func (c Cat) MakeSound() {
    sound_cat := "Мяу!"
    fmt.Printf("%s\n", sound_cat)
}