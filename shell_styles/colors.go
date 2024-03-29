package shell_styles

const escape = "\x1B"
const RESET_ALL = "\x1B[0m"
const fgBlue = "[34m"
const fgGreen = "[32m"
const fgDefault = "[39m"
const fgLightGreen = "[92m"
const fgDarkGray = "[90m"
const bgWhite = "[107m"
const bgGreen = "[42m"
const bgDefault = "[49m"

type Color struct {
	name   string
	prefix string
}

type Reset struct {
	name string
}

type Output interface {
	ToString() string
}

func (color Color) ToString() string {
	return color.prefix
}

func formatWithPrefix(text string) string {
	return escape + text
}

func (color Reset) ToString() string {
	return RESET_ALL
}

var FgDefault Output = Color{name: "fg Default", prefix: formatWithPrefix(fgDefault)}
var FgBlue Output = Color{name: "fg Blue", prefix: formatWithPrefix(fgBlue)}
var FgGreen Output = Color{name: "fg Green", prefix: formatWithPrefix(fgGreen)}
var FgLightGreen Output = Color{name: "fg LightGreen", prefix: formatWithPrefix(fgLightGreen)}
var FgDarkGray Output = Color{name: "fg DarkGray", prefix: formatWithPrefix(fgDarkGray)}
var BgDefault Output = Color{name: "bg default", prefix: formatWithPrefix(bgDefault)}
var BgWhite Output = Color{name: "bg white", prefix: formatWithPrefix(bgWhite)}
var BgGreen Output = Color{name: "bg green", prefix: formatWithPrefix(bgGreen)}
var ResetAll Output = Reset{name: "reset all styles"}
