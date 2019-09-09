package digitalclock

import "strings"

func mergeLines(outString []string) string {
	return strings.Join(outString[:], "\n")
}

func renderWidget(widgetSettings Settings) string {
	outputStrings := []string{}
	clockString, needBorder := renderClock(widgetSettings)
	if needBorder {
		outputStrings = append(outputStrings, mergeLines([]string{"", clockString, ""}))
	} else {
		outputStrings = append(outputStrings, clockString)
	}
	outputStrings = append(outputStrings, getDate())
	outputStrings = append(outputStrings, getUTC())
	outputStrings = append(outputStrings, getEpoch())
	return mergeLines(outputStrings)
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, renderWidget(*widget.settings), false
	})

}
