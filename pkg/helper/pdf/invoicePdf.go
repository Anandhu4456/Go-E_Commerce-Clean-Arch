package pdf

import "github.com/unidoc/unipdf/v3/creator"

type Client struct {
	creator *creator.Creator
}

type cellStyle struct {
	ColSpan         int
	HAlignment      creator.CellHorizontalAlignment
	BackgroundColor creator.Color
	BorderSide      creator.CellBorderSide
	BorderStyle     creator.CellBorderStyle
	BorderWidth     float64
	BorderColor     creator.Color
	Indent          float64
}

var cellStyles = map[string]cellStyle{
	"heading-left": {
		BackgroundColor: creator.ColorRGBFromHex("#332f3f"),
		HAlignment:      creator.CellHorizontalAlignmentLeft,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"heading-centered": {
		BackgroundColor: creator.ColorRGBFromHex("#332f3f"),
		HAlignment:      creator.CellHorizontalAlignmentCenter,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"left-highlighted": {
		BackgroundColor: creator.ColorRGBFromHex("#dde4e5"),
		HAlignment:      creator.CellHorizontalAlignmentLeft,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"centered-highlighted": {
		BackgroundColor: creator.ColorRGBFromHex("#dde4e5"),
		HAlignment:      creator.CellHorizontalAlignmentCenter,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"left": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"centered": {
		HAlignment: creator.CellHorizontalAlignmentCenter,
	},
	"gradingsys-head": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"gradingsys-row": {
		HAlignment: creator.CellHorizontalAlignmentCenter,
	},
	"conduct-head": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"conduct-key": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"conduct-val": {
		BackgroundColor: creator.ColorRGBFromHex("#dde4e5"),
		HAlignment:      creator.CellHorizontalAlignmentCenter,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     3,
	},
}
