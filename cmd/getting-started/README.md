# Unit0: Getting Started

### Installing Go
Just use the [Download and install](https://go.dev/doc/install) steps.

### Checking Version
```
$go version  
```

### Building Go
```
$go build main.go  
```

### Reducing the size of the binary
```
$go build -ldflags "-w -s" -o reduced_binary.exe main.go 
```

### Finding Help ###
```
$go help 
```

### Environment 
```
$go env 
```
### Cross-Compile on other OS  
- Building binary for Unix  
```
$goos=linux goarch=amd64 go build -o linux_binary main.go 
```

- Building binary for Mac  
```
$goos=darwin goarch=amd64 go build -o mac_binary main.go 
```