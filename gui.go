package main

import (
	"github.com/lxn/walk"
	g "github.com/lxn/walk/declarative"
)

func gui() {
	var te *walk.TextEdit
	mw := g.MainWindow{
		Name:    "MainWindow",
		MinSize: g.Size{640, 480},
		Layout:  g.Grid{Rows: 2},
		Children: []g.Widget{
			g.TextEdit{AssignTo: &te},
			g.Label{Text: "Hello"},
		},
	}
	mw.Run()

}
