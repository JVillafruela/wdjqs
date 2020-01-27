package wd

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
