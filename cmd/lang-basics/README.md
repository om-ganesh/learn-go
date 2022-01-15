## How to import the public Go Modules

- Initialize the module 
```
$go mod init module-test
```

- Import the module in your go file  
``` 
import "github.com/om-ganesh/goutils" 
```
- Use the desired public method(s)  
``` 
fmt.Println(goutils.Hello()) 
```
- Add the module to your project from github  
(Note, go.mod, go.sum file will be created)  
```
$go get github.com/om-ganesh/goutils 
```
- Run your project
```
$go build 
```