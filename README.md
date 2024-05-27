# caller

Wrapping Go errors with a caller name and other helpers

`caller.Wrap(err) error` - wrap error with caller name. Example

```go
package mypkg

import "github.com/spirin/caller"

func myFunc() error {
    err := doSmth()
    if err != nil {
        // "mypkg.myFunc: some err"
        return caller.Wrap(err)
    }
    ...
}

```

`caller.WrapPackage(err) error` - wrap error with caller package name. Example: "http: some err"

```go
package mypkg

import "github.com/spirin/caller"

func myFunc() error {
    err := doSmth()
    if err != nil {
        // "mypkg: some err"
        return caller.WrapPackage(err)
    }
    ...
}

```

`caller.Name() string` - get caller name, e.g., _"caller.testNameObj.Fn"_

`caller.NameFull() string` - get caller full name, e.g., _"testNameFullObj.Fn @ github.com/spirin/caller"_

`caller.Package() string` - get caller package name, e.g., _"caller"_
