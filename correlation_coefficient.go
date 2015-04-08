package correlation_coefficient

import "math"

func CorrelationCoefficient(x []float64, y []float64)(r float64){
  x_sum := 0.0
  for i := 0; i < len(x); i++ {
    x_sum += float64(x[i])
  }
  mean_x := float64(x_sum) / float64(len(x))

  y_sum :=0.0
  for i := 0; i < len(y); i++ {
    y_sum += float64(y[i])
  }
  mean_y :=float64(y_sum) / float64(len(y))

  cov := 0.0
  for i:=0; i < len(x); i++ {
    cov += (float64(x[i]) - mean_x) * (float64(y[i])  - mean_y)
  }

  var_x := 0.0
  for i:=0; i < len(x); i++ {
    var_x += math.Pow( (float64(x[i]) - mean_x), 2.0)
  }

  var_y := 0.0
  for i:=0; i < len(y); i++ {
    var_y += math.Pow( (float64(y[i]) - mean_y), 2.0)
  }

  //相関係数
  r = cov / math.Sqrt(var_x)
  r /= math.Sqrt(var_y)

  return r
}


