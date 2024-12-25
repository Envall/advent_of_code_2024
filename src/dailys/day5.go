package dailys

import (
	"fmt"
)

/*
plan
- datastruktur för rules (tex 47|53)
- iterera över alla 'updates' (tex 75,47,61,53,29)
- applicera alla regler på varje update
	- iterera över regler och verifiera om update ok
	- strunta i regel om den innehåller siffra som ej finns i update..
	- om fail -> fail
- om !fail. Spara mittensidan (u[len(u)/2+1]) (a/b ska vara floor om båda är int)
- summera sparade mittensidor..
*/

func Day5_1(path string) int {
	fmt.Println("Hello")

	return -1
}
