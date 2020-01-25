package wd

import "errors"

/* precomputed items for Joconde database
   source : db dumb in csv format

/* xsv frequency -s "Domaine" joconde.tsv | xsv table
field    value                      count
Domaine  dessin                     233733
Domaine  peinture                   64004
Domaine  estampe                    34775
Domaine  sculpture                  29998
Domaine  archéologie;néolithique    13824
Domaine  photographie;ethnologie    10847
Domaine  photographie               10586
Domaine  céramique                  9731
Domaine  estampe;ethnologie         6821
Domaine  archéologie;âge du bronze  5827
*/

type Dict map[string]string

var domains = Dict{
	"dessin":                    "Q93184",
	"peinture":                  "Q3305213",
	"estampe":                   "Q11060274",
	"sculpture":                 "Q860861",
	"archéologie;néolithique":   "",
	"photographie;ethnologie":   "",
	"photographie":              "Q125191",
	"céramique":                 "Q13464614", //céramique d'art ?
	"estampe;ethnologie":        "",
	"archéologie;âge du bronze": "",
}

// FindDomain : find qid for domain, subclass of visual arts
func FindDomain(domain string) (string, error) {
	if domain == "" {
		return "", nil
	}

	qid, found := domains[domain]
	if !found {
		return "", errors.New("Domain not found : " + domain)
	}
	if qid == "" {
		return "", errors.New("No QID for domain : " + domain)
	}
	return qid, nil
}
