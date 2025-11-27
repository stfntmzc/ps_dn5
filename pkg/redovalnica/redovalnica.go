package redovalnica

import f "fmt"

var StOcen int = 3
var MinOcena int = 3
var MaxOcena int = 3

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	_, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		f.Printf("Študenta z vpisno številko %s ni na seznamu.\n", vpisnaStevilka)
		return
	}
	if ocena < MinOcena || MaxOcena < ocena {
		f.Printf("Ocena mora biti število od %d do %d.\n", MinOcena, MaxOcena)
		return
	}

	student := studenti[vpisnaStevilka]
	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisnaStevilka] = student
}

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

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisnaStevilka, student := range studenti {
		avg := povprecje(studenti, vpisnaStevilka)
		f.Printf("%s %s: povprečna ocena %.2f -> ", student.Ime, student.Priimek, povprecje(studenti, vpisnaStevilka))
		if len(student.Ocene) < StOcen {
			f.Println("Premalo ocen za zaključek ocene")
		} else {
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
}

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
