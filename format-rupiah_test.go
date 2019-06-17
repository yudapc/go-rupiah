package rupiah

import "testing"

func TestFormatRupiah(t *testing.T) {
	amount := float64(10000)
	expected := "Rp 10.000"
	res := FormatRupiah(amount)
	if res != expected {
		t.Errorf("Our FormatRupiah function doesn't work, %f isn't %s\n", amount, res)
	}
}
