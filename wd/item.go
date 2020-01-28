package wd

import (
	"fmt"
	"os"
)

// QSandbox : The Sandbox item
const QSandbox = "Q4115189"

// PropWithLang : properties needing language qualifier
var PropWithLang = [...]string{"P1476"}

// Item : a WD item (missing qualifiers)
type Item struct {
	// Qid of item, empty for creation
	Qid string
	// language label & descriptions are written into eg. "fr"
	Lang        string
	Label       string
	Description string
	Properties  map[string]string
}

// Add : add a property to item
func (it *Item) Add(property, value string) {
	if it.Properties == nil {
		it.Properties = make(map[string]string)
	}
	it.Properties[property] = value
}

// WriteQS : writes QuickStatements in file
// https://www.wikidata.org/wiki/Help:QuickStatements
func (it *Item) WriteQS(fname string) error {
	const eol = "\n"
	out, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer out.Close()

	qid := it.Qid
	if qid == "" {
		qid = "LAST" // Item creation
		out.WriteString("CREATE " + eol)
	}

	for prop, value := range it.Properties {
		lang := ""
		if value[0:1] != "Q" {
			if IsPropertyLang(prop) {
				lang = it.Lang
			}
			value = quote(value, lang)
		}
		line := fmt.Sprintf("%s\t%s\t%s%s", qid, prop, value, eol)
		_, err := out.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

//Quote string, prefixing with lang code if needed
func quote(str, lang string) string {
	const dq = "\""
	var lp string
	if lang != "" && str != "" {
		lp = lang + ":"
	}
	return lp + dq + str + dq
}

//IsPropertyLang : returns true if property needs a language qualifier
func IsPropertyLang(prop string) bool {
	for _, p := range PropWithLang {
		if prop == p {
			return true
		}
	}
	return false
}
