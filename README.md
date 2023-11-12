[![codecov](https://codecov.io/gh/jmmills/go-fopt/graph/badge.svg?token=K3GFRHIUQ7)](https://codecov.io/gh/jmmills/go-fopt)

# fopt generates a functional options pattern for Go

```bash
./fopt
Error: accepts 2 arg(s), received 0
Usage:
  fopt <package> <type name> [flags]

Flags:
  -h, --help                 help for fopt
      --no-package-options   will disable the generation of package level option functions
      --only-builder         will genearte builder methods for functional options without package level option functions
  -o, --option string        defines the name of the option type (default "Option")
  -p, --prefix string        defines a prefix to prepend to option function names
      --with-builder         will generate builder methods on the functional options type
  -w, --write string         defines file to write rather than stdout
```

# Define Options

To define a struct field to generate an option for simply add an `fopt` struct tag with the name you want for the option.

```go
type MyFuncOptions struct {
  MagicEnabled `fopt:WithMagic`
}
```

Then generate the options source and write them to `options.go` for this structure

```bash
$ fopt . MyFuncOptions -w options.go
```

This will generate the functional options boiler plate code and a functional option named `WithMagic`

```go
type Option func(s *Config)

type Options []Option

func (opts Options) Apply(s *Config) Config {
        for _, opt := range opts {
                opt(s)
        }
        return *s
}

// Name will set the Name option for Config.
func WithMagic(value bool) Option {
        return func(s *MyFuncOptions) {
                s.MagicEnabled = value
        }
}
```

Once this code is generated you can then consume options in your code:

```go
func MyMagicFunc(opts ...Option) {
  cfg := Options(opts).Apply(new(MyFuncOptions))
  fmt.Println(cfg.MagicEnabled)
}
```

Which allows this function to accept this options like:

```go
MyMagicFunc(WithMagic(true)) // prints "true"
```

# Builder pattern

In some cases you may want to reduce the number of imports a user of your library needs to make in order to 
utilize functions or methods that consume functional options, for this use case we supply a builder pattern generator.

To use generate the builder pattern for functional options supply the `--with-builder` or `--only-builder` flags (the `--only-builder` flag will not generate the package level option functions):

```bash
$ fopt . MyFuncOptions --only-builder -w options.go
```

This will generate a method on the options slice similiar to this:

```go
func (opts *Options) WithMagic(value bool) Options {
        *opts = append(*opts, Option(func(s *MyFuncOptions) {
                s.MagicEnabled = value
        }))
        return *opts
}
```

This allows for a pattern in your types that make option building part of an interface:

```go

type MyMethodConfig struct {
  Magic bool `fopt:"WithMagic"`
}

func (s Struct) MyMethodOptions() (o Options) {
  return 
}
```

Which then couples the options of the method that accepts them

```go

o.MyMethod(
  value, 
  o.MyMethodOptions().WithMagic(true)...
)
```

## Multiple methods that accept functional options

In some cases you may have multiple methods in a type that accept functional options, but wish to still use the builder pattern.
To reduce clutter and avoid naming conflicts use the `--option` flag with the `--only-builder` flag to generate an options type for each method.

Given these configuration structs:
```go
type FooConfig struct {
  Tags []string `fopt:"WithTags"`
}

type BarConfig struct {
  Debug bool`fopt:"SetDebug"` 
}
```

You can generate option types for each:

```bash 
$ fopt . MyFuncOptions --option FooOption --only-builder -w foo_options.go
$ fopt . MyFuncOptions --option Baroption --only-builder -w bar_options.go
```

No each method can accept it's own options type:

```go
func (s Struct) Foo(opts ...FooOption) {
  var cfg FooConfig
  Options(opts).Apply(&cfg)
}

func (s Struct) Bar(opts ...BarOption) {
  var cfg BarConfig
  Options(opts).Apply(&cfg)
}
```

To make these easier we can now (optionally) consolidate our optional builders into a single type with methods:

```go
type Options struct{}

func (Options) Foo() (o FooOptions) { return }
func (Options) Bar() (o BarOptions) { return }

func (s Struct) Options() (o Options) { return }
```

Which allows us a nice looking call pattern:

```go
var s Struct

s.Foo(s.Options().Foo().WithTags("tag_a")...)
s.Bar(s.Options().Bar().SetDebug(true)...)
```
