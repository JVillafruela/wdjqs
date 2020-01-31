package wd

import (
	"fmt"
	"os"
)

// PropWithLang : properties needing language qualifier
var PropWithLang = [...]string{"P1476"}

// PropertyValue : property value with qualifiers and references
type PropertyValue struct {
	Value      string
	Qualifiers map[string]string
	Sources    map[string]string
}

// Item : a WD item
type Item struct {
	// Qid of item, empty for creation
	Qid string
	// language label & descriptions are written into eg. "fr"
	Lang        string
	Label       string
	Description string
	Properties  map[string]PropertyValue
}

// AddProperty : add a property to item
func (it *Item) AddProperty(property, value string) {
	if it.Properties == nil {
		it.Properties = make(map[string]PropertyValue)
	}
	pv, ok := it.Properties[property]
	if !ok {
		pv = PropertyValue{Qualifiers: make(map[string]string), Sources: make(map[string]string)}
	}
	pv.Value = value
	it.Properties[property] = pv
}

// AddQualifier : add a qualifier to a property. Do nothing if property doesn't exist
func (it *Item) AddQualifier(p, q, value string) {
	_, ok := it.Properties[p]
	if !ok {
		return
	}
	it.Properties[p].Qualifiers[q] = value
}

// AddSource : add a source (reference) to a property. Do nothing if property doesn't exist
func (it *Item) AddSource(p, s, value string) {
	_, ok := it.Properties[p]
	if !ok {
		return
	}
	it.Properties[p].Sources[s] = value
}

// AddSources : add several sources (references) to a property. Do nothing if property doesn't exist
func (it *Item) AddSources(p string, sv map[string]string) {
	_, ok := it.Properties[p]
	if !ok {
		return
	}
	for s, v := range sv {
		it.Properties[p].Sources[s] = v
	}
}

// AddSourcesToAll : add several sources (references) to each item property.
func (it *Item) AddSourcesToAll(sv map[string]string) {
	for p := range it.Properties {
		for s, v := range sv {
			it.Properties[p].Sources[s] = v
		}
	}
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
		out.WriteString(fmt.Sprintf("LAST\t%s\t%s%s", "L"+it.Lang, it.Label, eol))
		out.WriteString(fmt.Sprintf("LAST\t%s\t%s%s", "D"+it.Lang, it.Description, eol))
	}

	for prop, pv := range it.Properties {
		lang := ""
		value := pv.Value
		if value == "" {
			continue
		}
		if value[0:1] != "Q" && value != UnknownValue {
			if IsPropertyLang(prop) {
				lang = it.Lang
			}
			value = quote(value, lang)
		}
		qstr := it.formatQualifiers(prop)
		sstr := it.formatSources(prop)
		line := fmt.Sprintf("%s\t%s\t%s%s%s%s", qid, prop, value, qstr, sstr, eol)
		_, err := out.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

// format the qualifiers section for property if any
func (it *Item) formatQualifiers(prop string) string {
	return formatQS(it.Properties[prop].Qualifiers, it.Lang, false)
}

// format the sources section for property if any
func (it *Item) formatSources(prop string) string {
	return formatQS(it.Properties[prop].Sources, it.Lang, true)
}

func formatQS(props map[string]string, lang string, source bool) string {
	str := ""
	for id, v := range props {
		value := v
		if value[0:1] != "Q" && value[0:1] != "+" { //TODO Regexp for Qid & date
			if !IsPropertyLang(id) {
				lang = ""
			}
			value = quote(value, lang)
		}

		if source && id[0:1] == "P" {
			id = "S" + id[1:]
		}

		str = fmt.Sprintf("%s\t%s\t%s", str, id, value)
	}
	return str
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
