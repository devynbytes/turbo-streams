# Turbo Streams Go

This is a simple implementation of the [Turbo Sreams](https://turbo.hotwired.dev/handbook/streams) backend for Go.

## Usage

1. Create or parse a `html/template` to use
1. Create an instance of the `turbo_streams.Turbo` struct
2. Send an http response


```go

  func handleSomePath(w http.ResponseWriter, r *http.Request) {

      // Create or open a template
      tmp, _ := template.New("message").Parse(`<div>{{.}}</div>`)

      // Create the Turbo struct, with one of the 7 turbo streams actions
      turbo := turbo_streams.Turbo{
          Action:   turbo.Append,
          Template: template,
          Target:   "messages",
          Data:     "some message",
      }

      // Sends an http response
      turbo.SendMessage(response, request)
  }

```

