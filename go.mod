module github.com/sp-prog/go-ioc-container

go 1.22

//структура папок взята из:
//- https://github.com/golang-standards/project-layout/blob/master/README_ru.md
//- https://go.dev/doc/modules/layout
//- https://pkg.go.dev/cmd/go#hdr-Internal_Directories
//- книга ТЕЙВА ХАРШАНИ 2024 100 ошибок Go И КАК ИХ ИЗБЕЖАТЬ. Ошибка №12

require (
	github.com/brianvoe/gofakeit v3.18.0+incompatible
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
