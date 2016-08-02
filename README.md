# go-problems
Basic mathematical and real world problems resolved in go

### clone-url.go
A basic implementation for a webpage cloner, that saves the clone on a file and serve then in a webserver
#### Usage
```bash
go run clone-url.go -u "http://url.dev/clone" -p 8080
```

### collatz-conjecture.go
A basic implementation for the [collatz conjecture](https://en.wikipedia.org/wiki/Collatz_conjecture)
#### Usage
```bash
go run collatz-conjecture.go -n 5
```
Change the `-n` value to an integer. You can also pass `-r` for recursive execution until the program reaches 1

### fibonacci.go
A basic implementation for the [fibonacci numbers](https://en.wikipedia.org/wiki/Fibonacci_number)
#### Usage
```bash
go run fibonacci.go -n 12
```
Change the `-n` value to an integer representing how many fibonacci numbers you need

### fizz-buzz.go
A basic implementation for the [fizz buzz](https://en.wikipedia.org/wiki/Fizz_buzz).
#### Usage
```bash
go run fizz-buzz.go -n 15
```
Change the `-n` value to an integer representing how many fiz buzz you need

### gauss-legendre-pi.go
Estimate the value of pi based on [Gauss-Legendre method](https://en.wikipedia.org/wiki/Gauss%E2%80%93Legendre_method).
#### Usage
```bash
go run gauss-legendre-pi.go
```

### get-webpage-size.go
Estimate the html webpage size
#### Usage
```bash
go run get-webpage-size.go
```

### get-webpage-size-concurrently.go
Estimate the html webpage size using goroutines and channels
#### Usage
```bash
go run get-webpage-size-concurrently.go
```

### monte-carlo-pi.go
Estimate the value of pi based on [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method).
#### Usage
```bash
go run monte-carlo-pi.go
```

### reverse-string.go
Reverse a string
#### Usage
```bash
go run reverse-string.go -s "Hello World!"
```

### server.go
A basic server in go
#### Usage
```bash
go run server.go
```
