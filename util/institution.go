package util

func Institution(nuwb string) string {
	institutionCode := string([]rune(nuwb)[0])
	if institutionCode == "1" {
		kodeSD := string([]rune(nuwb)[1])
		if kodeSD == "1" {
			return "SD IT PUTRA"
		} else if kodeSD == "2" {
			return "SD IT PUTRI"
		}
	} else if institutionCode == "2" {
		return "SMP IT PUTRI"
	} else if institutionCode == "3" {
		return "SMA IT PUTRI"
	} else if institutionCode == "4" {
		return "MA PLUS"
	} else if institutionCode == "5" {
		return "SMP IT PUTRA"
	} else if institutionCode == "6" {
		return "IDAD"
	} else if institutionCode == "7" {
		return "SMP IT FULLDAY"
	} else if institutionCode == "8" {
		return "SMA IT FULLDAY"
	} else if institutionCode == "9" {
		return "DINIYYAH"
	}
	return ""
}
