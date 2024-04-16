# validate

A biblioteca `validate` possibilita validar valores dos campos de uma structs através de tags.

O `validate` não apenas permite validações específicas para os valores como CPF, CNPJ e
strings em branco, mas também amplia as funcionalidades do pacote `go-playground/validator`.

## Tags adicionais:

| Tag   | Descrição                           |
|-------|-------------------------------------|
| cpf   | Valida se o campo é um CPF válido.  |
| cnpj  | Valida se o campo é um CNPJ válido. |
| blank | Verifica se o campo está em branco. |

Veja a documentação completa com todas as tags disponível em [https://github.com/go-playground/validator](https://github.com/go-playground/validator/blob/master/README.md)
## Instalação
Para obter:
```bash
go get github.com/go-lets-go/validate@v1.0.1
go mod tidy
```
Importando:
```bash
import "github.com/go-lets-go/validate"
```
## Uso:
```go
validator := validate.NewValidate()
validations, err := validator.Struct(person)
```

## Implementação:
- [Exemplo API]()

## Referência

- [klassmann/cpfcnpj](https://github.com/klassmann/cpfcnpj)
- [go-playground/validator](https://github.com/go-playground/validator)