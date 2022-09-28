<!--
Please answer these questions before submitting your issue. Thanks!
-->

### What version of Go are you using (`go version`)?

<pre>
$ go version
local: go version go1.19.1 linux/amd64
</pre>

Although most test were run using gotip
<pre>
$ gotip version
go version devel go1.20-53773a5 Wed Sep 28 11:50:58 2022 +0000 linux/amd64
</pre>

### Does this issue reproduce with the latest release?
Yes


### What operating system and processor architecture are you using (`go env`)?

<details><summary><code>go env</code> Output</summary><br><pre>
$ go env
GO111MODULE=""
GOARCH="amd64"
GOBIN="/home/falves/go/bin"
GOCACHE="/home/falves/.cache/go-build"
GOENV="/home/falves/.config/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/falves/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/falves/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.19.1"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
GOWORK=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build1430203757=/tmp/go-build -gno-record-gcc-switches"
</pre></details>

### What did you do?

Given the introduction of generics and the improvements of the functionality in Go, it's probably time to start accepting distinct numeric Type arguments to math functions. The evaluation of whether the implementation of generics is beneficial to a math function should probably be done on a per-function basis.

This investigation was specific to the [Abs](https://github.com/golang/go/blob/master/src/math/abs.go) function of the math package. This functions receives a float64 data type as argument and returns the its absolute number also as a float64.
I saw that @randall77 and others worked on optimizations for the `Abs` function that rely on IEE 754 binary representation of floats. As such, the `Abs` function currently only supports float64 type as an argument.

The proposal was to use generics to take in any of the following basic number data types and object with their underlying types (~):
- ~int,
- ~int8,
- ~int16,
- ~int32,
- ~int64,
- ~float32,
- ~float64

The underlying methodology for the function `Abs` would not be modified.

The advantage of this proposal is that it makes it simpler for the user to interact with it. As the user will no longer be required to include an explicit conversion to and from float64 (in the cases where the original number is not a float64).

The possible disadvantage of this proposal is that it might degrade the improvements previously introduced to the function.

### Analysis and Methodology

Three scenarios were benchmarked:

1. Using current `Abs` implementation that uses the float64 data type as the parameter and return value.
2. Implementation of `Abs` function using generics as parameter and float64 as return value.
3. Implementation of `Abs` function using generics as parameter and return value.

The tests were run for the following inputs:
1. float64
2. int
3. int64
4. struct float64
5. struct int64
6. Type float64
7. Type int64

Assumption: The intended data type should be provided as an argument and the return of the `Abs` function.
Consequence: type conversions are only necessary in some cases:
i.e.:
When running test variant 1.1 - current `Abs` implementation with float64. No type conversion is needed:
```
var number float64 = 5
_ = Abs(number)
```
When running test variant 1.2 - current `Abs` implementation with int. Type conversion is needed for the argument and the return value:
```
var number int = 5
_ = int(Abs(float64(number)))
```
When running test variant 2.3 - implementation using generics for function parameter and return, tested with float64. No type conversion is needed:
```
var number int64 = 5
_ = GenericAbsT(number)
```

The type conversions are in place to perform an end-to-end comparison.  ## EDIT this?
Each of the benchmark test variants was run 100 times. Each benchmark test ran 1000000000 executions.
Benchmark tests were done using the gotip version mentioned previously.

Benchmark tests are aggreated using benchstat and analysed manually.
A bash script automates the testing and outputs the final result in `<type>Benchstat.txt` files under the `benchmark_result_<test_iteration>` directory.

Code and script reproduce the test can be found in: ##ADD GH LINK

### Example of test Process
```
$ gotip version
go version devel go1.20-53773a5 Wed Sep 28 11:50:58 2022 +0000 linux/amd64

$ gotip test -bench BenchmarkAbs -cpu 1 -benchmem -benchtime 1000000000x -count 100 > BenchmarkAbs.txt
$ gotip test -bench BenchmarkGenericAbsReturnFloat -cpu 1 -benchmem -benchtime 1000000000x -count 100 > BenchmarkGenericAbsReturnFloat.txt
$ gotip test -bench BenchmarkGenericAbsReturnT -cpu 1 -benchmem -benchtime 1000000000x -count 100 > BenchmarkGenericAbsReturnT.txt
...
$ benchstat BenchmarkAbs.txt BenchmarkGenericAbsReturnFloat.txt BenchmarkGenericAbsReturnT.txt ...
```
### Results 
The benchstat results have been aggregated further and edited to facilitate analysis

```
name \ time/op         float64      int          int64        structFloat64   structInt64   TypeFloat64   TypeInt64
Abs                    0.24ns ±7%   0.24ns ±8%   0.24ns ±9%   0.24ns ±6%      0.24ns ±7%    0.24ns ±17%   0.24ns ±9%
GenericAbsReturnFloat  0.24ns ±7%   0.24ns ±4%   0.24ns ±8%   0.23ns ±3%      0.24ns ±8%    0.25ns ±17%   0.24ns ±7%
GenericAbsReturnT      0.24ns ±6%   0.24ns ±6%   0.23ns ±2%   0.24ns ±7%      0.23ns ±3%    0.24ns ±6%    0.24ns ±8%

name \ bytesAlloc/op   float64      int          int64        structFloat64   structInt64   TypeFloat64   TypeInt64
Abs                    0.00B        0.00B        0.00B        0.00B           0.00B         0.00B         0.00B
GenericAbsReturnFloat  0.00B        0.00B        0.00B        0.00B           0.00B         0.00B         0.00B
GenericAbsReturnT      0.00B        0.00B        0.00B        0.00B           0.00B         0.00B         0.00B

name \ allocs/op       float64      int          int64        structFloat64   structInt64   TypeFloat64   TypeInt64
Abs                    0.00         0.00         0.00         0.00            0.00          0.00          0.00
GenericAbsReturnFloat  0.00         0.00         0.00         0.00            0.00          0.00          0.00
GenericAbsReturnT      0.00         0.00         0.00         0.00            0.00          0.00          0.00
```

### Conclusion

The results negate the existance of degredation in terms of performance in average execution time, the average number of bytes allocated per operation and the number of allocations per operation.

The recommendation from this analysis is to **apply generics with the proposed contraints to the `math.Abs` function as a mean to improve user experience and to reduce the need for type conversion**.

Please let me know if you're ok with this proposal, if you require any further testing or have any suggestions.