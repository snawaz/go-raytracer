// Copyright 2019 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
    "time"
    "os"
    "strconv"
    "math"
    "strings"

    "github.com/snawaz/goraytracer/raytracer"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("./goraytracer <width>")
        return
    }
    width, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Wrong input: ", err)
        return
    }
    height := int(float64(width) * 9.0/16.0)
    samplesPerPixel := 100
    raysPerSample := 50
    err = os.Mkdir("images", 0700)
    if err != nil && !strings.Contains(err.Error(), "exists") {
        fmt.Println("Mkdir failed: ", err)
        return
    }
    start := time.Now()

    raytracer.WriteImage(width, height, samplesPerPixel, raysPerSample, raytracer.RandomScene(), "images/tmp.ppm")

    elapsed := int(math.Round(float64(time.Now().Sub(start)) / 1000000000.0))
    rate := float64(width*height)/float64(elapsed)
    m := elapsed / 60
    s := elapsed % 60
    filename := fmt.Sprintf("images/%d.%dx%d.%d-%d.%dm-%ds.%d.ppm", int(rate), width, height, samplesPerPixel, raysPerSample, m, s, start.Unix())
    err = os.Rename("images/tmp.ppm", filename)
    if err != nil {
        fmt.Println("Rename failed: ", err)
        return
    }
    fmt.Printf("Elapsed time      :\033[32m %dm %ds\n\033[0m", m, s)
    fmt.Printf("Pixels per second :\033[32m %f\n\033[0m", rate)
    fmt.Printf("Image produced    :\033[32m %s\n\033[0m", filename)
}
