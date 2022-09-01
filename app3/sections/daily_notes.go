package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type DailyNotesParameters struct {
	Pages           int
	NotesParameters components.NotesParameters
}

type DailyNotes struct {
	parameters DailyNotesParameters
	day        calendar.Day
}

func NewDailyNotes(day calendar.Day, parameters DailyNotesParameters) (DailyNotes, error) {
	return DailyNotes{
		day:        day,
		parameters: parameters,
	}, nil
}

func (r DailyNotes) Build() ([]string, error) {
	pages := make([]string, 0, r.parameters.Pages)

	for i := 0; i < r.parameters.Pages; i++ {
		pages = append(pages, r.day.Format("2006-01-02"))
	}

	return pages, nil
}

func (r DailyNotes) Link(text string) string {
	return fmt.Sprintf(`\hyperlink{daily-notes-%s}{%s}`, r.day.Format("2006-01-02"), text)
}

func (r DailyNotes) Reference() string {
	return fmt.Sprintf(`daily-notes-%s`, r.day.Format("2006-01-02"))
}

func (r DailyNotes) Repeat() int {
	return r.parameters.Pages
}
