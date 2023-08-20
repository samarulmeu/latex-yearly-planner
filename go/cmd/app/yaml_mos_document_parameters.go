package app

import (
	"fmt"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/mosdocument"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/sections/mosannual"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/sections/mosdaily"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/sections/mosdailynotes"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/sections/mosdailyreflections"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/sections/mosdedicatednotes"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/sections/mosmonthly"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/sections/mosquarterly"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/sections/mosweekly"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/tex/texcalendar"
	"github.com/kudrykv/latex-yearly-planner/internal/core/calendar"
	"github.com/kudrykv/latex-yearly-planner/internal/core/entities"
	"gopkg.in/yaml.v3"
	"regexp"
	"strconv"
	"time"
)

type YAMLMOS struct {
	Document   YAMLDocumentParameters `yaml:"document"`
	Parameters YAMLParameters         `yaml:"parameters"`
	Sections   YAMLSections           `yaml:"sections"`
}

type YAMLSections struct {
	AnnualSection           YAMLAnnualSection           `yaml:"yearly"`
	QuarterlySection        YAMLQuarterlySection        `yaml:"quarterly"`
	MonthlySection          YAMLMonthlySection          `yaml:"monthly"`
	WeeklySection           YAMLWeeklySection           `yaml:"weekly"`
	DailySection            YAMLDailySection            `yaml:"daily"`
	DailyReflectionsSection YAMLDailyReflectionsSection `yaml:"daily_reflections"`
	DailyNotesSection       YAMLDailyNotesSection       `yaml:"daily_notes"`
	DedicatedNotesSection   YAMLDedicatedNotesSection   `yaml:"dedicated_notes"`
}

type YAMLAnnualSection struct {
	Enabled bool `yaml:"enabled"`

	Pages                    int             `yaml:"pages"`
	MonthsPerPage            int             `yaml:"months_per_page"`
	Columns                  int             `yaml:"columns"`
	ColumnWidth              entities.Length `yaml:"column_width"`
	ColumnSpacing            entities.Length `yaml:"column_spacing"`
	VerticalSpacing          entities.Length `yaml:"vertical_spacing"`
	UnderfullVerticalSpacing entities.Length `yaml:"underfull_vertical_spacing"`
}

type YAMLQuarterlySection struct {
	Enabled bool `yaml:"enabled"`

	CalendarsColumn          entities.Placement `yaml:"calendars_column"`
	CalendarsColumnWidth     entities.Length    `yaml:"calendars_column_width"`
	CalendarsVerticalSpacing entities.Length    `yaml:"calendars_vertical_spacing"`
	NotesColumnWidth         entities.Length    `yaml:"notes_column_width"`
	ColumnHeight             entities.Length    `yaml:"column_height"`
	ColumnSpacing            entities.Length    `yaml:"column_spacing"`
}

type YAMLMonthlySection struct {
	Enabled bool `yaml:"enabled"`

	NotesWidth     entities.Length `yaml:"notes_width"`
	NotesHeight    entities.Length `yaml:"notes_height"`
	Gap            entities.Length `yaml:"gap"`
	DateLeftOffset entities.Length `yaml:"date_left_offset"`
}

type YAMLWeeklySection struct {
	Enabled bool `yaml:"enabled"`
}

type YAMLDailySection struct {
	Enabled bool `yaml:"enabled"`
}

type YAMLDailyReflectionsSection struct {
	Enabled bool `yaml:"enabled"`
}

type YAMLDailyNotesSection struct {
	Enabled bool `yaml:"enabled"`
}

type YAMLDedicatedNotesSection struct {
	Enabled bool `yaml:"enabled"`

	IndexPages        int `yaml:"index_pages"`
	NotesPerIndexPage int `yaml:"notes_per_index_page"`
	NotesNumber       int `yaml:"notes_number"`
	PagesPerNote      int `yaml:"pages_per_note"`
}

type YAMLDocumentParameters struct {
	Layout YAMLLayout

	DebugOptions YAMLDebugOptions
}

