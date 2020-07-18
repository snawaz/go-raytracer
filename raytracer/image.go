
package raytracer

import (
    "fmt"
    "math"
    "os"
    "sync"
)

type Color struct {
    red int
    green int
    blue int
}

type Image struct {
    Width int
    Height int
    Colors []Color
}

func (color * Color) String() string {
	return fmt.Sprintf("%d %d %d", color.red, color.green, color.blue)
}

func (v * Vec) toColor(samplesPerPixel int) Color {
    color := func(f float64) int {
        return int(math.Floor(256.0 * math.Max(0.0, math.Min(0.999, math.Sqrt(f/float64(samplesPerPixel))))))
    }
    return Color { color(v.X), color(v.Y), color(v.Z) }
}

func rayColor(ray Ray, world Hittable, raysPerSample int) *Vec {
    if raysPerSample <= 0 {
        return Zeroes()
    }
    t := 0.5 * (ray.Direction.Unit().Y + 1.0)
    default_color := Ones().Scale(1.0 -t).Add(NewVec(0.5, 0.7, 1.0).Scale(t))
    rec := world.Hit(ray, 0.001, math.MaxFloat64)
    if rec != nil {
        scatteredRay, attenuation := rec.Mat.Scatter(ray, *rec)
       if scatteredRay != nil {
           c := rayColor(*scatteredRay, world, raysPerSample - 1)
           return c.Multiply(attenuation)
       }
       return Zeroes()
   }
   return default_color
}

func createImage(width, height, samplesPerPixel, raysPerSample int, world Hittable) Image {
    colors := make([]Color, width * height)
    camera := NewCamera(NewVec(13, 2, 3), Zeroes(), NewVec(0, 1, 0), 20.0, float64(width)/float64(height), 0.1, 10.0)
    var waitGroup sync.WaitGroup
    waitGroup.Add(height)
    processRows := func(j int) {
        for i := 0; i < width; i++ {
            k := j * width + i
            sampled := Zeroes()
            for k := 0; k < samplesPerPixel; k++ {
                u := (float64(i) + Sample()) / float64(width - 1)
                v := (float64(j) + Sample()) / float64(height - 1)
                sampled = sampled.Add(rayColor(camera.RayAt(u, v), world, raysPerSample))
            }
            colors[k] = sampled.toColor(samplesPerPixel)
        }
        waitGroup.Done()
    }
    for j := 0; j < height; j++ {
        go processRows(j)
    }
    waitGroup.Wait()
    return Image { width, height, colors }
}

func WriteImage(width, height, samplesPerPixel, raysPerSample int, world Hittable, filepath string) {
    image := createImage(width, height, samplesPerPixel, raysPerSample, world)
    file, _ := os.Create(filepath)
    fmt.Fprintln(file, "P3")
    fmt.Fprintf(file, "%d %d\n", image.Width, image.Height)
    fmt.Fprintln(file, "255")
    for j := image.Height -1; j >= 0; j-- {
        for i := 0 ; i < image.Width; i++ {
          k := j * image.Width + i
          fmt.Fprintln(file, image.Colors[k].String())
        }
    }
}
