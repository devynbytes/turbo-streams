package turbo_streams_test

import (
	"github.com/devybytes/turbo-streams-go"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTest(
	_ *testing.T,
	action turbo_streams.Action,
	template *template.Template,
	target string,
	data string) (func(t *testing.T), turbo_streams.Turbo, *httptest.ResponseRecorder, *http.Request) {

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	turbo := turbo_streams.Turbo{
		Action:   action,
		Template: template,
		Target:   target,
		Data:     data,
	}

	return func(t *testing.T) {}, turbo, recorder, request
}

func TestTurbo_SendMessage(t *testing.T) {
	t.Run("sets the turbo header", func(t *testing.T) {
		tmp, _ := template.New("message").Parse(`<div>{{.}}</div>`)
		teardownSuite, turbo, recorder, request := setupTest(t, turbo_streams.Append, tmp, "foo", "foobar")
		defer teardownSuite(t)
		turbo.SendMessage(recorder, request)

		got := recorder.Header().Get("Content-Type")
		want := "text/vnd.turbo-stream.html"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("sets the data", func(t *testing.T) {
		tmp, _ := template.New("message").Parse(`<div>{{.}}</div>`)
		teardownSuite, turbo, recorder, request := setupTest(t, turbo_streams.Append, tmp, "foo", "foobar")
		defer teardownSuite(t)
		turbo.SendMessage(recorder, request)

		got := recorder.Body.String()
		//want := "<div>foobar</div>"
		want := `
<turbo-stream action="append" target="foo">
	<template>
		<div>foobar</div>
	</template>
</turbo-stream>
`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
