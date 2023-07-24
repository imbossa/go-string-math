# go-string-math
> go-string-math is a data arithmetic method written by golang. It can be used for numeric string calculation and numerical calculation of very long digits

## Installation
```shell
go get github.com/imbossa/go-string-math
```

## Quickstart
```` go
import (
    math "github.com/imbossa/go-string-math"
    "fmt"
)

func Example() {
    fmt.Println(math.Add(1,"1"))
    fmt.Println(math.Add("-1.021","1.021"))
    fmt.Println(math.Add(1.000001,"-1.0000032321"))
    fmt.Println(math.Add("3232423534535241232.312312312554","974234345343232432322434353463413123123454632243445345.53453232334234"))
    
    fmt.Println(math.Subtract(1,"1"))
    fmt.Println(math.Subtract("-1.021","1.021"))
    fmt.Println(math.Subtract(1.000001,"-1.0000032321"))
    fmt.Println(math.Subtract("3232423534535241232.312312312554","97423434523123454632243445345.53453232334234"))
}
````


## License
This project is licensed under the [MIT License](/LICENSE).