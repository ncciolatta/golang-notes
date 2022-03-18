## Importing funcs
In Go if you wish to import functions, it needs to start with an Uppercase letter to sinalize.

This one is right
```go
func Writetest(){
    .....
}
```
The lowercase ones are only used locally, in the same files they were made.

## Funcs

- You can return more than one value in a func
- You can ignore a return with the usade of underline (_)

## Types
- Implicit
```go
variablename := "badabing badaboom"
```
- strong typed
```go
var variable_str string = "lmao"
```

- Theres no char type in go
- Errors have the type error, default value is nil

## Structs

The most similar thing to classes in Go
```go
type structure1 struct{
    example string
    ...
}
```