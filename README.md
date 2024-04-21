# Blog Posts
learn about making markdown parser for blog with golang.
using tdd.

## build
```sh
go build
```

## test
```sh
go test
```


## Performance

Old Test benchmark
```
pkg: alnoor/blogposts/renderer
BenchmarkRenderer-4   	   7027	   226383 ns/op
```

New Test benchmark
```
pkg: alnoor/blogposts/renderer
BenchmarkRenderer-4   	  69820	    17487 ns/op
```

New Algorithm speed calc = `226383 / 17487 ≈ 12.94578830`
so new algorithm faster 12 times than the old algorithm
