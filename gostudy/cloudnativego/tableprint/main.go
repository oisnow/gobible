package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		[]string{"haoyunpeng", "31", "man"},
		[]string{"gaolian", "32", "woman"},
		[]string{"happy", "31", "girl"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Age", "Sex"})
	table.AppendBulk(data)
	table.Render()

}
