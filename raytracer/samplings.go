
package raytracer

import (
    "math"
    "math/rand"
)

func Sample() float64 {
    return rand.Float64()
}

func SampleBetween(lo, hi float64) float64 {
    return lo + (hi - lo) * rand.Float64()
}

func SampleVecBetween(lo, hi float64) Vec {
    return Vec { SampleBetween(lo, hi), SampleBetween(lo, hi), SampleBetween(lo, hi) }
}

func SampleVec() Vec {
    return Vec {rand.Float64(), rand.Float64(), rand.Float64()}
}

func SampleUnitVector() *Vec {
    a := SampleBetween(0, 2 * math.Pi)
    z := SampleBetween(-1, 1)
    r := math.Sqrt(1 - z * z)
    return &Vec { r * math.Cos(a), r * math.Sin(a), z}
}

func SamplePoint() *Point {
    return NewVec(Sample(), Sample(), Sample())
}

func SamplePointInSphere() *Vec {
    for {
        p := &Vec {SampleBetween(-1, 1), SampleBetween(-1, 1), SampleBetween(-1, 1) }
        if p.LengthSquared() < 1 {
            return p
        }
    }
}
func SampleVecInUnitDisk() *Vec {
    for {
        x := SampleBetween(-1, 1)
        y := SampleBetween(-1, 1)
        v := Vec {x, y, 0} 
            if v.LengthSquared() < 1.0 {
            return &v
        }
    }
}
