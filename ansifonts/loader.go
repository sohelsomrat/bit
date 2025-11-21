// Package ansifonts provides a library for rendering text using ANSI art fonts.
//
// The loader package handles font discovery and loading from the fonts directory.
package ansifonts

import (
	"embed"
	"encoding/json"
	"path"
	"strings"
)

//go:embed fonts/*.bit
var EmbeddedFonts embed.FS

// LoadFont loads a font by name from the embedded fonts directory
func LoadFont(name string) (*Font, error) {
	// Construct the path using forward slashes for embedded filesystem
	fontPath := path.Join("fonts", name+".bit")

	fontBytes, err := EmbeddedFonts.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}

	var fontData FontData
	err = json.Unmarshal(fontBytes, &fontData)
	if err != nil {
		return nil, err
	}

	return &Font{
		Name:     fontData.Name,
		FontData: fontData,
	}, nil
}

// ListFonts returns a list of available font names from the embedded fonts directory
func ListFonts() ([]string, error) {
	entries, err := EmbeddedFonts.ReadDir("fonts")
	if err != nil {
		return nil, err
	}

	var fonts []string
	for _, entry := range entries {
		if !entry.IsDir() && path.Ext(entry.Name()) == ".bit" {
			fontName := strings.TrimSuffix(entry.Name(), ".bit")
			fonts = append(fonts, fontName)
		}
	}

	return fonts, nil
}
