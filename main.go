package main

import (
	"github.com/ssimunic/gosensors"
	"github.com/zserge/webview"
)

// Counter is a simple example of automatic Go-to-JS data binding
type Counter struct {
	Value *gosensors.Sensors `json:"value"`
}

// Add increases the value of a counter by n
func (c *Counter) Add(n string) {
	// c.Value = c.Value + n
}

// Update the values
func (c *Counter) Update() {
	c.Value = getSensors()
}

func main() {
	w := webview.New(webview.Settings{Title: "Gosus", Debug: true, Resizable: true})
	defer w.Exit()

	w.Dispatch(func() {
		// Inject data
		w.Bind("counter", &Counter{getSensors()})

		// Inject Compiled Proyect
		w.Eval(string(MustAsset("assets/dist/main.js")))
	})
	w.Run()
}

func getSensors() *gosensors.Sensors {
	sensors, err := gosensors.NewFromSystem()

	if err != nil {
		panic(err)
	}

	return sensors
}