func (r YAMLMOS) ToDocumentParameters() mosdocument.DocumentParameters {
	return mosdocument.DocumentParameters{
		Layout:       r.Document.Layout.ToLayout(),
		DebugOptions: r.Document.DebugOptions.ToDebugOptions(),
	}
}

func (r YAMLMOS) ToParameters() mos.Parameters {
	months := calendar.NewMonths(
		calendar.Weekday(r.Parameters.WeekdayStart),
		r.Parameters.StartDate.Year(), r.Parameters.StartDate.Month(),
		r.Parameters.EndDate.Year(), r.Parameters.EndDate.Month(),
	)

	return mos.Parameters{
		Months: months,
	}
}

func (r YAMLMOS) AnnualParameters() mosannual.SectionParameters {
	return mosannual.SectionParameters{
		Enabled: r.Sections.AnnualSection.Enabled,

		Pages:                    r.Sections.AnnualSection.Pages,
		MonthsPerPage:            r.Sections.AnnualSection.MonthsPerPage,
		Columns:                  r.Sections.AnnualSection.Columns,
		ColumnWidth:              r.Sections.AnnualSection.ColumnWidth,
		ColumnSpacing:            r.Sections.AnnualSection.ColumnSpacing,
		VerticalSpacing:          r.Sections.AnnualSection.VerticalSpacing,
		UnderfullVerticalSpacing: r.Sections.AnnualSection.UnderfullVerticalSpacing,
	}
}

func (r YAMLMOS) LittleCalendarParameters() texcalendar.CalendarLittleParameters {
	return texcalendar.CalendarLittleParameters{
		ShowWeekNumbers:     r.Parameters.ShowWeekNumbers,
		WeekNumberPlacement: r.Parameters.WeekNumberPlacement,
	}
}

func (r YAMLMOS) QuarterlyParameters() mosquarterly.SectionParameters {
	return mosquarterly.SectionParameters{
		Enabled: r.Sections.QuarterlySection.Enabled,

		CalendarsColumn:          r.Sections.QuarterlySection.CalendarsColumn,
		CalendarsColumnWidth:     r.Sections.QuarterlySection.CalendarsColumnWidth,
		NotesColumnWidth:         r.Sections.QuarterlySection.NotesColumnWidth,
		ColumnHeight:             r.Sections.QuarterlySection.ColumnHeight,
		CalendarsVerticalSpacing: r.Sections.QuarterlySection.CalendarsVerticalSpacing,
		ColumnSpacing:            r.Sections.QuarterlySection.ColumnSpacing,
	}
}

func (r YAMLMOS) MonthlyParameters() mosmonthly.SectionParameters {
	return mosmonthly.SectionParameters{
		Enabled: r.Sections.MonthlySection.Enabled,

		Gap:         r.Sections.MonthlySection.Gap,
		NotesWidth:  r.Sections.MonthlySection.NotesWidth,
		NotesHeight: r.Sections.MonthlySection.NotesHeight,
	}
}

func (r YAMLMOS) LargeCalendarParameters() texcalendar.CalendarLargeParameters {
	return texcalendar.CalendarLargeParameters{
		ShowWeekNumbers:     r.Parameters.ShowWeekNumbers,
		WeekNumberPlacement: r.Parameters.WeekNumberPlacement,
		DateLeftOffset:      r.Sections.MonthlySection.DateLeftOffset,
	}
}

func (r YAMLMOS) WeeklyParameters() mosweekly.SectionParameters {
	return mosweekly.SectionParameters{
		Enabled: r.Sections.WeeklySection.Enabled,
	}
}

func (r YAMLMOS) DailyParameters() mosdaily.SectionParameters {
	return mosdaily.SectionParameters{
		Enabled: r.Sections.DailySection.Enabled,
	}
}

func (r YAMLMOS) DailyReflectionsParameters() mosdailyreflections.SectionParameters {
	return mosdailyreflections.SectionParameters{
		Enabled: r.Sections.DailyReflectionsSection.Enabled,
	}
}

