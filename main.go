package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	_ "github.com/lib/pq"
)

func main() {
	config := LoadConfig()
	initDatabase(config.Database.Host, config.Database.Port,
		config.Database.User, config.Database.Password, config.Database.Dbname)
	initGUI()
}

func initDatabase(host string, port string, user string, password string, dbname string) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Scoober")
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
