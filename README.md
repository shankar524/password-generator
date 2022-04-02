<!-- [![Build Status](https://github.com/shankar524/password-generator/workflows/test%20and%20build/badge.svg)](https://github.com/shankar524/password-generator/actions?workflow=test%20and%20build) -->
[![Coverage Status](https://coveralls.io/repos/github/shankar524/password-generator/badge.svg)](https://coveralls.io/github/shankar524/password-generator)
[![Go Report Card](https://goreportcard.com/badge/github.com/jandelgado/golang-ci-template-github-actions)](https://goreportcard.com/report/github.com/jandelgado/golang-ci-template-github-actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
# password-generator
`password-generator` is command line interface app to generate random password. This app can be used to generate any random text from provided rules.

### usage

- Clone repo:
```
git clone  git@github.com:shankar524/password-generator.git
cd password-generator
go mod download
```
- build 
```
go build -o ./pwd-gen .
```
-check if successfully build
```
./pwd-gen version
```
-generate random text
    - without rules(copies text to clipboard)
```
./pwd-gen g
```
    - with rules
```terminal
./pwd-gen g --numbers=7 --symbols=2 --down=2 --up=3 --copy=false
```
or
```terminal
./pwd-gen g -n=7 -s=2 -d=2 -u=3 -c=false
```
-help?
```
./pwd-gen help
```

#### Available Flags

| Flag   | Short |Type   |Default value |Meaning|
| -------| ------|-------|--------------|:------------------------------------------------------:|
| symbols| s     |number |0             |number of symbols to generate in return text| 
| up     | u     |number |4             | number of upper case letters to generate in return text|
| down   | d     |number |3             | number of down case letters to generate in return text|
| numbers| n     |number |0             | number of numerical digits to generate in return text|
| copy   | c     |boolean|true          | copy generated text to clipboard.If set false then prints text in console itself|
