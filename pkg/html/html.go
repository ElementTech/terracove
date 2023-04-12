package html

import (
	_ "embed"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/jatalocks/terracove/internal/types"
)

func CreateHTML(suites []types.TerraformModuleStatus, path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()
	tmpl, err := template.New("htmlTmpl").Funcs(template.FuncMap{
		"strip": func(s string) string {
			var result strings.Builder
			for i := 0; i < len(s); i++ {
				b := s[i]
				if ('a' <= b && b <= 'z') ||
					('A' <= b && b <= 'Z') ||
					('0' <= b && b <= '9') ||
					b == ' ' {
					result.WriteByte(b)
				}
			}
			return result.String()
		},
		"formatTime": func(s string) string {
			t, err := time.Parse(time.RFC3339, s)
			if err != nil {
				panic(err)
			}

			humanReadable := t.Format("Monday, Jan 2, 2006 at 3:04pm")
			return humanReadable
		},
	}).Parse(htmlTmpl)

	if err != nil {
		return err
	}

	err = tmpl.Execute(file, suites)

	if err != nil {
		return err
	}

	return nil
}
