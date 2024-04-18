# AWS::IAM::Role Policy resource check


## Installation process
To install, follow undermentioned steps:  
- If you haven't already, initialize go module
```console
> go mod init <mod_name>
```
- In the root directory (the one containing `go.mod`) open console
- Type command:
```console
> go get github.com/grysj/remitly-assignment/...
```
If a `go.sum` file appeared, you are ready to go (pun intended)!

## Example usage
Implemented method `Check` from subpackage `check` takes param of type `interface{}` and returns tuple consisting of type `bool` and `error`
```go
res, err := check.Check(param)
```
In case of parameter of type `[]byte` method `Check` interprets parameter value as JSON. Type `string` however interprets parameter value as `filePath` to JSON.
If a parameter `param` has a resource field containing "*" it will return `false` (`true` in any other case assuming JSON is correct)

## Tests
To run tets, type in console:
```console
> go test run github.com/grysj/remitly-assignment/<method>
```
`method` can be one of 3 things:
- check - will run tests connected with correct output of a 'Check' method
- read - will run tests connected with saving JSON to a parameter
- parse - will run tets connected with parsing JSON "Resource" field

Example:
```console
> go test run github.com/grysj/remitly-assignment/check
```
will run tests from subpackage `check`
