package eventmanager

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Student struct {
	ID                     string         `bson:"_id"`
	ClassCode              string         `bson:"classCode"`
	AttendancePeriodSetMap map[string]int `bson:"attendancePeriodSetMap"`
}

func Export(students []Student, dates []AcamedicDate, classCode string) string {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	f.SetCellValue("Sheet1", "A1", "roll_no")
	f.SetCellValue("Sheet1", "B1", "class")

	periods := [][]string{{"C", "one"},
		{"D", "two"},
		{"E", "three"},
		{"F", "four"},
		{"G", "five"},
		{"H", "six"},
		{"I", "seven"},
		{"J", "total_percent"},
	}
	for _, period := range periods {
		f.SetCellValue("Sheet1", period[0]+strconv.Itoa(1), period[1])
	}

	working_days := len(dates)

	for ind, student := range students {
		pos := strconv.Itoa(ind + 2)
		f.SetCellValue("Sheet1", "A"+pos, student.ID)
		f.SetCellValue("Sheet1", "B"+pos, student.ClassCode)

		var sum int
		var percentage float32
		final_percentage := 0.0
		final_sum := 0
		for i := 0; i < 7; i++ {
      sum = 0
			for _, date := range dates {
				value, exists := student.AttendancePeriodSetMap[strconv.Itoa(date.ID)]
				values := numberToBoolArray(value)
				if exists {
					if values[i] {
						sum++
					}
				}
				percentage = float32(sum) / float32(working_days) * 100
				f.SetCellValue("Sheet1", periods[i][0]+pos, fmt.Sprintf("%.2f", percentage))
			}
		}
		for _, date := range dates {
			_, exists := student.AttendancePeriodSetMap[strconv.Itoa(date.ID)]
			if exists {
				final_sum++
			}
		}
		final_percentage = float64(final_sum) / float64(working_days) * 100
		f.SetCellValue("Sheet1", "J"+pos, fmt.Sprintf("%.2f", final_percentage))
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(1)

	// Save spreadsheet by the given path.
	fileName := classCode + "_sheet.xlsx"
	if err := f.SaveAs(classCode + "_sheet.xlsx"); err != nil {
		fmt.Println(err)
	}
	return fileName
}

func numberToBoolArray(num int) [7]bool {
	// Define an array to store the boolean values
	var boolArray [7]bool

	// Iterate over each bit position from right to left
	for i := 6; i >= 0; i-- {
		// Check if the bit at the current position is set (1) or unset (0)
		boolArray[i] = (num>>uint(i))&1 == 1
	}

	return boolArray
}
