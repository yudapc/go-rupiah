# go-rupiah

This is helper for Golang format rupiah

Import:

```
import "github.com/yudapc/go-rupiah"
```

Example format from number (float64) to rupiah:

```
  amount := 1008050000
  amountFloat := float64(amount)
  formatRupiah := rupiah.FormatRupiah(amountFloat)
  fmt.Println(formatRupiah)  // Output Rp 1.008.050.000
```

Example from rupiah to amount int64:

```
  rupiah := "Rp 1.008.050.000"
  getAmount := rupiah.GetAmount(formatRupiah)
  fmt.Println(getAmount) // Output 1008050000
```

Run Test:

```
go test
```

Check coverage:

```
go test -coverprofile=coverage.out
```

Preview coverage test:

```
go tool cover -html=coverage.out
```