func (r YAMLMOS) DailyNotesParameters() mosdailynotes.SectionParameters {
	return mosdailynotes.SectionParameters{
		Enabled: r.Sections.DailyNotesSection.Enabled,
	}
}

func (r YAMLMOS) DedicatedNotesParameters() mosdedicatednotes.SectionParameters {
	return mosdedicatednotes.SectionParameters{
		Enabled: r.Sections.DedicatedNotesSection.Enabled,

		IndexPages:        r.Sections.DedicatedNotesSection.IndexPages,
		NotesPerIndexPage: r.Sections.DedicatedNotesSection.NotesPerIndexPage,
		NotesNumber:       r.Sections.DedicatedNotesSection.NotesNumber,
		PagesPerNote:      r.Sections.DedicatedNotesSection.PagesPerNote,
	}
}

type YAMLDimensions struct {
	Width  entities.Length `yaml:"width"`
	Height entities.Length `yaml:"height"`
}

func (r YAMLDimensions) ToDimensions() mosdocument.Dimensions {
	return mosdocument.Dimensions{
		Width:  r.Width,
		Height: r.Height,
	}
}

type YAMLLayout struct {
	Dimensions  YAMLDimensions  `yaml:"dimensions"`
	Margin      YAMLMargin      `yaml:"margin"`
	MarginNotes YAMLMarginNotes `yaml:"margin_notes"`
}

func (r YAMLLayout) ToLayout() mosdocument.Layout {
	return mosdocument.Layout{
		Dimensions:  r.Dimensions.ToDimensions(),
		Margin:      r.Margin.ToMargin(),
		MarginNotes: r.MarginNotes.ToMarginNotes(),
	}
}

type YAMLMargin struct {
	Top    entities.Length `yaml:"top"`
	Right  entities.Length `yaml:"right"`
	Bottom entities.Length `yaml:"bottom"`
	Left   entities.Length `yaml:"left"`
}

func (r YAMLMargin) ToMargin() mosdocument.Margin {
	return mosdocument.Margin{
		Top:    r.Top,
		Right:  r.Right,
		Bottom: r.Bottom,
		Left:   r.Left,
	}
}

type YAMLMarginNotes struct {
	Width     entities.Length `yaml:"width"`
	Separator entities.Length `yaml:"separator"`
	Placement string          `yaml:"placement"`
}

func (r YAMLMarginNotes) ToMarginNotes() mosdocument.MarginNotes {
	return mosdocument.MarginNotes{
		Width:     r.Width,
		Separator: r.Separator,
		Placement: r.Placement,
	}
}

type YAMLDebugOptions struct {
	Enabled bool `yaml:"enabled"`

	HighlightLinks  bool `yaml:"highlight_links"`
	HighlightFrames bool `yaml:"highlight_frames"`
}

func (r YAMLDebugOptions) ToDebugOptions() mosdocument.DebugOptions {
	return mosdocument.DebugOptions{
		Enabled:         r.Enabled,
		HighlightLinks:  r.HighlightLinks,
		HighlightFrames: r.HighlightFrames,
	}
}

type YAMLParameters struct {
	StartDate           YAMLDate           `yaml:"start_date"`
	EndDate             YAMLDate           `yaml:"end_date"`
	WeekdayStart        YAMLWeekday        `yaml:"weekday_start"`
	ShowWeekNumbers     bool               `yaml:"show_week_numbers"`
	WeekNumberPlacement entities.Placement `yaml:"week_number_placement"`
	DateLeftOffset      entities.Length    `yaml:"date_left_offset"`
	DateTopOffset       entities.Length    `yaml:"date_top_offset"`
	FormatAMPM          bool               `yaml:"format_ampm"`
	Notes               YAMLNotes          `yaml:"notes"`
}

type YAMLNotes struct {
	Pattern string `yaml:"pattern"`
}

var dateRegex = regexp.MustCompile(`(?i)(^\d{4}),?\s* (january|february|march|april|may|june|july|august|september|november|october|december)$`)

