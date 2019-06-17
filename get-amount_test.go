package rupiah

import "testing"

func TestGetAmount(t *testing.T) {
	rupiah := "Rp 10.000"
	expected := int64(10000)
	res := GetAmount(rupiah)
	if res != expected {
		t.Errorf("Our GetAmount function doesn't work, %s isn't %d\n", rupiah, res)
	}
}
