package main

import "fmt"

func main() {

	fmt.Println("-------------For condicional-------------")
	// For condicional
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------For while-------------")
	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("-------------For forever-------------")
	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever > 12 {
			break
		}
	}

	fmt.Println("-------------For Range-------------")
	//For Range
	arreglo := [8]int{0, 1, 4, 6, 10, 9}
	fmt.Println("Arreglo:", arreglo)

	fmt.Println("Primer ejemplo")
	for i, j := range arreglo {
		fmt.Printf("indice i: %d tiene como valor #%d\n", i, j)
	}

	fmt.Println("Segundo ejemplo")
	for i := range arreglo {
		fmt.Printf("Valor de i: %d\n", i)
	}

	fmt.Println("Tercer ejemplo")
	for _, j := range arreglo {
		fmt.Printf("Valor de i: %d\n", j)
	}

	fmt.Println("-------------For con funciones-------------")
	// For con funciones
	for i := preFor(); condicion(i); i = postFor(i) {
		fmt.Printf("Valor de i: %d", i)
		if i == 7 {
			fmt.Printf(" así que saldremos del ciclo...\n")
			break /// este ejemplo es para usar el break
		}
		fmt.Printf("\n")
	}

	fmt.Println("-------------For con goto tags-------------")
	var i int
CICLO:
	fmt.Println("estamos fuera del for")
	for i < 10 {
		if i == 6 {
			i = i + 3
			fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
			goto CICLO2
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
CICLO2:
	fmt.Printf("ciclo 2 Valor de i: %d\n", i)
	if i == 9 {
		fmt.Printf("Valor de i: %d\n", i)
		i = i + 3
		fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
		goto CICLO
	}
	fmt.Printf("terminamos\n")

	fmt.Println("-------------Reto for decreciente-------------")
	// Reto for decreciente
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}

}

func preFor() int {
	fmt.Println("prefor i")
	return 0
}
func postFor(i int) int {
	fmt.Println("postFor sumemos i")
	i++
	return i
}
func condicion(i int) bool {
	fmt.Println("condicion i")

	return (i < 10)
}
