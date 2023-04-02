package business

import "fmt"

type Solar struct {
	Name  string
	Netto float64
}

type Wind struct {
	Name  string
	Netto float64
}

type Energy interface {
	Solar | Wind
	Cost() float64
}

func (s *Solar) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *s)
}

func (w *Wind) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *w)
}

// I've changed the type from any to Energy
func PrintGeneric[T Energy](t T) string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, t)
}

// TASK: Print Slice using Generics
//PrintSlice: accept any slice of type T and print it's elements using PrintGeneric func!
func PrintSlice[T Energy](tt []T) {
	for i, t := range tt {
		fmt.Printf("%d: %s\n", i, PrintGeneric[T](t))
	}
}

var kinetecoPrint string = "Kineteco Deal:"
