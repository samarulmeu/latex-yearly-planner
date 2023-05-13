package mosannualbody

import (
	"bytes"
	"context"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/sections/mosannual"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/tex/parbox"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/tex/texcalendar"
	"text/template"
)

type Body struct {
	templateTree *template.Template

	global              mos.Parameters
	calendarsParameters texcalendar.CalendarLittleParameters
}

func New(global mos.Parameters, calendarsParameters texcalendar.CalendarLittleParameters) Body {
	return Body{
		global:              global,
		calendarsParameters: calendarsParameters,
	}
}

func (r Body) GenerateComponent(
	_ context.Context, pageNumber mosannual.PageNumber, sectionParameters mosannual.SectionParameters,
) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)

	littleCalendars := texcalendar.NewCalendarsLittle(r.global.Months, r.calendarsParameters)

	from := (pageNumber - 1) * sectionParameters.MonthsPerPage
	to := pageNumber * sectionParameters.MonthsPerPage
	columnIndex := 0

	if to > len(littleCalendars) {
		to = len(littleCalendars)
	}

	for i, littleCal := range littleCalendars[from:to] {
		columnIndex = (columnIndex + 1) % sectionParameters.Columns
		buffer.WriteString(parbox.New(sectionParameters.ColumnWidth).SetContent(littleCal).Render())

		if columnIndex == 0 {
			buffer.WriteString("\n\n")

			if i+1 != sectionParameters.MonthsPerPage {
				if sectionParameters.VerticalSpacing.IsSpecial() {
					buffer.WriteString(sectionParameters.VerticalSpacing.String())
				} else {
					buffer.WriteString(`\vspace{` + sectionParameters.VerticalSpacing.String() + `}`)
				}
			}
		} else {
			if sectionParameters.ColumnSpacing.IsSpecial() {
				buffer.WriteString(sectionParameters.ColumnSpacing.String())
			} else {
				buffer.WriteString(`\hspace{` + sectionParameters.ColumnSpacing.String() + `}`)
			}
		}
	}

	return buffer.Bytes(), nil
}
