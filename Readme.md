
[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]
[![GoDoc][doc-img]][doc-url]

### Install

```bash
go get github.com/pkg4go/convert
```

### Example

```go
import "github.com/pkg4go/convert"

...

s := "10086"
i, _ := convert.Int(s)
i32, _ := convert.Int32(s)
i64, _ := convert.Int64(s)
f32, _ := convert.Float32(s)
f64, _ := convert.Float64(s)

...

```

### License
MIT

[doc-img]: http://img.shields.io/badge/GoDoc-reference-green.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/pkg4go/convert
[travis-img]: https://img.shields.io/travis/pkg4go/convert.svg?style=flat-square
[travis-url]: https://travis-ci.org/pkg4go/convert
[license-img]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
