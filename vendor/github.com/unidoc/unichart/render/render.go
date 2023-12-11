package render

import (
	"image/color"
	"io"
)

// Renderable is a function that can be called to render custom
// elements on the chart.
type Renderable func(r Renderer, canvasBox Box, defaults Style)

// Renderer represents a chart renderer.
type Renderer interface {
	// ResetStyle resets all the style related settings of the renderer.
	ResetStyle()

	// GetDPI gets the DPI for the renderer.
	GetDPI() float64

	// SetDPI sets the DPI for the renderer.
	SetDPI(dpi float64)

	// SetClassName sets the current class name.
	SetClassName(string)

	// SetStrokeColor sets the current stroke color.
	SetStrokeColor(color.Color)

	// SetFillColor sets the current fill color.
	SetFillColor(color.Color)

	// SetStrokeWidth sets the stroke width.
	SetStrokeWidth(width float64)

	// SetStrokeDashArray sets the stroke dash array.
	SetStrokeDashArray(dashArray []float64)

	// MoveTo moves the cursor to the specified point.
	MoveTo(x, y int)

	// LineTo draws a line to the specified point, starting from the previous one.
	LineTo(x, y int)

	// QuadCurveTo draws a quad curve. `cx` and `cy` are the Bézier control points.
	QuadCurveTo(cx, cy, x, y int)

	// ArcTo draws an arc with a given center (`cx`, `cy`), a given set of
	// radii (`rx`, `ry`), a `startAngle` and `deltaAngle` (in radians).
	ArcTo(cx, cy int, rx, ry, startAngle, delta float64)

	// Close finalizes a shape, closing the path.
	Close()

	// Stroke strokes the current path.
	Stroke()

	// Fill fills the current path.
	Fill()

	// FillStroke fills and strokes the current path.
	FillStroke()

	// Circle draws a circle at the given coordinates, with a given radius.
	Circle(radius float64, x, y int)

	// SetFont sets the current font.
	SetFont(font Font)

	// SetFontColor sets the current font color.
	SetFontColor(color.Color)

	// SetFontSize sets the current font size.
	SetFontSize(size float64)

	// Text draws a text chunk.
	Text(body string, x, y int)

	// MeasureText measures the specified text.
	MeasureText(body string) Box

	// SetTextRotation sets the rotation of the text.
	SetTextRotation(radians float64)

	// ClearTextRotation clears rotation of the text.
	ClearTextRotation()

	// Save saves the rendered data to the given writer.
	Save(w io.Writer) error
}

// RendererProvider is a function that returns a renderer.
type RendererProvider func(int, int) (Renderer, error)

// ChartRenderable represents a chart renderable component.
type ChartRenderable interface {
	Width() int
	SetWidth(width int)

	Height() int
	SetHeight(height int)

	Render(rp RendererProvider, w io.Writer) error
}
