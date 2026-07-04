package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var aliases = []struct {
	Name string
}{
	{"Box"}, {"VStack"}, {"HStack"}, {"Text"}, {"Flex"}, {"Center"},
	{"Container"}, {"Separator"}, {"Spacer"},
	{"Heading"}, {"Code"}, {"Link"}, {"Blockquote"}, {"List"}, {"Kbd"}, {"Mark"},
	{"Button"}, {"OutlineButton"}, {"Checkbox"}, {"Switch"}, {"Input"}, {"TextArea"},
	{"Badge"}, {"BadgeSolid"}, {"BadgeSubtle"}, {"BadgeOutline"}, {"BadgeVariant"},
	{"Card"}, {"Alert"}, {"Spinner"}, {"Avatar"}, {"Progress"},
	{"Dialog"}, {"DialogContent"}, {"Menu"}, {"MenuItem"}, {"Tooltip"}, {"Popper"},
	{"Window"}, {"App"}, {"Portals"}, {"Portal"},
}

const tmpl = `package main

import "raitui"

var (
{{- range .}}
	{{.Name | pad}} = raitui.{{.Name}}
{{- end}}
)
`

func main() {
	funcs := template.FuncMap{
		"pad": func(s string) string {
			max := 0
			for _, a := range aliases {
				if len(a.Name) > max {
					max = len(a.Name)
				}
			}
			return s + strings.Repeat(" ", max-len(s)+1)
		},
	}

	t := template.Must(template.New("aliases").Funcs(funcs).Parse(tmpl))

	path := "aliases.go"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "raitui: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := t.Execute(f, aliases); err != nil {
		fmt.Fprintf(os.Stderr, "raitui: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s\n", path)
}
