package rnatranscription

func ToRNA(dna string) string {
	changed := ""
	for _, content := range dna {
		switch {
		case string(content) == "A":
			changed += "U"
		case string(content) == "C":
			changed += "G"
		case string(content) == "G":
			changed += "C"
		case string(content) == "T":
			changed += "A"
		}
	}
	return changed
}
