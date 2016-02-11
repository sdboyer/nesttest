## Nesting Testing

This is a simple test repo to explore and demonstrate how nesting a repository inside itself can be used to address a non-encapsulating vendor directory - a situation that can arise when building a proper language package manager.

Our basic file layout here looks like this:

```
quux.go
bar/
   bar.go
foo/
    baz/
        baz.go
    foo.go
    other.go
    vendor/
        github.com/sdboyer/nesttest/(RECURSION)
```

`foo` is the main package, and it calls the single function in each of the other packages (`bar`, `baz`, `quux`), then quits. When run without any vendor directory, this is the output:

```
Hello from foo/main in main area
Hello from other foo file under main
Hello from bar in main area
Hello from baz in main area
Hello from quux in main area
```

When run with itself fully replicated under the vendor tree, this is the output:

```
Hello from foo/main in main area
Hello from other foo file under main
Hello from bar in vendor area
Hello from baz in vendor area
Hello from quux in vendor area
```

And when run with only the packages not shadowed by the vendor tree (`bar` and `quux`), this is the output:

```
Hello from foo/main in main area
Hello from other foo file under main
Hello from bar in vendor area
Hello from baz in main area
Hello from quux in vendor area
```

All of this is what we’d want and expect.
