package main

import (
    "fmt"
    "math"
)

const EPSYLON = 0.0001

func sqrt(x float64) (z float64) {
    i := 0
    z = 1.0
    for diff := math.MaxFloat64; diff > EPSYLON; {
        old_z := z
        z -= (z*z - x) / (2*z)
        diff = math.Abs(old_z - z)

        fmt.Printf("iteration %d, z is %f and diff %f\n", i, z, diff)
        i++
    }

    return
}

func main() {
    val := 10789347598437985
    fmt.Printf("%f\n", sqrt(float64(val)))
    fmt.Printf("real sqrt %f\n", math.Sqrt(float64(val)))
}

