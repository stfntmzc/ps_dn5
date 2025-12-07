// Package redovalnica omogoča upravljanje ocen študentov.
// Vsebuje funkcije za dodajanje ocen, izpis vseh ocen in izračun končnega uspeha.
// Funkcije, ki se začnejo z veliko začetnico, so izvozene in jih lahko uporablja paket main.
// Funkcija povprecje ostaja skrita znotraj paketa.
package redovalnica

import f "fmt"

// Student predstavlja podatke o študentu, vključno z imenom, priimkom in seznamom ocen.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// DodajOceno doda oceno študentu, ki ga identificira vpisna številka.
// Če ocena ni v interval [1,10] (ali pa podan interval) ali študent ne obstaja, se izpiše opozorilo.
// Ocena se doda neposredno v seznam ocen študenta.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena < 1 || 10 < ocena {
		f.Println("Ocena mora biti število od 1 do 10.")
		return
	}
	_, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		f.Printf("Študenta z vpisno številko %s ni na seznamu.\n", vpisnaStevilka)
		return
	}

	student := studenti[vpisnaStevilka]
	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisnaStevilka] = student
}

// IzpisVsehOcen izpiše vse ocene vseh študentov v obliki:
// Vpisna številka - Ime Priimek: [ocena1 ocena2 ...]
func IzpisVsehOcen(studenti map[string]Student) {
	f.Println("VSE OCENE:")
	for vpisnaStevilka, student := range studenti {
		f.Printf("%s - %s %s: [", vpisnaStevilka, student.Ime, student.Priimek)
		for i, ocena := range student.Ocene {
			f.Printf("%d", ocena)
			if i < len(student.Ocene)-1 {
				f.Printf(" ")
			}
		}
		f.Printf("]\n")
	}
}

// IzpisiKoncniUspeh izračuna povprečne ocene vseh študentov in jih kategorizira:
// Odličen študent (povprečje >= 9), Povprečen študent (>=6), Neuspešen študent (<6)
func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisnaStevilka, student := range studenti {
		avg := povprecje(studenti, vpisnaStevilka)
		f.Printf("%s %s: povprečna ocena %.2f -> ", student.Ime, student.Priimek, povprecje(studenti, vpisnaStevilka))
		switch {
		case avg >= 9.0:
			f.Println("Odličen študent!")
		case avg >= 6.0:
			f.Println("Povprečen študent")
		default:
			f.Println("Neuspešen študent")
		}
	}
}

// povprecje izračuna povprečno oceno študenta. Funkcija je skrita (ne izvozi se).
// Če študent ne obstaja, vrne -1.0.
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	_, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		return -1.0
	}

	var avg float64 = 0.0
	for _, ocena := range studenti[vpisnaStevilka].Ocene {
		avg += float64(ocena)
	}

	return (avg / float64(len(studenti[vpisnaStevilka].Ocene)))
}
