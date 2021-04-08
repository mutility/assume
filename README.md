# assume

The assume package helps clarify your invariant assumptions.

[![CI](https://github.com/mutility/assume/actions/workflows/build.yaml/badge.svg)](https://github.com/mutility/assume/actions/workflows/build.yaml)

## Usage

Use like one of the following:

```go
if assume.True(op()) { phew() }
if assume.False(op()) { eek() }
if err := assume.Ok(op()); err != nil { eek() }
if err := assume.Err(op()); err != nil { whew() }
```

When built normally, this will issue panics when the desired invariant fails. This is useful to highlight bad callers of your code during your tests.

> ⚠️ Warning: this is only intended for testing invariants. If you come to rely on `assume.Ok(err)` to issue a panic, your code may misbehave when built without verification.

## Building

By default, these functions will verify their inputs and return them. If the assumptions are not met, they will panic.

Build or test with `-tags=noverify` to elide the panics, resulting in inlined functions. Typically you should run your tests verification enabled; this will help identify local misuse of your interface. Then consider building your "release" binaries with noverify to remove the verification overhead.
