// Package ansifonts provides a library for rendering text using ANSI art fonts.
//
// The renderer package handles text rendering with various formatting options.
package ansifonts

// RenderText renders text using the specified font with default options
func RenderText(text string, font *Font) []string {
	// Use default options for basic rendering
	options := DefaultRenderOptions()

	// Render using the local implementation
	return RenderTextWithFont(text, font.FontData, options)
}

// RenderTextWithColor renders text using the specified font and applies ANSI color
func RenderTextWithColor(text string, font *Font, colorCode string) []string {
	// Get hex color from ANSI code using the canonical color map, default to white
	hexColor := "#FFFFFF"
	if color, ok := ANSIColorMap[colorCode]; ok {
		hexColor = color
	}

	// Use options with the specified color
	options := DefaultRenderOptions()
	options.TextColor = hexColor

	// Render using the local implementation
	return RenderTextWithFont(text, font.FontData, options)
}

// RenderTextWithOptions renders text using the specified font and advanced options
func RenderTextWithOptions(text string, font *Font, options RenderOptions) []string {
	// Render using the local implementation
	return RenderTextWithFont(text, font.FontData, options)
}
