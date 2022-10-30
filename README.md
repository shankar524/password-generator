[![Coverage Status](https://coveralls.io/repos/github/shankar524/password-generator/badge.svg)](https://coveralls.io/github/shankar524/password-generator)
[![Go Report Card](https://goreportcard.com/badge/github.com/jandelgado/golang-ci-template-github-actions)](https://goreportcard.com/report/github.com/jandelgado/golang-ci-template-github-actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
# password-generator
`password-generator` is command line interface app to generate random password. This app can be used to generate any random text from provided rules.

### requirement
Should have go installed. Go can be downloaded from [here](https://go.dev/dl/).

### App usage

- Clone repo:
```shell
git clone  git@github.com:shankar524/password-generator.git
cd password-generator
go mod download
```
- build
```shell
go build -o ./pwd-gen .
```
or
```shell
make build
```
- check if successfully build
```shell
./pwd-gen version
```
- generate random text
    - without rules(copies text to clipboard)
```shell
./pwd-gen g
```
    - with rules
```shell
./pwd-gen g --numbers=7 --symbols=2 --down=2 --up=3 --copy=false
```
or
```shell
./pwd-gen g -n=7 -s=2 -d=2 -u=3 -c=false
```
- help?
```shell
./pwd-gen help
```

#### Available Flags

| Flag   | Short |type   |Default value |Meaning        |
| -------| ------|-------|--------------|:-------------:|
| symbol | s     |number |0             |number of symbols to generate in return text|
| up     | u     |number |4             | number of upper case letters to generate in return text|
| down   | d     |number |3             | number of down case letters to generate in return text|
| number | n     |number |0             | number of numerical digits to generate in return text|
| copy   | c     |boolean|true          | copy generated text to clipboard.If set false then prints text in console itself|

### Package usage
This package can also be used as common library to generate random text from provided text rule.
`text.TextBuilder` implements `text.Builder` interface which include following two function to help build text rule
- AddRule(TextRule) Builder
- Build() (Generator, error)
Also `text.text` implements `text.Generator` interface which allows us to generate text
- Generate() string

Sample code usage:

```go
  import(
    "github.com/shankar524/password-generator/src/text"
  )
  // create textBuilder
  textBuilder := text.TextBuilder{}
  textBuilder.AddRule(text.TextRule{For: text.LOWERCASE, Length: 6})
  textBuilder.AddRule(text.TextRule{For: text.NUMBER, Length: 2})
  textBuilder.AddRule(text.TextRule{For: text.UPPERCASE, Length: 2})
  textBuilder.AddRule(text.TextRule{For: text.SYMBOL, Length: 2})

  generator, err:= textBuilder.Build()
  if err!= nil {
    fmt.printf("error creating text generator. Error: %s", err.Error())
    return
  }

  fmt.Print(generator.Generate())

  // generate default generator
  defaultBuilder := text.TextBuilder{}
  defaultGenerator, _ := defaultBuilder.Build()
  fmt.Print(defaultGenerator.Generate())
```