type YAMLDate time.Time

func (r *YAMLDate) UnmarshalYAML(value *yaml.Node) error {
	if value.Tag != "!!str" {
		return entities.ErrNotAString
	}

	if !dateRegex.MatchString(value.Value) {
		return fmt.Errorf("expected '<year>, <month_name>':%w", entities.ErrBadPattern)
	}

	matches := dateRegex.FindAllStringSubmatch(value.Value, -1)[0]
	stringYear, stringMonth := matches[1], matches[2]

	year, err := strconv.Atoi(stringYear)
	if err != nil {
		return fmt.Errorf("parse year: %w", err)
	}

	month, err := parseMonth(stringMonth)
	if err != nil {
		return fmt.Errorf("parse month: %w", err)
	}

	*r = YAMLDate(time.Date(year, month, 1, 0, 0, 0, 0, time.Local))

	return nil
}

func (r *YAMLDate) Year() int {
	return time.Time(*r).Year()
}

func (r *YAMLDate) Month() time.Month {
	return time.Time(*r).Month()
}

func parseMonth(month string) (time.Month, error) {
	switch {
	case regexp.MustCompile("(?i)january").MatchString(month):
		return time.January, nil
	case regexp.MustCompile("(?i)february").MatchString(month):
		return time.February, nil
	case regexp.MustCompile("(?i)march").MatchString(month):
		return time.March, nil
	case regexp.MustCompile("(?i)april").MatchString(month):
		return time.April, nil
	case regexp.MustCompile("(?i)may").MatchString(month):
		return time.May, nil
	case regexp.MustCompile("(?i)june").MatchString(month):
		return time.June, nil
	case regexp.MustCompile("(?i)july").MatchString(month):
		return time.July, nil
	case regexp.MustCompile("(?i)august").MatchString(month):
		return time.August, nil
	case regexp.MustCompile("(?i)september").MatchString(month):
		return time.September, nil
	case regexp.MustCompile("(?i)october").MatchString(month):
		return time.October, nil
	case regexp.MustCompile("(?i)november").MatchString(month):
		return time.November, nil
	case regexp.MustCompile("(?i)december").MatchString(month):
		return time.December, nil
	default:
		return 0, fmt.Errorf("unknown month: %s", month)
	}
}

var weekdayRegex = regexp.MustCompile(`(?i)(^monday|tuesday|wednesday|thursday|friday|saturday|sunday)$`)

type YAMLWeekday calendar.Weekday

func (r *YAMLWeekday) UnmarshalYAML(value *yaml.Node) error {
	if value.Tag != "!!str" {
		return entities.ErrNotAString
	}

	if !weekdayRegex.MatchString(value.Value) {
		return fmt.Errorf("expected '<weekday>':%w", entities.ErrBadPattern)
	}

	weekday, err := parseWeekday(value.Value)
	if err != nil {
		return fmt.Errorf("parse weekday: %w", err)
	}

	*r = YAMLWeekday(weekday)

	return nil
}

func parseWeekday(value string) (calendar.Weekday, error) {
	switch {
	case regexp.MustCompile("(?i)monday").MatchString(value):
		return calendar.Monday, nil
	case regexp.MustCompile("(?i)tuesday").MatchString(value):
		return calendar.Tuesday, nil
	case regexp.MustCompile("(?i)wednesday").MatchString(value):
		return calendar.Wednesday, nil
	case regexp.MustCompile("(?i)thursday").MatchString(value):
		return calendar.Thursday, nil
	case regexp.MustCompile("(?i)friday").MatchString(value):
		return calendar.Friday, nil
	case regexp.MustCompile("(?i)saturday").MatchString(value):
		return calendar.Saturday, nil
	case regexp.MustCompile("(?i)sunday").MatchString(value):
		return calendar.Sunday, nil
	default:
		return calendar.Weekday{}, fmt.Errorf("unknown weekday: %s", value)
	}
}