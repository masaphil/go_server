package source

type SkillID string

func ToSkillIDList(data []string) []SkillID {
	slice := make([]SkillID, 0)
	for _, st := range data {
		slice = append(slice, SkillID(st))
	}
	return slice
}

//slillIdリストからstringリストへ
func ToStringListSkillID(data []SkillID) []string {
	slice := make([]string, 0)
	for _, data := range data {
		slice = append(slice, string(data))
	}
	return slice
}
