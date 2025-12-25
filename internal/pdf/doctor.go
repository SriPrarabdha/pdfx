package pdf

import "os/exec"

type DependencyStatus struct {
	Name        string
	Command     string
	Description string
	Available   bool
}

func CheckDependencies() []DependencyStatus {
	deps := []DependencyStatus{
		{
			Name:        "Ghostscript",
			Command:     ghostscriptCmd(),
			Description: "PDF compression",
		},
		{
			Name:        "Pandoc",
			Command:     "pandoc",
			Description: "Text/Markdown to PDF",
		},
		{
			Name:        "LibreOffice",
			Command:     libreOfficeCmd(),
			Description: "DOCX/PPTX/XLSX to PDF",
		},
	}

	for i := range deps {
		_, err := exec.LookPath(deps[i].Command)
		deps[i].Available = (err == nil)
	}

	return deps
}
