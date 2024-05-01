package primitives

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Calendar struct {
	*tview.Table
	// options       []string
	date time.Time
}

// NewCalendar takes either 0 or 1 input. If there are no inputs, it takes time.Now(),
// and if there is 1 or above, it uses the first argument and discards the rest.
func NewCalendar(dates ...time.Time) *Calendar {
	var today time.Time

	// Check if a date was provided
	if len(dates) > 0 {
		today = dates[0]
	} else {
		// Set a default value if no date is provided
		today = time.Now()
	}

	if today.Year() < 1970 || today.Month() > 12 {
		panic("unsupported date used in calendar")
	}
	fmt.Println("1")
	// Create a new table
	table := tview.NewTable()
	table.SetBorder(true).SetTitle(fmt.Sprintf("%s %d\n", today.Month(), today.Year()))
	fmt.Println("2")

	// Create a new calendar
	calendar := &Calendar{
		Table: table,
		date:  today,
	}
	fmt.Println("3")

	return calendar
}

func (c *Calendar) Draw(screen tcell.Screen) {
	c.Box.DrawForSubclass(screen, c)

	// Clear the table before redrawing
	c.Clear()

	// Get the dimensions of the table
	//_, _, width, height := c.GetInnerRect()

	// Draw the content of the calendar
	// For example, you can draw the days of the month in each cell

	// Print the current month and year as the title
	year, month, _ := c.date.Date()
	title := fmt.Sprintf("%s %d", month.String(), year)
	c.SetTitle(title)

	// Generate the slice containing all the days of the month
	daysOfMonth := GenerateMonth(year, month)

	// Iterate over the days and set the content of each cell
	for i, day := range daysOfMonth {
		// Calculate the row and column for the current day
		row := i / 7 // Each row represents a week
		col := i % 7 // Each column represents a day of the week

		// Set the content of the cell
		cell := tview.NewTableCell(day.Format("2"))
		c.SetCell(row, col, cell)
	}
}

func GenerateMonth(inYear int, inMonth time.Month) []time.Time {
	res := []time.Time{}
	for i := 1; i <= 31; i++ {
		tempDate := time.Date(inYear, inMonth, i, 0, 0, 0, 0, time.Now().Location())
		if tempDate.Month() > inMonth {
			break
		}
		res = append(res, tempDate)
	}

	return res
}
