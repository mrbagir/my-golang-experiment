package main

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	const (
		fontPtSize = 18.0
		wd         = 100.0
	)
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	pdf.SetFont("Times", "", fontPtSize)
	_, lineHt := pdf.GetFontSize()
	pdf.AddPage()
	pdf.SetMargins(10, 10, 10)
	lines := pdf.SplitLines([]byte(lorem()), wd)
	ht := float64(len(lines)) * lineHt
	y := (297.0 - ht) / 2.0
	pdf.SetDrawColor(128, 128, 128)
	pdf.SetFillColor(255, 255, 210)
	x := (210.0 - (wd)) / 2.0
	pdf.Rect(x, y, wd, ht, "FD")
	pdf.SetY(y)
	for _, line := range lines {
		// pdf.CellFormat(190.0, lineHt, string(line), "", 1, "C", false, 0, "")
		pdf.MultiCell(190.0, lineHt, string(line), "1", "C", false)
	}

	pdf.OutputFileAndClose("output.pdf")
}

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
		"officia deserunt mollit anim id est laborum.",
	}
}

func lorem() string {
	return strings.Join(loremList(), " ")
}
