package main

import (
	"fmt"

	"github.com/fjukstad/kvik/eutils"
	"github.com/fjukstad/kvik/genenames"

	"honnef.co/go/js/dom"
)

func search(gene string) {
	d := dom.GetWindow().Document()
	s := d.GetElementByID("summary")

	doc, err := genenames.GetDoc(gene)
	if err != nil {
		s.SetInnerHTML("Could not convert gene symbol to entrez id. " + err.Error())
		return
	}

	geneSummary, err := eutils.GeneSummary(doc.EntrezId)
	if err != nil {
		s.SetInnerHTML(err.Error())
		return
	}
	summary := geneSummary.Summary
	s.SetInnerHTML(summary)
}

func main() {
	d := dom.GetWindow().Document()
	p := d.GetElementByID("symbol").(*dom.HTMLInputElement)
	submit := d.GetElementByID("submit").(*dom.HTMLInputElement)

	p.Focus()
	p.AddEventListener("keyup", false, func(event dom.Event) {
		ke := event.(*dom.KeyboardEvent)
		input := event.Target().(*dom.HTMLInputElement)
		fmt.Println(input.Value)
		gene := input.Value
		if ke.KeyCode == 13 {
			go search(gene)
		}
	})

	submit.AddEventListener("click", false, func(event dom.Event) {
		gene := p.Value
		go search(gene)
	})

}
