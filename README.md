# i18n

Ejemplo básico:

```go
	f := i18n.New()
	f.SetMoney(CURRENCY_EUR)
```

Importe en céntimos/centavos

```go
    var i64 int64 = 123456
    log.Println( f.Cent(&i64) )
```

Mostrara `1.234,56 €`

