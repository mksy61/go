package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	bs := `Neptün'ün halkaları, Neptün etrafında yer alan ve beş ana halkadan oluşan bir halka sistemidir. 
	Başta "yaylar" olarak adlandırılan halkalar, 22 Temmuz 1984'te Patrice Bouchet, Reinhold Häfner ve Jean Manfroid'dan oluşan ekip tarafından 
	Şili'deki La Silla Gözlemevi'nde ve William Hubbard liderliğindeki bir program kapsamında 
	F. Vilas ve L. R. Elicer tarafından Cerro Tololo Amerikaarası Gözlemevi'nde keşfedildi. `

	for _, r := range bs {
		fmt.Printf("%c", r)
	}
	fmt.Println()
	fmt.Println(utf8.RuneCountInString(bs))
}
