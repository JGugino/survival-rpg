package engine

import "fmt"

const (
	INFO_PREFIX    = "[ENGINE:INFO]"
	ERROR_PREFIX   = "[ENGINE:ERROR]"
	WARNING_PREFIX = "[ENGINE:WARNING]"
)

func PrintInfo(message string) {
	print(fmt.Sprintf("%s - %s \n", INFO_PREFIX, message))
}

func PrintError(message string) {
	print(fmt.Sprintf("%s - %s \n", ERROR_PREFIX, message))
}

func PrintWarning(message string) {
	print(fmt.Sprintf("%s - %s \n", WARNING_PREFIX, message))
}
