package renderer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

type Renderer struct {
	templates *template.Template
}

func NewRenderer(path string) *Renderer {
	t, err := template.ParseGlob(path + "*")
	if err != nil {
		panic(err)
	}
	return &Renderer{
		templates: t,
	}
}

func (r *Renderer) Render(id string, data []byte) ([]byte, error) {
	var buf bytes.Buffer
	switch id {
	case "1":
		var t templateOne
		err := json.Unmarshal(data, &t)
		if err != nil {
			return nil, err
		}
		if t.Docker_file == nil {
			return nil, fmt.Errorf("missing required parameters")
		}
		r.templates.ExecuteTemplate(&buf, "template1.yml", t)
	case "2":
		var t templateTwo
		err := json.Unmarshal(data, &t)
		if err != nil {
			return nil, err
		}
		if t.Command == nil || t.Sub_command == nil {
			return nil, fmt.Errorf("missing required parameters")
		}
		r.templates.ExecuteTemplate(&buf, "template2.yml", t)
	default:
		return nil, fmt.Errorf("invalid template id")
	}
	return buf.Bytes(), nil
}
