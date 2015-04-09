# Install

```command
$ go get github.com/shigekid/go_correlation_coefficient
```

# Example

```sample.go
package main

import (
  "fmt"
  r "github.com/shigekid/go_correlation_coefficient"
)

func main(){
  slice_1 := []float64{0,1,2,3}
  slice_2 := []float64{0,1,2,2}
  fmt.Println(r.CorrelationCoefficient(slice_1, slice_2))
}

```
```command
$ go run sample.go
0.9438798074485389
```

