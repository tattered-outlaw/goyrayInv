package engine

import (
	"image/color"
	"math"
	"sort"
	"sync"
)

const maxIntersections = 32

type Scene struct {
	pointLights       []PointLight
	camera            Camera
	shapes            []*Shape
	intersectionsPool *sync.Pool
	rayPool           *sync.Pool
	tuplePool         *sync.Pool
}

func NScene(pointLights []PointLight, camera Camera) *Scene {
	intersectionsPool := &sync.Pool{
		New: func() interface{} {
			//fmt.Println("new intersections")
			return &Intersections{Buf: make([]Intersect, maxIntersections)}
		},
	}
	rayPool := &sync.Pool{
		New: func() interface{} {
			return &Ray{&Tuple{}, &Tuple{}}
		},
	}
	tuplePool := &sync.Pool{
		New: func() interface{} {
			return &Tuple{}
		},
	}
	scene := Scene{
		pointLights:       pointLights,
		camera:            camera,
		shapes:            make([]*Shape, 0),
		intersectionsPool: intersectionsPool,
		rayPool:           rayPool,
		tuplePool:         tuplePool,
	}
	return &scene
}

func (scene *Scene) AddShape(shape Shape) {
	shape = shape.calculateInverseTransformations()
	scene.shapes = append(scene.shapes, &shape)
}

func (scene *Scene) GetPixel(x, y int) color.Color {
	ray := scene.camera.rayForPixel(x, y)
	return colorAt(scene, &ray, scene.intersectionsPool)
}

type Intersect struct {
	T     float64
	Shape *Shape
}

type PointLight struct {
	Position  Tuple
	Intensity Color
}

type Intersections struct {
	Buf []Intersect
	Len int // write index
}

func (i *Intersections) Add(t float64, shape *Shape) {
	if i.Len == maxIntersections {
		panic("too many intersections")
	}
	i.Buf[i.Len].T = t
	i.Buf[i.Len].Shape = shape
	i.Len++
}

func (i *Intersections) Return(pool *sync.Pool) {
	i.Len = 0
	pool.Put(i)
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

func colorAt(scene *Scene, ray *Ray, intersectionsPool *sync.Pool) Color {
	intersections := intersectionsPool.Get().(*Intersections)
	defer intersections.Return(intersectionsPool)
	intersectShapes(scene, ray, intersections, scene.rayPool)
	intersects := intersections.Buf[:intersections.Len]
	if len(intersects) == 0 {
		return Black
	} else {
		hitIndex := hitIndex(&intersects)
		if hitIndex == -1 {
			return Black
		} else {
			hit := intersects[hitIndex]
			hitRecord := createHitRecord(&hit, ray, scene.tuplePool)
			return shadeHit(scene, hitRecord)
		}
	}

}

func hitIndex(intersects *[]Intersect) int {
	for i, intersect := range *intersects {
		if intersect.T > 0 {
			return i
		}
	}
	return -1
}

func intersectShapes(scene *Scene, worldRay *Ray, intersections *Intersections, rayPool *sync.Pool) {
	shapes := scene.shapes
	for _, shape := range shapes {
		intersectShape(shape, worldRay, intersections, rayPool)
	}
	intersects := intersections.Buf[:intersections.Len]
	sort.Slice(intersects, func(i, j int) bool {
		return intersects[i].T < intersects[j].T
	})
}

func createHitRecord(intersect *Intersect, ray *Ray, tuplePool *sync.Pool) HitRecord {
	point := ray.Position(intersect.T)
	normalV := intersect.Shape.normalAt(&point, tuplePool)
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

func shadeHit(scene *Scene, record HitRecord) Color {
	total := Black
	for _, light := range scene.pointLights {
		shadowed := isShadowed(scene, light, record.overPoint)
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

func isShadowed(scene *Scene, light PointLight, point Tuple) bool {
	v := light.Position.Sub(point)
	distance := v.Magnitude()
	direction := v.Normalize()
	rayPool := scene.rayPool
	ray := rayPool.Get().(*Ray)
	defer rayPool.Put(ray)
	ray.Origin = &point
	ray.Direction = &direction
	pool := scene.intersectionsPool
	intersections := pool.Get().(*Intersections)
	defer intersections.Return(pool)

	intersectShapes(scene, ray, intersections, scene.rayPool)
	intersects := intersections.Buf[:intersections.Len]

	return hitIndex(&intersects) >= 0 && intersects[0].T < distance
}
