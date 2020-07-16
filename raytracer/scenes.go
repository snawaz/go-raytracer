
package raytracer

/*
    where
        allObjects = [groundObject] ++ randomObjects ++ fixedObjects
        groundObject = Sphere (vec 0 (-1000) 0) 1000 (Material (Lambertian (toColor $ from 0.5)))
        randomObjects = fst $ foldl' mkObject ([], mkStdGen seed) $ (,) <$> [-11..11] <*> [-11..11]
        mkObject (objects, g) (a, b) = (fromMaybe objects (fmap (\o-> o:objects) object), g2)
            where
                ((chooseMat, x, z), g1) = sampleFraction3 g
                center = vec (a + 0.9 * x) 0.2 (b + 0.9 * z)
                (object, g2) = if len (center - vec 4 0.2 0) > 0.9 then createRandomObject g1 else (Nothing, g1)
                createRandomObject g3 = (Just obj, g5)
                    where
                        ((p1, p2), g4) = samplePoint2 g3
                        (d1, g5) = sampleFraction g4
                        mat = if chooseMat < 0.8
                                 then Material $ Lambertian (toColor $ p1 * p2)
                                 else if chooseMat < 0.95
                                    then Material $ Metal (fmap (\i -> 0.5 + 0.5 * i) $ toColor p1) (0.5 * d1)
                                    else Material $ Dielectric 1.5
                        obj = Sphere center 0.2 mat
        fixedObjects =  [
                (Sphere (vec 0 1 0) 1.0 (Material (Dielectric 1.5))),
                (Sphere (vec (-4) 1 0) 1.0 (Material (Lambertian (toColor $ vec 0.4 0.2 0.1)))),
                (Sphere (vec 4 1 0) 1.0 (Material (Metal (toColor $ vec 0.7 0.6 0.5) 0.0)))
            ]

*/

func RandomScene() Hittable {
    fixed_point := NewVec(4, 0.2, 0)
    var objects []Hittable
    objects = append(objects, &Sphere{*NewVec(0, -1000, 0), 1000, &Lambertian{*VecFrom(0.5)}})
    for a := -11; a <= 11; a++ {
        for b := -11 ; b <= 11; b++ {
            x := Sample()
            z := Sample()
            center := NewVec(float64(a) + 0.9 * x, 0.2, float64(b) + 0.9 * z)
            if center.Substract(fixed_point).Length() > 0.9 {
                p1 := SamplePoint()
                p2 := SamplePoint()
                d1 := Sample()
                chooseMat := Sample()
                var mat Material
                if chooseMat < 0.8 {
                    p := p1.Multiply(p2)
                    mat = &Lambertian{*p}
                } else if chooseMat < 0.95 {
                    p := p1.Scale(0.5).Shift(0.5)
                    mat = &Metal{*p, 0.5 * d1}
                } else {
                    mat = &Dielectric { 1.5 }
                }
                objects = append(objects, &Sphere{*center, 0.2, mat})
            }
        }
    }
    objects = append(objects, &Sphere{*NewVec(0, 1, 0), 1.0, &Dielectric{1.5}})
    objects = append(objects, &Sphere{*NewVec(-4, 1, 0), 1.0, &Lambertian{*NewVec(0.4, 0.2, 0.1)}})
    objects = append(objects, &Sphere{*NewVec(4, 1, 0), 1.0, &Metal{*NewVec(0.7, 0.6, 0.5), 0.0}})
    return &HittableList { objects }
}

