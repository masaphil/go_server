package paper

import "github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"

//選択肢
type Option string

//source.Option から paper.optionに変換メソッド
//optionの順番はシャッフル
func FromSourceOption(o []source.Option) []Option {
	source.ShuffleSourceOptions(o)
	slice := make([]Option, 0)
	for _, i := range o {
		slice = append(slice, Option(string(i)))
	}
	return slice
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
