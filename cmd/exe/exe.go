package main

import (
	"context"
	f "fmt"
	"log"
	"os"

	"github.com/stfntmzc/ps_dn5/pkg/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Ocene študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: 3,
				Usage: "najmanjše število ocen za pozitivno oceno",
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Value: 1,
				Usage: "najmanjša možna ocena",
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Value: 10,
				Usage: "največja možna ocena",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {

			redovalnica.StOcen = c.Int("stOcen")
			redovalnica.MinOcena = c.Int("minOcena")
			redovalnica.MaxOcena = c.Int("maxOcena")

			studenti := map[string]redovalnica.Student{
				"01010101": {Ime: "Adam", Priimek: "Bohorič", Ocene: []int{10, 10}},
				"12121212": {Ime: "Bine", Priimek: "Cvetković", Ocene: []int{8, 9, 9}},
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
			f.Println("Dodajanje ocene 8 študentu, ki ni na seznamu (vpisna številka 56565656):")
			redovalnica.DodajOceno(studenti, "56565656", 8)
			f.Println("")
			f.Printf("Dodajanje ocene, ki je izven intervala [%d,%d]:\n", c.Int("minOcena"), c.Int("maxOcena"))
			redovalnica.DodajOceno(studenti, "23232323", 23)
			f.Println("")

			f.Println("IzpisVsehOcen ===========================")
			f.Println("")
			redovalnica.IzpisVsehOcen(studenti)
			f.Println("")

			f.Println("IzpisiKoncniUspeh ===========================")
			f.Println("")
			redovalnica.IzpisiKoncniUspeh(studenti)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
