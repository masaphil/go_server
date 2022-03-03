package valueobject

import "testing"

func TestGenerateNewID(t *testing.T) {
	id := GenerateNewID()
	idStr, err := NewID(id.String())
	if err != nil {
		t.Errorf("parse error: id is %s", id.String())
	}
	t.Log(idStr)

	if _, err = NewID("aaaaaaa"); err == nil {
		t.Error("文字列の形式が違った場合は成功しません")
	}
}
