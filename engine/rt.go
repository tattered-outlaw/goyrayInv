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
	T         float64
	shape     *Shape
	point     Tuple
	overPoint Tuple
	eyeV      Tuple
	normalV   Tuple
	inside    bool
}

func colorAt(shapes []*Shape, lights []PointLight, ray Ray) Color {
	intersects := intersectShapes(shapes, ray)
	if len(intersects) == 0 {
		return Black
	} else {
		hitIndex := hitIndex(intersects)
		if hitIndex == -1 {
			return Black
		} else {
			hit := intersects[hitIndex]
			hitRecord := createHitRecord(hit, ray)
			return shadeHit(lights, shapes, hitRecord)
		}
	}

}

func hitIndex(intersects []Intersect) int {
	for i, intersect := range intersects {
		if intersect.T > 0 {
			return i
		}
	}
	return -1
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
	overPoint := point.Add(normalV.Scale(EPSILON))
	return HitRecord{
		T:         intersect.T,
		shape:     intersect.Shape,
		point:     point,
		overPoint: overPoint,
		eyeV:      eyeV,
		normalV:   normalV,
		inside:    inside,
	}
}

func shadeHit(lights []PointLight, shapes []*Shape, record HitRecord) Color {
	total := Black
	for _, light := range lights {
		shadowed := isShadowed(light, shapes, record.overPoint)
		total = total.Add(lighting(record.shape.getMaterial(), light, record.point, record.eyeV, record.normalV, shadowed))
	}
	return total
}

func lighting(material Material, light PointLight, point Tuple, eyeV Tuple, normalV Tuple, shadowed bool) Color {
	effectiveColor := material.Color.Multiply(light.Intensity)
	lightV := light.Position.Sub(point).Normalize()
	ambient := effectiveColor.Scale(material.Ambient)
	lightDotNormal := lightV.Dot(normalV)
	diffuse := Black
	specular := Black
	if !shadowed && lightDotNormal >= 0 {
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

func isShadowed(light PointLight, shapes []*Shape, point Tuple) bool {
	v := light.Position.Sub(point)
	distance := v.Magnitude()
	direction := v.Normalize()
	ray := Ray{point, direction}
	intersections := intersectShapes(shapes, ray)
	return hitIndex(intersections) >= 0 && intersections[0].T < distance
}
