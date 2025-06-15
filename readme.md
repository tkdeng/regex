# Go Regex

A High Performance Regex Package That Uses A Cache.

After calling a regex, the compiled output gets cached to improve performance.

This package uses the core go regexp package under the hood, so its free from external dependencies.

## Installation

```shell
go get github.com/tkdeng/regex
```

## Usage

```go
import (
  "github.com/tkdeng/regex"
)

func main(){
  // pre compile a regex into the cache
  regex.Comp(`re`)

  // compile a regex and safely escape user input
  regex.Comp(`re %1`, `this will be escaped .*`); // output: this will be escaped \.\*
  regex.Comp(`re %1`, `hello \n world`); // output: hello \\n world (note: the \ was escaped, and the n is literal)

  // use %n to reference a param
  // use %{n} for param indexes with more than 1 digit
  regex.Comp(`re %1 and %2 ... %{12}`, `param 1`, `param 2` ..., `param 12`);

  // return an error instead of panic on failed compile
  reg, err := regex.CompTry(`re`)


  // manually escape a string
  // note: the compile methods params are automatically escaped
  regex.Escape(`(.*)? \$ \\$ \\\$ regex hack failed`)

  // determine if a regex is valid, and can be compiled by this module
  regex.IsValid(`re`)

  // determine if a regex is valid, and can be compiled by the builtin RE2 Regexp module
  regex.IsValidRE2(`re`)

  // run a replace function (most advanced feature)
  regex.Comp(`(?flags)re(capture group)`).RepFunc(myByteArray, func(b func(int) []byte) []byte {
    b(0) // get the string
    b(1) // get the first capture group

    return []byte("")

    // returning nil will stop the loop early
    return nil
  })

  // replace with a string
  regex.Comp(`re (capture)`).Rep(myByteArray, []byte("test $1"))

  // replace with a string literal
  regex.Comp(`re`).RepLit(myByteArray, []byte("all capture groups ignored (ie: $1)"))


  // return a bool if a regex matches a byte array
  regex.Comp(`re`).Match(myByteArray)

  // split a byte array in a similar way to JavaScript
  regex.Comp(`re|(keep this and split like in JavaScript)`).Split(myByteArray) // [][]byte
  regex.Comp(`re|(keep this and split)`).SplitStr(myByteArray) // []string

  // a regex string is modified before compiling, to add a few other features
  `use \' in place of ` + "`" + ` to make things easier`
  `(?#This is a comment in regex)`

  // direct access to compiled *regexp.Regexp
  regex.Comp("re").RE


  // another helpful function
  // this method makes it easier to return results to a regex function
  regex.JoinBytes("string", []byte("byte array"), 10, 'c', b(2))

  // the above method can be used in place of this one
  append(append(append(append([]byte("string"), []byte("byte array")...), []byte(strconv.Itoa(10))...), 'c'), b(2)...)
}
```
