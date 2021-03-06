package pkg

import "github.com/gosuri/uiprogress"

type Progress struct {
	Bar *uiprogress.Bar
	channel chan string
	State string
}

func (progress *Progress) SetText(text string) {
	if text != "" {
		progress.State = text
	}
}
