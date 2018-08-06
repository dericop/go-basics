package main

import "fmt"

func main(){
	i, j := 42, 2701
	p := &i // Referencia en memoria
	fmt.Println(*p) //Acceder al contenido de la referencia de memoria
	*p = 21
	fmt.Println(i) 

	p = &j
	*p = *p / 37
	fmt.Println(j)
}