package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("test2.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	start := time.Now()
	fmt.Println("Time Start:", start.Format("15:04:05.000"))
	defer func() {
		end := time.Now()
		fmt.Println("Time End:", end.Format("15:04:05.000"))
		fmt.Println("Duration:", end.Sub(start))
	}()

	streamRows, _ := f.Rows("Sheet1")
	
	for streamRows.Next() {
		cells, err := streamRows.Columns()
		if err != nil {
			fmt.Println(err)
			return
		}

		for i, _ := range cells {
			fmt.Printf("%s%d", string(rune(65+i)), rowIndex)
		}
		fmt.Println()
		rowIndex++
	}

	// rows, err := f.GetRows("Sheet1", excelize.Options{RawCellValue: false, })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// for i, row := range rows {
	// 	for j, _ := range row {
	// 		fmt.Printf("%s%d", string(65+i), j+1)
	// 	}
	// 	fmt.Printf("\n")
	// }

	return

	// pdf := gofpdf.New("P", "mm", "A4", "")
	// pdf.AddPage()
	// pdf.SetFont("Arial", "", 10)

	// lineWidth := 65
	// pageWidth, _ := pdf.GetPageSize()
	// marginLeft, _, _, _ := pdf.GetMargins()
	// maxWidth := pageWidth - 2*marginLeft

	// for i, row := range rows {
	// 	if i == 0 {
	// 		continue
	// 	}

	// 	for j, cell := range row {
	// 		if j >= len(rows[0]) || rows[0][j] == "" {
	// 			continue
	// 			// break
	// 		}

	// 		label := strings.ToUpper(rows[0][j])
	// 		text := fmt.Sprintf("%s: %s", label, cell)

	// 		wrapped := wrapText(text, lineWidth)

	// 		if j == 0 {
	// 			numbered := fmt.Sprintf("%d. %s", i, wrapped)
	// 			pdf.SetX(marginLeft)
	// 			pdf.MultiCell(maxWidth, 5, numbered, "", "L", false)
	// 		} else {
	// 			pdf.SetX(marginLeft)
	// 			pdf.MultiCell(maxWidth, 5, wrapped, "", "L", false)
	// 		}
	// 	}
	// }

	// err = pdf.OutputFileAndClose("test.pdf")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func wrapText(s string, maxLen int) string {
	words := strings.Fields(s)
	var lines []string
	var currentLine string

	for _, word := range words {
		for len(word) > maxLen {
			if currentLine != "" {
				lines = append(lines, currentLine)
				currentLine = ""
			}
			lines = append(lines, word[:maxLen])
			word = word[maxLen:]
		}

		if len(currentLine)+len(word)+1 > maxLen {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return strings.Join(lines, "\n")
}
