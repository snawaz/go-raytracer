
package raytracer

import (
    "math"
)

type HitRecord struct {
    P        Point
    Normal   Vec
    Mat      Material
    T        float64
    IsFrontFace bool
}

type Hittable interface {
    Hit(ray Ray, min, max float64) *HitRecord
}

type Sphere struct {
    Center   Point
    Radius   float64
    Mat      Material
}

func (s * Sphere) Hit(ray Ray, min, max float64) *HitRecord {
    oc := ray.Origin.Substract(&s.Center)
    a := ray.Direction.LengthSquared()
    half_b := oc.Dot(&ray.Direction)
    c := oc.LengthSquared() - s.Radius * s.Radius
    discriminant := half_b * half_b - a * c
    if discriminant <= 0 {
        return nil
    }
    recordFn := func (t float64) *HitRecord {
        point := ray.PointAt(t)
        normal := point.Substract(&s.Center).Scale(1.0/s.Radius)
        front_face := ray.Direction.Dot(normal) < 0
        if front_face == false {
            normal = normal.Negate()
        }
        return &HitRecord {*point, *normal, s.Mat, t, front_face}
    }
    root := math.Sqrt(discriminant)
    t1 := (-half_b - root) / a
    t2 := (-half_b + root) / a
    if t1 < max && t1 > min {
        return recordFn(t1)
    } else if t2 < max && t2 > min {
        return recordFn(t2)
    }
    return nil
}

type HittableList struct {
    Items []Hittable
}

func (hl * HittableList) Hit(ray Ray, min, max float64) *HitRecord {
    current_max := max
    var current_rec *HitRecord
    for _, h := range hl.Items {
        rec := h.Hit(ray, min, current_max)
        if rec != nil {
            current_rec = rec
            current_max = rec.T
        }
    }
    return current_rec
}

func reflect(v, n *Vec) *Vec {
    return v.Substract(n.Scale(2 * v.Dot(n)))
}

func refract(unitRay, normal *Vec, etai_over_etat float64) *Vec {
    cos_theta := unitRay.Negate().Dot(normal)
    r_out_parallel := unitRay.Add(normal.Scale(cos_theta)).Scale(etai_over_etat)
    r_out_perp := normal.Scale(math.Sqrt(1.0 - r_out_parallel.LengthSquared())).Negate()
    return r_out_perp.Add(r_out_parallel)
}

func schlick(cosine, refIdx float64) float64 {
    r0 := math.Pow((1-refIdx) / (1+refIdx), 2)
    return r0 + (1 - r0) * math.Pow(1 - cosine, 5)
}

type Material interface {
    Scatter(ray Ray, record HitRecord) (*Ray, *Vec)
}

type Lambertian struct {
    ColorV Vec
}

type Metal struct {
    ColorV Vec
    Fuzz float64
}

type Dielectric struct {
    RefIdx float64
}

func (m * Lambertian) Scatter(ray Ray, record HitRecord) (*Ray, *Vec) {
    newRay := NewRay(&record.P, record.Normal.Add(SampleUnitVector()))
    return &newRay, &m.ColorV
}

func (m * Metal) Scatter(ray Ray, record HitRecord) (*Ray, *Vec) {
    reflected := reflect(ray.Direction.Unit(), &record.Normal)
    if reflected.Dot(&record.Normal) > 0 {
        direction := reflected.Add(SamplePointInSphere().Scale(math.Min(m.Fuzz, 1.0)))
        newRay := NewRay(&record.P, direction)
        return &newRay, &m.ColorV
    } else {
        return nil, nil
    }
}

func (m * Dielectric) Scatter(ray Ray, record HitRecord) (*Ray, *Vec) {
    etai_over_etat := m.RefIdx
    if record.IsFrontFace {
        etai_over_etat = 1.0 / m.RefIdx
    }
    cos_theta := math.Min(1.0, ray.Direction.Unit().Negate().Dot(&record.Normal))
    sin_theta := math.Sqrt (1.0 - cos_theta * cos_theta)
    reflect_probability := schlick(cos_theta, etai_over_etat)
    var scatter_direction *Vec
    if (etai_over_etat * sin_theta) > 1.0 || reflect_probability > Sample() {
        scatter_direction = reflect(ray.Direction.Unit(), &record.Normal)
    } else {
        scatter_direction = refract(ray.Direction.Unit(), &record.Normal, etai_over_etat)
    }
    newRay := NewRay(&record.P, scatter_direction)
    return &newRay, Ones()
}
