package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"scoober/database"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var red = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}

func main() {
	config := LoadConfig()
	database.StartDB(config.Database.Host, config.Database.Port,
		config.Database.User, config.Database.Password, config.Database.Dbname)
	initGUI()
}

func initGUI() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Scoober"),
			app.Size(unit.Dp(800), unit.Dp(600)),
		)
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type C = layout.Context
type D = layout.Dimensions

func draw(w *app.Window) error {
	var ops op.Ops
	var startButton widget.Clickable
	th := material.NewTheme(gofont.Collection())
	var activeGame bool = false

	// listen for events in the window.
	for e := range w.Events() {
		switch e := e.(type) {
		// window rendering
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			if startButton.Clicked() {
				activeGame = true
				fmt.Println("Button has been clicked and game has started")
			}
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					addTeamBoxes,
				),
				// add each layout item below
				layout.Rigid(
					func(gtx C) D {
						// define the margins for the start button
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(35),
							Left:   unit.Dp(35),
						}
						return margins.Layout(gtx,
							// add the start button
							func(gtx C) D {
								var text string
								if !activeGame {
									text = "Start Game"
								} else {
									text = "Stop Game"
								}
								btn := material.Button(th, &startButton, text)
								return btn.Layout(gtx)
							},
						)
					},
				),
			)

			e.Frame(gtx.Ops)
		// window destroy
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func addTeamBoxes(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Spacing:   layout.SpaceSides,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(300, 450), red)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(300, 450), red)
		}),
	)
}
