package turbo_streams

import (
	"bytes"
	"html/template"
	"net/http"
)

const turboStream = `
<turbo-stream action="{{.Action}}"{{if .Target}} target="{{.Target}}"{{end}}{{if .Targets}} targets="{{.Targets}}"{{end}}>
	<template>
		{{template "[[.]]" .Data}}
	</template>
</turbo-stream>
`

type Action string

const (
	After   Action = "after"
	Append  Action = "append"
	Before  Action = "before"
	Prepend Action = "prepend"
	Remove  Action = "remove"
	Replace Action = "replace"
	Update  Action = "update"
)

type Turbo struct {
	Action   Action
	Template *template.Template
	Target   string
	Targets  string
	Data     interface{}
}

func (t *Turbo) SendMessage(w http.ResponseWriter, r *http.Request) {
	var turboTemplate, err = t.wrapInTurboStream("message")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	wrapped, err := t.Template.New("data").Parse(turboTemplate)

	if err != nil {
		http.Error(w, "Error parsing template", 500)
	}

	t.Template = wrapped
	t.sendHTTP(w)
}

// sendHTTP sends a turbo template as an HTTP response
func (t *Turbo) sendHTTP(rw http.ResponseWriter) {
	rw.Header().Add("Content-type", "text/vnd.turbo-stream.html")
	err := t.Template.Execute(rw, t)
	if err != nil {
		return
	}
}

//func (h *Turbo) sendSocket(hub *Hub) {
//	var buf bytes.Buffer
//	h.Template.Execute(&buf, h)
//	hub.broadcast <- buf.Bytes()
//}

func (t *Turbo) wrapInTurboStream(name string) (string, error) {
	var buffer bytes.Buffer
	stream, _ := template.New("stream").Delims("[[", "]]").Parse(turboStream)
	err := stream.Execute(&buffer, name)
	return buffer.String(), err
}
