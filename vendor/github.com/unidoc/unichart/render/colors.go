package render

import (
	"fmt"
	"image/color"
)

var (
	// ColorTransparent is a transparent (alpha zero) color.
	ColorTransparent = color.RGBA{R: 255, G: 255, B: 255, A: 0}

	// ColorWhite is white.
	ColorWhite = color.RGBA{R: 255, G: 255, B: 255, A: 255}

	// ColorBlue is the basic theme blue color.
	ColorBlue = color.RGBA{R: 0, G: 116, B: 217, A: 255}

	// ColorCyan is the basic theme cyan color.
	ColorCyan = color.RGBA{R: 0, G: 217, B: 210, A: 255}

	// ColorGreen is the basic theme green color.
	ColorGreen = color.RGBA{R: 0, G: 217, B: 101, A: 255}

	// ColorRed is the basic theme red color.
	ColorRed = color.RGBA{R: 217, G: 0, B: 116, A: 255}

	// ColorOrange is the basic theme orange color.
	ColorOrange = color.RGBA{R: 217, G: 101, B: 0, A: 255}

	// ColorYellow is the basic theme yellow color.
	ColorYellow = color.RGBA{R: 217, G: 210, B: 0, A: 255}

	// ColorBlack is the basic theme black color.
	ColorBlack = color.RGBA{R: 51, G: 51, B: 51, A: 255}

	// ColorLightGray is the basic theme light gray color.
	ColorLightGray = color.RGBA{R: 239, G: 239, B: 239, A: 255}

	// ColorAlternateBlue is a alternate theme color.
	ColorAlternateBlue = color.RGBA{R: 106, G: 195, B: 203, A: 255}

	// ColorAlternateGreen is a alternate theme color.
	ColorAlternateGreen = color.RGBA{R: 42, G: 190, B: 137, A: 255}

	// ColorAlternateGray is a alternate theme color.
	ColorAlternateGray = color.RGBA{R: 110, G: 128, B: 139, A: 255}

	// ColorAlternateYellow is a alternate theme color.
	ColorAlternateYellow = color.RGBA{R: 240, G: 174, B: 90, A: 255}

	// ColorAlternateLightGray is a alternate theme color.
	ColorAlternateLightGray = color.RGBA{R: 187, G: 190, B: 191, A: 255}
)

var (
	// defaultBackgroundColor is the default chart background color.
	defaultBackgroundColor = ColorWhite

	// defaultBackgroundStrokeColor is the default chart border color.
	defaultBackgroundStrokeColor = ColorWhite

	// defaultCanvasColor is the default chart canvas color.
	defaultCanvasColor = ColorWhite

	// defaultCanvasStrokeColor is the default chart canvas stroke color.
	defaultCanvasStrokeColor = ColorWhite
)

var (
	// defaultColors are a couple default series colors.
	defaultColors = []color.Color{
		ColorBlue,
		ColorGreen,
		ColorRed,
		ColorCyan,
		ColorOrange,
	}

	// defaultAlternateColors are a couple alternate colors.
	defaultAlternateColors = []color.Color{
		ColorAlternateBlue,
		ColorAlternateGreen,
		ColorAlternateGray,
		ColorAlternateYellow,
		ColorBlue,
		ColorGreen,
		ColorRed,
		ColorCyan,
		ColorOrange,
	}
)

// ColorWithAlpha returns a copy of the color with a given alpha.
func ColorWithAlpha(c color.RGBA, a uint8) color.Color {
	return color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: a,
	}
}

// ColorIsZero returns true if the all the color components are zero.
func ColorIsZero(c color.Color) bool {
	if c == nil {
		return true
	}

	r, g, b, a := c.RGBA()
	return r == 0 && g == 0 && b == 0 && a == 0
}

// ColorToString returns a string representation of the color.
func ColorToString(c color.Color) string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf("rgba(%d,%d,%d,%d)", r, g, b, a)
}

// GetDefaultColor returns a color from the default list by index.
// NOTE: the index will wrap around (using a modulo).
func GetDefaultColor(index int) color.Color {
	finalIndex := index % len(defaultColors)
	return defaultColors[finalIndex]
}

// GetAlternateColor returns a color from the default list by index.
// NOTE: the index will wrap around (using a modulo).
func GetAlternateColor(index int) color.Color {
	finalIndex := index % len(defaultAlternateColors)
	return defaultAlternateColors[finalIndex]
}

// ColorPalette is a set of colors that.
type ColorPalette interface {
	BackgroundColor() color.Color
	BackgroundStrokeColor() color.Color
	CanvasColor() color.Color
	CanvasStrokeColor() color.Color
	AxisStrokeColor() color.Color
	TextColor() color.Color
	GetSeriesColor(index int) color.Color
}

// DefaultColorPalette is the default color palatte.
var DefaultColorPalette defaultColorPalette

type defaultColorPalette struct{}

func (dp defaultColorPalette) BackgroundColor() color.Color {
	return defaultBackgroundColor
}

func (dp defaultColorPalette) BackgroundStrokeColor() color.Color {
	return defaultBackgroundStrokeColor
}

func (dp defaultColorPalette) CanvasColor() color.Color {
	return defaultCanvasColor
}

func (dp defaultColorPalette) CanvasStrokeColor() color.Color {
	return defaultCanvasStrokeColor
}

func (dp defaultColorPalette) AxisStrokeColor() color.Color {
	return DefaultLineColor
}

func (dp defaultColorPalette) TextColor() color.Color {
	return DefaultTextColor
}

func (dp defaultColorPalette) GetSeriesColor(index int) color.Color {
	return GetDefaultColor(index)
}

// AlternateColorPalette is an alternate color palatte.
var AlternateColorPalette alternateColorPalette

type alternateColorPalette struct{}

func (ap alternateColorPalette) BackgroundColor() color.Color {
	return defaultBackgroundColor
}

func (ap alternateColorPalette) BackgroundStrokeColor() color.Color {
	return defaultBackgroundStrokeColor
}

func (ap alternateColorPalette) CanvasColor() color.Color {
	return defaultCanvasColor
}

func (ap alternateColorPalette) CanvasStrokeColor() color.Color {
	return defaultCanvasStrokeColor
}

func (ap alternateColorPalette) AxisStrokeColor() color.Color {
	return DefaultLineColor
}

func (ap alternateColorPalette) TextColor() color.Color {
	return DefaultTextColor
}

func (ap alternateColorPalette) GetSeriesColor(index int) color.Color {
	return GetAlternateColor(index)
}
