package main

import (
	"fmt"

	"example.com/generic/business"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
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
	business.PrintGeneric(windwest) // type inferece

	ss := []business.Solar{solar2k, solar3k}
	business.PrintSlice[business.Solar](ss)
	business.PrintSlice[business.Wind]([]business.Wind{windwest, windwest})

	business.Cost(10, solar2k.Netto)
	business.Cost(10, float64(0.45))
	business.Cost[float64](10, 0.45)

	fmt.Printf("index: %d\n", slices.Index(ss, solar2k))
	business.SortByCost(ss)

	contracts := make(map[int]business.Solar)
	contracts[1] = solar2k
	contracts[2] = solar2k
	contracts[3] = solar3k
	contracts[4] = solar3k

	contractIDs := maps.Keys(contracts)
	fmt.Println(contractIDs)

	ss2 := business.SolarSlice{solar2k, solar3k}
	business.PrintSlice(ss2)
	business.PrintSlice2(ss2)

}
