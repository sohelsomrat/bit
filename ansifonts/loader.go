// Package ansifonts provides a library for rendering text using ANSI art fonts.
//
// The loader package handles font discovery and loading from the fonts directory.
package ansifonts

import (
	"embed"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed fonts/*.bit
var EmbeddedFonts embed.FS

// findFontsDir searches for the fonts directory relative to the package
func findFontsDir() (string, error) {
	// Try package-relative fonts directory first (for library usage)
	// This assumes the fonts are in ansifonts/fonts/
	if _, err := os.Stat("ansifonts/fonts"); err == nil {
		return "ansifonts/fonts", nil
	}

	// Try current directory (for when running from ansifonts/ directory)
	if _, err := os.Stat("fonts"); err == nil {
		return "fonts", nil
	}

	// Try parent directory (for cmd/ executables)
	if _, err := os.Stat("../ansifonts/fonts"); err == nil {
		return "../ansifonts/fonts", nil
	}

	// Legacy fallback for old structure
	if _, err := os.Stat("../fonts"); err == nil {
		return "../fonts", nil
	}

	return "", os.ErrNotExist
}

// LoadFont loads a font by name from the fonts directory
func LoadFont(name string) (*Font, error) {
	fontsDir, err := findFontsDir()
	if err != nil {
		return nil, err
	}

	// Try to load from the fonts directory
	fontPath := filepath.Join(fontsDir, name+".bit")

	// Check if the file exists
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		// Try with different case variations
		titleCaser := cases.Title(language.English)
		variations := []string{
			strings.ToLower(name) + ".bit",
			strings.ToUpper(name) + ".bit",
			titleCaser.String(strings.ToLower(name)) + ".bit",
		}

		for _, variation := range variations {
			fontPath = filepath.Join(fontsDir, variation)
			if _, err := os.Stat(fontPath); err == nil {
				break
			}
		}
	}

	fontBytes, err := os.ReadFile(fontPath)
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

// ListFonts returns a list of available font names from the fonts directory
func ListFonts() ([]string, error) {
	fontsDir, err := findFontsDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(fontsDir)
	if err != nil {
		return nil, err
	}

	var fonts []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".bit" {
			fontName := strings.TrimSuffix(entry.Name(), ".bit")
			fonts = append(fonts, fontName)
		}
	}

	return fonts, nil
}
