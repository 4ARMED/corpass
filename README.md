# Corpass

> Corporation name password generator

Corpass is a simple password generator for making lists of potential passwords based on a company name. There's probably lots of other tools with more generic applicability but this is mine based on years of finding people do similar things when it comes to setting organisation passwords.

## Installation

The easiest way is to use Go, which obviously requires you to have Go installed! It needs version 1.21 and up.

```bash
$ go install github.com/4armed/corpass@latest
```

## Usage

```bash
Usage: corpass [string]

  -generators string
        comma separated list of generators to use (default "upperlower,leet,numerics,punctuation")
  -verbose
        enable verbose logging
  -version
        print version and exit
```

In its simplest incarnation, just run `corpass` and pass a single argument which is the name (or basically any string) you want to generate from. Generated values are output to stdout. If you specify `-verbose` this goes to stderr so you can redirect as needed. For example:

```bash
$ corpass -verbose foo 1> /dev/null
time=2023-08-29T20:42:14.717+01:00 level=DEBUG msg="using generators" names="[upperlower leet numerics punctuation]"
```

There are a few generators included as standard and it should be easiest to implement more by just adding them to the `generators` package and satisfying the `generators.Generator` interface.

By default it will run with all generators but you can override this to run specific ones.

### upperlower

This generator iterates over the provided string and goes through every permutation of upper and lower case versions of each character. For example:

```bash
$ corpass -generators upperlower foo
FOO
FOo
FoO
Foo
fOO
fOo
foO
foo
```

### leet

This generate converts common characters to their l33tspeak equivalent. For example, e to 3.

```bash
$ corpass -generators leet foo
f00
foo
```

### numerics

Adds numbers 0 to 9 (by default) to the beginning and end of words.

```bash
$ corpass -generators numerics foo
0foo
1foo
2foo
3foo
4foo
5foo
6foo
7foo
8foo
9foo
foo
foo0
foo1
foo2
foo3
foo4
foo5
foo6
foo7
foo8
foo9
```

### punctuation

Adds punctation "special" characters to the beginning and end of values.

```bash
$ corpass -generators punctuation foo
!foo
"foo
#foo
$foo
%foo
&foo
'foo
(foo
)foo
*foo
+foo
,foo
-foo
.foo
/foo
:foo
;foo
<foo
=foo
>foo
?foo
@foo
[foo
\foo
]foo
^foo
_foo
`foo
foo
foo!
foo"
foo#
foo$
foo%
foo&
foo'
foo(
foo)
foo*
foo+
foo,
foo-
foo.
foo/
foo:
foo;
foo<
foo=
foo>
foo?
foo@
foo[
foo\
foo]
foo^
foo_
foo`
foo{
foo|
foo}
foo~
{foo
|foo
}foo
~foo
```

## Combining Generators

Generators are additive and run in the order provided. This means you can combine them. Each generator's output is printed first before passed to the next one so you get all combinations.

```bash
$ corpass -generators upperlower,leet foo
F00
F0O
FO0
FOO
FOo
FoO
Foo
f00
f0O
fO0
fOO
fOo
foO
foo
```

## Contributing

Feel free to raise issues and submit pull requests for any changes. They will be much appreciated!