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

type dict map[string]string

var domains = dict{
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

/*
xsv frequency -s 17 joconde.tsv | xsv table
field                 value                       count
Matériaux-techniques  (NULL)                      56700
Matériaux-techniques  mine de plomb               33175
Matériaux-techniques  peinture à l'huile;toile    24905
Matériaux-techniques  peinture à l'huile, toile   12782
Matériaux-techniques  silex                       7395
Matériaux-techniques  bronze                      6190
Matériaux-techniques  plâtre                      6032
Matériaux-techniques  peinture à l'huile (toile)  5630
Matériaux-techniques  terre cuite                 5586
Matériaux-techniques  fer                         5553
*/

var materials = dict{
	"(NULL)":                     "",
	"bronze":                     "Q34095", // copper alloy
	"fer":                        "",
	"mine de plomb":              "Q868239", //TODO toile support de peinture
	"peinture à l'huile (toile)": "Q296955",
	"peinture à l'huile;toile":   "Q296955",
	"peinture à l'huile, toile":  "Q296955",
	"plâtre":                     "Q3392817",
	"silex":                      "Q83087",
	"terre cuite":                "Q60424",
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

// FindMaterial : find qid for material
func FindMaterial(material string) (string, error) {
	if material == "" || material == "(NULL)" {
		return "", nil
	}

	qid, found := materials[material]
	if !found {
		return "", errors.New("Material not found : " + material)
	}
	if qid == "" {
		return "", errors.New("No QID for material : " + material)
	}
	return qid, nil
}
