package engine

import (
	"image/color"
	"math"
	"sort"
)

type Scene struct {
	pointLights []PointLight
	camera      Camera
	shapes      []*Shape
}

func NScene(pointLights []PointLight, camera Camera, shapes []*Shape) Scene {
	for _, s := range shapes {
		s.calculateInverseTransformations()
	}
	scene := Scene{
		pointLights: pointLights,
		camera:      camera,
		shapes:      shapes,
	}
	return scene
}

func (scene Scene) GetPixel(x, y int) color.Color {
	ray := scene.camera.rayForPixel(x, y)
	return colorAt(scene.shapes, scene.pointLights, ray)
}

type Intersect struct {
	T     float64
	Shape *Shape
}

type PointLight struct {
	Position  Tuple
	Intensity Color
}

type HitRecord struct {
	T       float64
	shape   *Shape
	point   Tuple
	eyeV    Tuple
	normalV Tuple
	inside  bool
}

func colorAt(shapes []*Shape, lights []PointLight, ray Ray) Color {
	intersects := intersectShapes(shapes, ray)
	if len(intersects) == 0 {
		return Black
	} else {
		hit := intersects[0]
		hitRecord := createHitRecord(hit, ray)
		return shadeHit(lights, hitRecord)
	}

}

func intersectShapes(shapes []*Shape, worldRay Ray) []Intersect {
	intersects := make([]Intersect, 0)
	for _, shape := range shapes {
		intersects = append(intersects, shape.intersect(worldRay)...)
	}
	sort.Slice(intersects, func(i, j int) bool {
		return intersects[i].T < intersects[j].T
	})
	return intersects
}

func createHitRecord(intersect Intersect, ray Ray) HitRecord {
	point := ray.Position(intersect.T)
	normalV := intersect.Shape.normalAt(point)
	eyeV := ray.Direction.Negate()
	inside := false
	if normalV.Dot(eyeV) < 0 {
		inside = true
		normalV = normalV.Negate()
	}
	return HitRecord{
		T:       intersect.T,
		shape:   intersect.Shape,
		point:   point,
		eyeV:    eyeV,
		normalV: normalV,
		inside:  inside,
	}
}

func shadeHit(lights []PointLight, record HitRecord) Color {
	total := Black
	for _, light := range lights {
		total = total.Add(lighting(record.shape.getMaterial(), light, record.point, record.eyeV, record.normalV))
	}
	return total
}

func lighting(material Material, light PointLight, point Tuple, eyeV Tuple, normalV Tuple) Color {
	effectiveColor := material.Color.Multiply(light.Intensity)
	lightV := light.Position.Sub(point).Normalize()
	ambient := effectiveColor.Scale(material.Ambient)
	lightDotNormal := lightV.Dot(normalV)
	diffuse := Black
	specular := Black
	if lightDotNormal >= 0 {
		diffuse = effectiveColor.Scale(material.Diffuse * lightDotNormal)
		reflectV := reflect(lightV.Scale(-1), normalV)
		reflectDotEye := reflectV.Dot(eyeV)
		if reflectDotEye > 0 {
			specular = light.Intensity.Scale(material.Specular * math.Pow(reflectDotEye, material.Shininess))
		}
	}
	return ambient.Add(diffuse).Add(specular)
}

func reflect(in Tuple, normal Tuple) Tuple {
	return in.Sub(normal.Scale(2 * in.Dot(normal)))
}
