package redovalnica

import f "fmt"

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

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
