# Blog Posts
learn about making markdown parser for blog with golang.
using tdd.


## dependances
- install gow (go watcher)
```sh
go install github.com/mitranim/gow@latest
```
- download [tailwindcss cli](https://github.com/tailwindlabs/tailwindcss/releases/latest)
- add to tailwind cli to path

## build
```sh
go build
```

## test
```sh
go test
```

## run in devmode
```
  ## run go server and restart on files change
  gow -e=go,gohtml,css -w .. run main.go
  ## generate tailwindcss classes watch changes
  tailwind -i static/style.css -o dest/style.css --watch
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
