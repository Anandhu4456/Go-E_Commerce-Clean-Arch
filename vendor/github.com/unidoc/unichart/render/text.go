package render

import (
	"strings"

	"github.com/unidoc/unichart/mathutil"
)

// TextHorizontalAlign is an enum for the horizontal alignment options.
type TextHorizontalAlign int

const (
	// TextHorizontalAlignUnset is an unset state for text horizontal alignment.
	TextHorizontalAlignUnset TextHorizontalAlign = iota

	// TextHorizontalAlignLeft aligns a string horizontally so that it's left
	// ligature starts at horizontal pixel 0.
	TextHorizontalAlignLeft

	// TextHorizontalAlignCenter left aligns a string horizontally so that
	// there are equal pixels to the left and to the right of a string within
	// a box.
	TextHorizontalAlignCenter

	// TextHorizontalAlignRight right aligns a string horizontally so that
	// the right ligature ends at the right-most pixel of a box.
	TextHorizontalAlignRight
)

// TextWrap is an enum for the word wrap options.
type TextWrap int

const (
	// TextWrapUnset is the unset state for text wrap options.
	TextWrapUnset TextWrap = iota

	// TextWrapNone will spill text past horizontal boundaries.
	TextWrapNone

	// TextWrapWord will split a string on words (i.e. spaces) to fit within
	// a horizontal boundary.
	TextWrapWord

	// TextWrapRune will split a string on a rune (i.e. utf-8 codepage) to fit
	// within a horizontal boundary.
	TextWrapRune
)

// TextVerticalAlign is an enum for the vertical alignment options.
type TextVerticalAlign int

const (
	// TextVerticalAlignUnset is the unset state for vertical alignment options.
	TextVerticalAlignUnset TextVerticalAlign = iota

	// TextVerticalAlignBaseline aligns text according to the "baseline" of
	// the string, or where a normal ascender begins.
	TextVerticalAlignBaseline

	// TextVerticalAlignBottom aligns the text according to the lowers pixel
	// of any of the ligatures (ex. g or q both extend below the baseline).
	TextVerticalAlignBottom

	// TextVerticalAlignMiddle aligns the text so that there is an equal
	// amount of space above and below the top and bottom of the ligatures.
	TextVerticalAlignMiddle

	// TextVerticalAlignMiddleBaseline aligns the text vertically so that there
	// is an equal number of pixels above and below the baseline of the string.
	TextVerticalAlignMiddleBaseline

	// TextVerticalAlignTop alignts the text so that the top of the ligatures
	// are at y-pixel 0 in the container.
	TextVerticalAlignTop
)

// TextStyle encapsulates text style options.
type TextStyle struct {
	HorizontalAlign TextHorizontalAlign
	VerticalAlign   TextVerticalAlign
	Wrap            TextWrap
}

type text struct{}

var (
	// Text contains utilities for text.
	Text = &text{}
)

func (t text) Measure(r Renderer, text string, style Style) Box {
	style.GetTextOptions().WriteToRenderer(r)
	defer r.ResetStyle()

	return r.MeasureText(text)
}

// Draw draws text with a given style.
func (t text) Draw(r Renderer, text string, x, y int, style Style) {
	style.GetTextOptions().WriteToRenderer(r)
	defer r.ResetStyle()

	r.Text(text, x, y)
}

// DrawWithin draws the text within a given box.
func (t text) DrawWithin(r Renderer, text string, box Box, style Style) {
	style.GetTextOptions().WriteToRenderer(r)
	defer r.ResetStyle()

	lines := Text.WrapFit(r, text, box.Width(), style)
	linesBox := Text.MeasureLines(r, lines, style)

	y := box.Top

	switch style.GetTextVerticalAlign() {
	case TextVerticalAlignBottom, TextVerticalAlignBaseline:
		y = y - linesBox.Height()
	case TextVerticalAlignMiddle:
		y = y + (box.Height() >> 1) - (linesBox.Height() >> 1)
	case TextVerticalAlignMiddleBaseline:
		y = y + (box.Height() >> 1) - linesBox.Height()
	}

	var tx, ty int
	for _, line := range lines {
		lineBox := r.MeasureText(line)
		switch style.GetTextHorizontalAlign() {
		case TextHorizontalAlignCenter:
			tx = box.Left + ((box.Width() - lineBox.Width()) >> 1)
		case TextHorizontalAlignRight:
			tx = box.Right - lineBox.Width()
		default:
			tx = box.Left
		}
		if style.TextRotationDegrees == 0 {
			ty = y + lineBox.Height()
		} else {
			ty = y
		}

		r.Text(line, tx, ty)
		y += lineBox.Height() + style.GetTextLineSpacing()
	}
}

func (t text) WrapFit(r Renderer, value string, width int, style Style) []string {
	switch style.TextWrap {
	case TextWrapRune:
		return t.WrapFitRune(r, value, width, style)
	case TextWrapWord:
		return t.WrapFitWord(r, value, width, style)
	}

	return []string{value}
}

func (t text) WrapFitWord(r Renderer, value string, width int, style Style) []string {
	style.WriteToRenderer(r)

	var output []string
	var line string
	var word string
	var textBox Box

	for _, c := range value {
		if c == rune('\n') {
			output = append(output, t.Trim(line+word))
			line = ""
			word = ""
			continue
		}

		textBox = r.MeasureText(line + word + string(c))

		if textBox.Width() >= width {
			output = append(output, t.Trim(line))
			line = word
			word = string(c)
			continue
		}

		if c == rune(' ') || c == rune('\t') {
			line = line + word + string(c)
			word = ""
			continue
		}

		word = word + string(c)
	}

	return append(output, t.Trim(line+word))
}

func (t text) WrapFitRune(r Renderer, value string, width int, style Style) []string {
	style.WriteToRenderer(r)

	var output []string
	var line string
	var textBox Box

	for _, c := range value {
		if c == rune('\n') {
			output = append(output, line)
			line = ""
			continue
		}

		textBox = r.MeasureText(line + string(c))

		if textBox.Width() >= width {
			output = append(output, line)
			line = string(c)
			continue
		}
		line = line + string(c)
	}
	return t.appendLast(output, line)
}

func (t text) Trim(value string) string {
	return strings.Trim(value, " \t\n\r")
}

func (t text) MeasureLines(r Renderer, lines []string, style Style) Box {
	style.WriteTextOptionsToRenderer(r)
	var output Box
	for index, line := range lines {
		lineBox := r.MeasureText(line)
		output.Right = mathutil.MaxInt(lineBox.Right, output.Right)
		output.Bottom += lineBox.Height()
		if index < len(lines)-1 {
			output.Bottom += +style.GetTextLineSpacing()
		}
	}
	return output
}

func (t text) appendLast(lines []string, text string) []string {
	if len(lines) == 0 {
		return []string{text}
	}
	lastLine := lines[len(lines)-1]
	lines[len(lines)-1] = lastLine + text
	return lines
}
