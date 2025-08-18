# i18n

## Formato monetario

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

## Idiomas

```go
    f := i18n.New()
	if err := f.Langs("./languages"); err != nil {
		t.Fatalf("Error loading languages: %v", err)
	}

    f.SetLang("es-ES")
    fmt.Println(  f.T("hola") )
```

En la carpeta indicada, se espera ficheros como `es-ES.json`, etc, con el formato:

```json
{
    "translation": {
        "hola": "Hola",
        "adios": "Adiós"
    }
}
```

