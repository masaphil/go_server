package paper

type Result struct {

	//正答数
	Point *int

	//偏差値
	Deviation *float32

	//問題数
	Volume int

	//平均正答数
	Average float32

	//TODO:標準偏差 valueobjectにする
	Hyojun float32
}

func initResult(a float32, ql *[]*Quiz, h float32) *Result {
	//TODO: point>volumeを禁止する制約

	v := len(*ql)

	return &Result{
		Point: nil,

		Deviation: nil,

		Volume: v,

		Average: a,

		Hyojun: h,
	}
}

//result完成させるメソッド
func (r *Result) completeResult(p int) *Result {
	dev := calculateDeviation(p, r.Average, r.Hyojun)
	return &Result{
		Point:     &p,
		Deviation: dev,
		Volume:    r.Volume,
		Average:   r.Average,
		Hyojun:    r.Hyojun,
	}
}

//偏差値を算出するメソッド
func calculateDeviation(p int, ave float32, hyojun float32) *float32 {

	//標準偏差算出
	deviation := (float32((float32(p)-ave))/hyojun)*10 + 50

	return &deviation
}
