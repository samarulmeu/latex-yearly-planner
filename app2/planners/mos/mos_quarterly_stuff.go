package mos

import (
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosQuarterlyContents struct {
	quarter calendar.Quarter
}

func (r mosQuarterlyContents) Build() ([]string, error) {
	return []string{texcalendar.NewQuarter(r.quarter).BuildPage()}, nil
}