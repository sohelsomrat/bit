package export

// ExportFormat represents a supported export format
type ExportFormat struct {
	Name        string // Canonical key (e.g., "TXT")
	Extension   string
	Description string
}

// Available export formats
var SupportedFormats = []ExportFormat{
	{
		Name:        "TXT",
		Extension:   ".txt",
		Description: "Plain text file",
	},
	{
		Name:        "GO",
		Extension:   ".go",
		Description: "Go source code",
	},
	{
		Name:        "JS",
		Extension:   ".js",
		Description: "JavaScript source code",
	},
	{
		Name:        "PY",
		Extension:   ".py",
		Description: "Python source code",
	},
	{
		Name:        "RS",
		Extension:   ".rs",
		Description: "Rust source code",
	},
	{
		Name:        "SH",
		Extension:   ".sh",
		Description: "Bash script",
	},
}
