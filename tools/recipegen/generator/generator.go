package generator

import "strings"

type Entry map[string]string

func (e Entry) ReplaceVars(template string) string {
	for k, v := range e {
		template = strings.ReplaceAll(template, "{"+k+"}", v)
	}
	return template
}
