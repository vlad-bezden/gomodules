# Using modules inside of the project

An example on how to use internal module from the package

Module can be references from the same project, and it doesn't have to have path to the GitHub. Three things are required:
1. Module must have go.mod file.
```
module stringutils
```
2. In the root directory go.mod must specify imported module
3. Since there is not path to the GitHub and Go doesn't know where to get this module from it must have 'replace' statement pointing to physical location of the module

```
module modules

require modules/stringutils v1.0.0
replace modules/stringutils => ./stringutils
```

## Usage

### Run without compilation
`go run main.go`

### Compile and create executable
 `go build -o modules`.

### Execute
`./modules`
