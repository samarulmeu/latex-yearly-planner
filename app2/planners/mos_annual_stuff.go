package planners

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosAnnualHeader struct {
	year   calendar.Year
	layout Layout

	left  string
	right []string
}

func newMOSAnnualHeader(year calendar.Year, layout Layout, left string, right []string) mosAnnualHeader {
	return mosAnnualHeader{
		year:   year,
		layout: layout,

		left:  left,
		right: right,
	}
}

func (m mosAnnualHeader) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.year)

	header, err := texsnippets.Build(texsnippets.MOSHeader, map[string]string{
		"MarginNotes": `\renewcommand{\arraystretch}{2.042}` + texYear.Months() + `\qquad{}` + texYear.Quarters(),
		"Header": `{\renewcommand{\arraystretch}{1.8185}\begin{tabularx}{\linewidth}{@{}lY|` + strings.Join(strings.Split(strings.Repeat("r", len(m.right)), ""), "|") + `|@{}}
\Huge ` + m.left + `{\Huge\color{white}{Q}} & & Calendar & To Do & Notes \\ \hline
\end{tabularx}}`,
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
