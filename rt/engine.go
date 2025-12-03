package rt

import (
	"image/color"
	"math"
	"sort"
	"sync"

	. "goray/math"
)

const maxIntersections = 128

type Engine struct {
	scene             *Scene
	intersectionsPool *sync.Pool
	rayPool           *sync.Pool
	tuplePool         *sync.Pool
}

type Ray struct {
	Origin    *Tuple
	Direction *Tuple
}

func NRay(origin, direction *Tuple) Ray {
	result := Ray{}
	result.Origin = origin
	result.Direction = direction
	return result
}

func (ray *Ray) Position(t float64) Tuple {
	return ray.Origin.Add(ray.Direction.Scale(t))
}

func (ray *Ray) TransformToShape(s *Shape, localRayBuffer *Ray) {
	t := s.InverseTransformation
	o := ray.Origin
	d := ray.Direction
	for r := 0; r < 4; r++ {
		localRayBuffer.Origin[r] = t[r][0]*o[0] + t[r][1]*o[1] + t[r][2]*o[2] + t[r][3]*o[3]
		localRayBuffer.Direction[r] = t[r][0]*d[0] + t[r][1]*d[1] + t[r][2]*d[2] + t[r][3]*d[3]
	}
}

type Intersection struct {
	t     float64
	shape *Shape
}

type Intersections struct {
	array      []Intersection
	writeIndex int
}

func (i *Intersections) Add(t float64, shape *Shape) {
	if i.writeIndex == maxIntersections {
		panic("too many intersections - consider increasing maxIntersections or changing bounding strategy")
	}
	i.array[i.writeIndex].t = t
	i.array[i.writeIndex].shape = shape
	i.writeIndex++
}

func (i *Intersections) Return(pool *sync.Pool) {
	i.writeIndex = 0
	pool.Put(i)
}

type HitRecord struct {
	t         float64
	shape     *Shape
	point     Tuple
	overPoint Tuple
	eyeV      Tuple
	normalV   Tuple
	inside    bool
}

func NEngine(scene *Scene) *Engine {
	intersectionsPool := &sync.Pool{
		New: func() interface{} {
			return &Intersections{array: make([]Intersection, maxIntersections)}
		},
	}
	rayPool := &sync.Pool{
		New: func() interface{} {
			return &Ray{&Tuple{}, &Tuple{}}
		},
	}
	return &Engine{
		scene:             scene,
		intersectionsPool: intersectionsPool,
		rayPool:           rayPool,
	}
}

func GetPixel(engine *Engine, x, y int) color.Color {
	scene := engine.scene
	ray := scene.camera.rayForPixel(x, y)
	return colorAt(engine, &ray)
}

func colorAt(engine *Engine, ray *Ray) Color {
	intersectionsPool := engine.intersectionsPool
	intersections := intersectionsPool.Get().(*Intersections)
	defer intersections.Return(intersectionsPool)
	intersectShapes(engine, ray, intersections)
	intersectionsSlice := intersections.array[:intersections.writeIndex]
	if len(intersectionsSlice) == 0 {
		return Black
	} else {
		hitIndex := getHitIndex(intersectionsSlice)
		if hitIndex == -1 {
			return Black
		} else {
			hit := intersectionsSlice[hitIndex]
			hitRecord := createHitRecord(&hit, ray)
			return shadeHit(engine, hitRecord)
		}
	}

}

func intersectShapes(engine *Engine, worldRay *Ray, intersections *Intersections) {
	scene := engine.scene
	shapes := scene.shapes
	for _, shape := range shapes {
		IntersectShape(engine, shape, worldRay, intersections)
	}
	intersectionsSlice := intersections.array[:intersections.writeIndex]
	sort.Slice(intersectionsSlice, func(i, j int) bool {
		return intersectionsSlice[i].t < intersectionsSlice[j].t
	})
}

func IntersectShape(engine *Engine, shape *Shape, worldRay *Ray, intersections *Intersections) {
	localRay := engine.rayPool.Get().(*Ray)
	defer engine.rayPool.Put(localRay)
	worldRay.TransformToShape(shape, localRay)

	shape.strategy.LocalIntersect(engine, shape, localRay, intersections)
}

func shadeHit(engine *Engine, record *HitRecord) Color {
	total := Black
	for _, light := range engine.scene.pointLights {
		shadowed := isInShadow(engine, light, record.overPoint)
		total = total.Add(lighting(record.shape.getMaterial(), light, record.point, record.eyeV, record.normalV, shadowed))
	}
	return total
}

func lighting(material *Material, light *PointLight, point Tuple, eyeV Tuple, normalV Tuple, shadowed bool) Color {
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

func getHitIndex(intersectionsSlice []Intersection) int {
	for i, intersect := range intersectionsSlice {
		if intersect.t > 0 {
			return i
		}
	}
	return -1
}

func createHitRecord(intersect *Intersection, ray *Ray) *HitRecord {
	point := ray.Position(intersect.t)
	normalV := normalAt(intersect.shape, &point)
	eyeV := ray.Direction.Negate()
	inside := false
	if normalV.Dot(eyeV) < 0 {
		inside = true
		normalV = normalV.Negate()
	}
	overPoint := point.Add(normalV.Scale(EPSILON))
	return &HitRecord{
		t:         intersect.t,
		shape:     intersect.shape,
		point:     point,
		overPoint: overPoint,
		eyeV:      eyeV,
		normalV:   normalV,
		inside:    inside,
	}
}

func normalAt(shape *Shape, worldPoint *Tuple) Tuple {
	localPoint := worldToObject(shape, *worldPoint)
	localNormal := shape.strategy.LocalNormalAt(shape, &localPoint)
	return normalToWorld(shape, localNormal)
}

func worldToObject(shape *Shape, point Tuple) Tuple {
	if shape.Parent != nil {
		point = worldToObject(shape.Parent, point)
	}
	return shape.InverseTransformation.MulT(point)
}

func normalToWorld(shape *Shape, normal Tuple) Tuple {
	normal = shape.TransposeInverse.MulT(normal)
	normal[3] = 0
	normal = normal.Normalize()
	if shape.Parent != nil {
		normal = normalToWorld(shape.Parent, normal)
	}
	return normal
}

func reflect(in Tuple, normal Tuple) Tuple {
	return in.Sub(normal.Scale(2 * in.Dot(normal)))
}

func isInShadow(engine *Engine, light *PointLight, point Tuple) bool {
	v := light.Position.Sub(point)
	distance := v.Magnitude()
	direction := v.Normalize()
	ray := engine.rayPool.Get().(*Ray)
	defer engine.rayPool.Put(ray)
	ray.Origin = &point
	ray.Direction = &direction
	intersections := engine.intersectionsPool.Get().(*Intersections)
	defer intersections.Return(engine.intersectionsPool)

	intersectShapes(engine, ray, intersections)
	intersectionsSlice := intersections.array[:intersections.writeIndex]

	return getHitIndex(intersectionsSlice) >= 0 && intersectionsSlice[0].t < distance
}
