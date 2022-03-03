package source

import (
	"errors"
	"math/rand"
)

type Option string

//[]optionに同じ要素が存在していないことを確認
func checkDuplicatedOptions(ol []Option) error {
	uniq := ol
	for _, o := range ol {
		cnt := 0
		for _, u := range uniq {
			if o == u {
				cnt += 1
			}
			if cnt > 1 {
				return errors.New("[]optionに重複した要素が存在します")
			}
		}
	}
	return nil
}

//[]optionに答えのoptionが含まれているか確認
func checkContains(ol []Option, a Option) error {

	for _, v := range ol {
		if a == v {
			return nil
		}
	}

	return errors.New("[]optionに正答が含まれていません")
}

//source.optionリストの要素をシャッフル
func ShuffleSourceOptions(data []Option) []Option {

	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	return data
}

//stringリストからoptionリストへ変換
func ToOptionList(data []string) []Option {

	slice := make([]Option, 0)
	for _, v := range data {
		slice = append(slice, Option(v))
	}

	return slice
}

//optionリストからstringリストにする
func ToStringListOption(data []Option) []string {
	slice := make([]string, 0)
	for _, data := range data {
		slice = append(slice, string(data))
	}
	return slice
}
