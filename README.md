# intersect
![Test](https://github.com/orisano/intersect/workflows/Test/badge.svg)

A library of intersect to sorted array. inspired by this article. https://lemire.me/blog/2019/01/16/faster-intersections-between-sorted-arrays-with-shotgun/

## Installation
```bash
go get github.com/orisano/intersect
```

## Benchmark
```
goos: darwin
goarch: amd64
pkg: github.com/orisano/intersect
BenchmarkIntersect/Naive-8         	       9	 116047739 ns/op
BenchmarkIntersect/Galloping-8     	  756350	      1669 ns/op
BenchmarkIntersect/Shotgun1-8      	 1000000	      1219 ns/op
BenchmarkIntersect/Shotgun4-8      	 1522401	       789 ns/op
PASS
ok  	github.com/orisano/intersect	34.603s
```

## Author
Nao Yonashiro (@orisano)

## License
Apache 2.0

