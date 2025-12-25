package pdf

import "runtime"

// ghostscript executable name per platform
func ghostscriptCmd() string {
	if runtime.GOOS == "windows" {
		return "gswin64c"
	}
	return "gs"
}

// libreoffice executable name per platform
func libreOfficeCmd() string {
	if runtime.GOOS == "windows" {
		return "soffice"
	}
	return "libreoffice"
}
