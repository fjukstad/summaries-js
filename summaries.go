package main

import (
	"fmt"

	"github.com/fjukstad/kvik/eutils"
	"github.com/fjukstad/kvik/genenames"

	"honnef.co/go/js/dom"
)

func main() {
	d := dom.GetWindow().Document()
	s := d.GetElementByID("summary")
	p := d.GetElementByID("symbol").(*dom.HTMLInputElement)
	p.Focus()
	p.AddEventListener("keyup", false, func(event dom.Event) {
		ke := event.(*dom.KeyboardEvent)
		if ke.KeyCode == 13 {
			go func() {
				input := event.Target().(*dom.HTMLInputElement)
				fmt.Println(input.Value)

				doc, err := genenames.GetDoc(input.Value)
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
			}()
		}
	})
}
