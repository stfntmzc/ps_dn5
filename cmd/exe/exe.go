package main

import (
	f "fmt"

	"github.com/stfntmzc/ps_dn5/pkg/redovalnica"
)

func main() {

	studenti := map[string]redovalnica.Student{
		"01010101": {Ime: "Adam", Priimek: "Bohorič", Ocene: []int{10, 10, 10, 10, 9}},
		"12121212": {Ime: "Bine", Priimek: "Cvetković", Ocene: []int{8, 9, 9, 8, 9}},
		"23232323": {Ime: "Jan", Priimek: "Beržan", Ocene: []int{8, 9, 9, 9, 9}},
		"34343434": {Ime: "Andraž", Priimek: "Novinšek", Ocene: []int{6, 10, 9, 8, 7}},
		"45454545": {Ime: "Janez", Priimek: "Novak", Ocene: []int{6, 6, 5, 5, 5}},
		"67676767": {Ime: "Jurij", Priimek: "Dalmatin", Ocene: []int{9, 9, 9, 9}},
		"78787878": {Ime: "Primož", Priimek: "Trubar", Ocene: []int{6, 6, 6, 6, 6}},
	}

	f.Println("DodajOceno ===========================")
	f.Println("")
	f.Println("Dodajanje ocene Janu Beržanu:")
	f.Println(studenti["23232323"])
	redovalnica.DodajOceno(studenti, "23232323", 8)
	f.Println(studenti["23232323"])
	f.Println("")
	f.Println("Dodajanje ocene študentu, ki ni na seznamu (vpisna številka 56565656):")
	redovalnica.DodajOceno(studenti, "56565656", 8)
	f.Println("")
	f.Println("Dodajanje ocene, ki je izven intervala [1,10]:")
	redovalnica.DodajOceno(studenti, "23232323", 23)
	f.Println("")

	f.Println("IzpisVsehOcen ===========================")
	f.Println("")
	redovalnica.IzpisVsehOcen(studenti)
	f.Println("")

	f.Println("IzpisiKoncniUspeh ===========================")
	f.Println("")
	redovalnica.IzpisiKoncniUspeh(studenti)
}
