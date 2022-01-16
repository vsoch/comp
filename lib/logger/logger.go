package logger

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	white  = "\033[37m"
)

type Logger struct{}

// A Row holds the row content and basic type
type Row struct {
	Content []string
	Type    string
}

// A struct can be instantiated by a logger to print a table
type Table struct {
	Rows         []Row
	Padding      int
	HeaderColor  string
	DividerColor string
	TableColor   string
}

// Print a single string in a color of choice
func (l *Logger) Red(message string) {
	l.printColor(message, red)
}
func (l *Logger) Green(message string) {
	l.printColor(message, green)
}
func (l *Logger) Yellow(message string) {
	l.printColor(message, yellow)
}

func (l *Logger) Blue(message string) {
	l.printColor(message, blue)
}
func (l *Logger) Purple(message string) {
	l.printColor(message, purple)
}
func (l *Logger) White(message string) {
	l.printColor(message, white)
}
func (l *Logger) Cyan(message string) {
	l.printColor(message, cyan)
}

// printColor will print colored messages
func (l *Logger) printColor(message string, color string) {
	fmt.Printf("%s", string(color))
	fmt.Printf(message)
	fmt.Println(string(reset))
}

// formatColor is the same, but returns the string
func (l *Logger) formatColor(message string, color string) string {
	return fmt.Sprintf("%s%s%s", string(color), message, string(reset))
}

func (l *Logger) RedColor(message string) string {
	return l.formatColor(message, red)
}
func (l *Logger) GreenColor(message string) string {
	return l.formatColor(message, green)
}
func (l *Logger) YellowColor(message string) string {
	return l.formatColor(message, yellow)
}

func (l *Logger) BlueColor(message string) string {
	return l.formatColor(message, blue)
}
func (l *Logger) PurpleColor(message string) string {
	return l.formatColor(message, purple)
}
func (l *Logger) WhiteColor(message string) string {
	return l.formatColor(message, white)
}
func (l *Logger) CyanColor(message string) string {
	return l.formatColor(message, cyan)
}

// Table returns the same logger with a writer
// log = logger.Logger{}
// table = log.Table()
func (l *Logger) Table() *Table {
	rows := []Row{}
	return &Table{Rows: rows, Padding: 1, HeaderColor: "cyan", DividerColor: "blue", TableColor: ""}
}

// Functions for the table!

// SetPadding for the table
func (t *Table) SetPadding(padding int) {
	t.Padding = padding
}

// SetHeader sets the header for the table
func (t *Table) SetHeader(fields []string) {
	t.AddHeader(fields)
	t.AddDivider(fields)
}

// AddDivider adds a divider the length of each field
func (t *Table) AddDivider(fields []string) {
	row := Row{Type: "divider"}
	for _, field := range fields {
		divider := strings.Repeat("-", len(field))
		row.Content = append(row.Content, divider)
	}
	t.Rows = append(t.Rows, row)
}

// AddRow adds a row to the list to be printed later
func (t *Table) AddRow(fields []string) {
	row := Row{Type: "row", Content: fields}
	t.Rows = append(t.Rows, row)
}

// AddHeader adds a Header row
func (t *Table) AddHeader(fields []string) {
	row := Row{Type: "header", Content: fields}
	t.Rows = append(t.Rows, row)
}

// formatHeader adds color for the header (and reset string)
func (t *Table) formatHeader(row []string) string {
	color := t.getColor(t.HeaderColor)
	return fmt.Sprintf("%s%s%s", string(color), strings.Join(row, "\t"), string(reset))
}

// formatDivider adds the DividerColor to a row (and reset string)
func (t *Table) formatDivider(row []string) string {
	color := t.getColor(t.DividerColor)
	return string(color) + strings.Join(row, "\t") + string(reset)
}

// formatRow adds the TableColor to a row (and reset string)
func (t *Table) formatRow(row []string) string {
	color := t.getColor(t.TableColor)
	return string(color) + strings.Join(row, "\t") + string(reset)
}

// getColor gets a color format string based on a string
func (t *Table) getColor(color string) string {
	switch color {
	case "red":
		return red
	case "reset":
		return reset
	case "green":
		return green
	case "yellow":
		return yellow
	case "blue":
		return blue
	case "cyan":
		return cyan
	case "purple":
		return purple
	case "white":
		return white
	}
	return ""
}

// Print the table to the terminal
func (t *Table) Print() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, t.Padding, '\t', tabwriter.AlignRight)
	for _, row := range t.Rows {

		formatted := ""
		switch row.Type {
		case "header":
			formatted = t.formatHeader(row.Content)
		case "divider":
			formatted = t.formatDivider(row.Content)
		default:
			formatted = t.formatRow(row.Content)
		}
		fmt.Fprintln(writer, formatted)
	}
	writer.Flush()
}
