package main

import (
	"log"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func loremList() []string {
	return []string{
		"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod " +
			"tempor incididunt ut labore et dolore magna aliqua.",
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
			"aliquip ex ea commodo consequat.",
		"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum " +
			"dolore eu fugiat nulla pariatur.",
		"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui " +
			"officia deserunt mollit anim id est laborum.",
	}
}

func lorem() string {
	return strings.Join(loremList(), " ")
}

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetFontLocation("fonts/")
	pdf.AddFont("OpenSans-Bold", "", "OpenSans-Bold.json")
	pdf.AddFont("Roboto-Regular", "", "Roboto-Regular.json")

	pdf.SetFont("OpenSans-Bold", "", 16)
	pdf.SetTextColor(25, 25, 25)

	pdf.AddPage()
	// pdf.SetFont("Helvetica", "B", 16)
	pdf.Cell(40, 10, "Malcolm Lockyer")
	pdf.Ln(-1)

	pdf.SetFont("Roboto-Regular", "", 14)
	pdf.Cell(40, 10, "Stuff")

	pdf.Ln(-1)
	pdf.MultiCell(0, 5, lorem(), "", "", false)
	pdf.Ln(-1)
	// pdf.SetFont("", "I", 0)
	pdf.Cell(0, 5, "(end of excerpt)")

	err := pdf.OutputFileAndClose("MalcolmLockyer.pdf")

	if err != nil {
		log.Panic(err)
	}
}
