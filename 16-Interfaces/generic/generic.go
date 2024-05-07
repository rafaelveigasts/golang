package main

import (
	ftm "fmt"
)

func generica(interfaceGenerica interface{}) {
	ftm.Println(interfaceGenerica) 
} // interface vazia, aceita qualquer tipo

func main(){
	generica("String")
	generica(1)
	generica(true)
	generica(1.5)
}
