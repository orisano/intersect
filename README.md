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
BenchmarkIntersect/n=100000000,nTest=32,mMax=1000000000/Naive-8   	      10	 107656768 ns/op
BenchmarkIntersect/n=100000000,nTest=32,mMax=1000000000/Galloping-8         	  797413	      1477 ns/op
BenchmarkIntersect/n=100000000,nTest=32,mMax=1000000000/Shotgun1-8          	 1000000	      1109 ns/op
BenchmarkIntersect/n=100000000,nTest=32,mMax=1000000000/Shotgun4-8          	 1671450	       701 ns/op
BenchmarkIntersect/n=10000,nTest=7000,mMax=1000000000/Naive-8               	   21763	     54934 ns/op
BenchmarkIntersect/n=10000,nTest=7000,mMax=1000000000/Galloping-8           	   26590	     44635 ns/op
BenchmarkIntersect/n=10000,nTest=7000,mMax=1000000000/Shotgun1-8            	    8660	    132265 ns/op
BenchmarkIntersect/n=10000,nTest=7000,mMax=1000000000/Shotgun4-8            	    8527	    140649 ns/op
PASS
ok  	github.com/orisano/intersect	37.614s
```

## Author
Nao Yonashiro (@orisano)

## License
Apache 2.0

