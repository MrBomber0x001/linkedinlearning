package main

import (
	"fmt"

	"example.com/generic/business"
)

func main() {
	solar2k := business.Solar{"Solar 2000", 3.500}
	solar3k := business.Solar{"Solar 3000", 3.00}
	windwest := business.Wind{"Wind West", 3.500}

	fmt.Println(solar2k.Print())
	fmt.Println(solar3k.Print())
	fmt.Println(windwest.Print())

	fmt.Printf("First generic functions, - %s", business.PrintGeneric[business.Solar](solar2k))
	business.PrintGeneric[business.Wind](windwest)

	ss := []business.Solar{solar2k, solar3k}
	business.PrintSlice[business.Solar](ss)
	business.PrintSlice[business.Wind]([]business.Wind{windwest, windwest})

}
