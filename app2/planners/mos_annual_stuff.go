package planners

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosAnnualHeader struct {
	year   calendar.Year
	layout Layout
}

func newMOSAnnualHeader(year calendar.Year, layout Layout) mosAnnualHeader {
	return mosAnnualHeader{year: year, layout: layout}
}

func (m mosAnnualHeader) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.year)

	header, err := texsnippets.Build(texsnippets.MOSHeader, map[string]string{
		"MarginNotes": `\renewcommand*{\arraystretch}{1.8185}` + texYear.Months() + `\qquad{}` + texYear.Quarters(),
		"Header":      "hello world header",
	})

	if err != nil {
		return nil, fmt.Errorf("build header: %w", err)
	}

	return []string{header}, nil
}

type mosAnnualContents struct {
	year calendar.Year
}

func (m mosAnnualContents) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.year)

	return []string{texYear.BuildCalendar()}, nil
}