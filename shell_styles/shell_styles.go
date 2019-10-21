package shell_styles

const escape = "\x1B"
const reset_all = "\x1B[0m"
const blue = "[34m"
const green = "[32m"

type Color struct {
	name   string
	prefix string
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

var Blue Output = Color{name: "blue", prefix: formatWithPrefix(blue)}
var Green Output = Color{name: "green", prefix: formatWithPrefix(green)}
