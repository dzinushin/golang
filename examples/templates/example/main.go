package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

func main() {
	//vararg("111", "2222")
	tmpl := `
Name: {{ .Name }}
{{ (test "LLL") }}
by env: '{{ (by_env "dev" "DEVEL" "ote" "OTE" "prod" "PRODUCTION" "default" "SOMESHIT") }}'
`
	envName := "dev"
	InterpolateManifest("k8s", envName, tmpl)
}

func vararg(args ...string) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func InterpolateManifest(name string, envName string, tmpl string) (string, error) {
	var err error
	t := template.New(name)

	t.Funcs(template.FuncMap{
		"test": func(p1 string) string {
			return "LALA" + p1
		},
		"by_env": func(args ...string) string {
			var vals = make(map[string]string)
			var key = ""
			for i, v := range args {
				fmt.Printf("i: %d v: %s\n", i, v)
				if i%2 == 0 {
					key = v
				} else {
					vals[key] = v
				}
			}
			fmt.Println(vals)
			result, present := vals[envName]
			if !present {
				result, present = vals["default"]
				if !present {
					panic(fmt.Sprintf("Can't determine value in 'by_env' template func for env name: %s", envName))
				}
			}
			return result
		},
	})

	t, err = t.Parse(tmpl)
	if err != nil {
		log.Fatalf("Error parse template: %v", err)
	}
	config := map[string]interface{}{
		"Name":      "Name",
		"Version":   "1.0.1",
		"Namespace": "ns",
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, config); err != nil {
		log.Fatalf("Error interpolate app manifest template: %v", err)
	}
	resultManifest := tpl.String()
	log.Printf("interpolated manifest: \n%s\n", resultManifest)

	return resultManifest, nil
}
