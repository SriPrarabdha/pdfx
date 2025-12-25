package pdf

import "runtime"

func ghostscriptCmd() string {
	if runtime.GOOS == "windows" {
		return "gswin64c"
	}
	return "gs"
}

func libreOfficeCmd() string {
	if runtime.GOOS == "windows" {
		return "soffice"
	}
	return "libreoffice"
}
