package internal

import (
	"image/color"
	"math"
	"sort"
	"sync"
)

const maxIntersections = 128
const groupSplitThreshold = 10
const maxDepth = 10

type Engine struct {
	scene             *Scene
	intersectionsPool *sync.Pool
	rayPool           *sync.Pool
	tuplePool         *sync.Pool
}

func NewEngine(scene *Scene) *Engine {
	intersectionsPool := &sync.Pool{
		New: func() interface{} {
			return &Intersections{array: make([]Intersection, maxIntersections)}
		},
	}
	rayPool := &sync.Pool{
		New: func() interface{} {
			return &Ray{Tuple{}, Tuple{}}
		},
	}

	calculateBounds(scene.rootGroup)
	rootGroup := scene.rootGroup
	rootGroup.children = unGroup(rootGroup, true)
	divideGroup(rootGroup, groupSplitThreshold)

	return &Engine{
		scene:             scene,
		intersectionsPool: intersectionsPool,
		rayPool:           rayPool,
	}
}

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func newRay(origin, direction Tuple) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (ray *Ray) Position(t float64) Tuple {
	return ray.Origin.Add(ray.Direction.Scale(t))
}

type Intersection struct {
	t      float64
	object SceneObject
}

type Intersections struct {
	array      []Intersection
	writeIndex int
}

func (i *Intersections) add(t float64, object SceneObject) {
	if i.writeIndex == maxIntersections {
		panic("too many intersections - consider increasing maxIntersections or changing bounding strategy")
	}
	i.array[i.writeIndex].t = t
	i.array[i.writeIndex].object = object
	i.writeIndex++
}

func (i *Intersections) Return(pool *sync.Pool) {
	i.writeIndex = 0
	pool.Put(i)
}

type HitRecord struct {
	t         float64
	object    SceneObject
	point     Tuple
	overPoint Tuple
	eyeV      Tuple
	normalV   Tuple
	reflectV  Tuple
	inside    bool
}

func GetPixel(engine *Engine, x, y int) color.Color {
	scene := engine.scene
	ray := scene.camera.rayForPixel(x, y)
	return colorAt(engine, ray, 0)
}

func colorAt(engine *Engine, ray *Ray, depth int) Color {
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
			return shadeHit(engine, hitRecord, depth)
		}
	}

}

func intersectShapes(engine *Engine, worldRay *Ray, intersections *Intersections) {
	scene := engine.scene
	rootGroup := scene.rootGroup
	intersectObject(engine, rootGroup, worldRay, intersections)
	intersectionsSlice := intersections.array[:intersections.writeIndex]
	sort.Slice(intersectionsSlice, func(i, j int) bool {
		return intersectionsSlice[i].t < intersectionsSlice[j].t
	})
}

func intersectObject(engine *Engine, object SceneObject, worldRay *Ray, intersections *Intersections) {
	if !object.getCommonState().isIdentity {
		localRayBuffer := engine.rayPool.Get().(*Ray)
		defer engine.rayPool.Put(localRayBuffer)

		t := object.getCommonState().inverseTransformation

		o := worldRay.Origin
		d := worldRay.Direction

		localRayBuffer.Origin[0] = t[0][0]*o[0] + t[0][1]*o[1] + t[0][2]*o[2] + t[0][3]*o[3]
		localRayBuffer.Origin[1] = t[1][0]*o[0] + t[1][1]*o[1] + t[1][2]*o[2] + t[1][3]*o[3]
		localRayBuffer.Origin[2] = t[2][0]*o[0] + t[2][1]*o[1] + t[2][2]*o[2] + t[2][3]*o[3]
		localRayBuffer.Origin[3] = t[3][0]*o[0] + t[3][1]*o[1] + t[3][2]*o[2] + t[3][3]*o[3]

		localRayBuffer.Direction[0] = t[0][0]*d[0] + t[0][1]*d[1] + t[0][2]*d[2] + t[0][3]*d[3]
		localRayBuffer.Direction[1] = t[1][0]*d[0] + t[1][1]*d[1] + t[1][2]*d[2] + t[1][3]*d[3]
		localRayBuffer.Direction[2] = t[2][0]*d[0] + t[2][1]*d[1] + t[2][2]*d[2] + t[2][3]*d[3]
		localRayBuffer.Direction[3] = t[3][0]*d[0] + t[3][1]*d[1] + t[3][2]*d[2] + t[3][3]*d[3]

		object.localIntersect(engine, localRayBuffer, intersections)
	} else {
		object.localIntersect(engine, worldRay, intersections)
	}
}

func shadeHit(engine *Engine, record *HitRecord, depth int) Color {
	resultColor := Black
	for _, light := range engine.scene.pointLights {
		shadowed := isInShadow(engine, light, record.overPoint)
		surface := resultColor.Add(lighting(record.object.getCommonState().material, light, record.point, record.eyeV, record.normalV, shadowed))
		resultColor = resultColor.Add(surface)
		if depth < maxDepth-1 {
			reflected := reflectedColor(engine, record, depth)
			resultColor = resultColor.Add(reflected)
		}
	}
	return resultColor
}

func lighting(material *Material, light *PointLight, point Tuple, eyeV Tuple, normalV Tuple, shadowed bool) Color {
	effectiveColor := material.pattern.localPatternAt(point).Multiply(light.Intensity)
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

func reflectedColor(engine *Engine, record *HitRecord, depth int) Color {
	reflectivity := record.object.getCommonState().material.Reflectivity
	if reflectivity == 0 {
		return Black
	}
	reflectRay := newRay(record.overPoint, record.reflectV)
	return colorAt(engine, reflectRay, depth+1).Scale(reflectivity)
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
	normalV := normalAt(intersect.object, &point)
	eyeV := ray.Direction.Negate()
	inside := false
	if normalV.Dot(eyeV) < 0 {
		inside = true
		normalV = normalV.Negate()
	}
	overPoint := point.Add(normalV.Scale(EPSILON))
	reflectV := reflect(ray.Direction, normalV)
	return &HitRecord{
		t:         intersect.t,
		object:    intersect.object,
		point:     point,
		overPoint: overPoint,
		eyeV:      eyeV,
		normalV:   normalV,
		reflectV:  reflectV,
		inside:    inside,
	}
}

func normalAt(object SceneObject, worldPoint *Tuple) Tuple {
	localPoint := worldToObject(object, *worldPoint)
	localNormal := object.localNormalAt(&localPoint)
	return normalToWorld(object, localNormal)
}

func worldToObject(object SceneObject, point Tuple) Tuple {
	state := object.getCommonState()
	if state.parent != nil {
		point = worldToObject(state.parent, point)
	}
	return state.inverseTransformation.MulT(point)
}

func normalToWorld(object SceneObject, normal Tuple) Tuple {
	state := object.getCommonState()
	normal = state.transposeInverse.MulT(normal)
	normal[3] = 0
	normal = normal.Normalize()
	if state.parent != nil {
		normal = normalToWorld(state.parent, normal)
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
	ray.Origin = point
	ray.Direction = direction
	intersections := engine.intersectionsPool.Get().(*Intersections)
	defer intersections.Return(engine.intersectionsPool)

	intersectShapes(engine, ray, intersections)
	intersectionsSlice := intersections.array[:intersections.writeIndex]

	hitIndex := getHitIndex(intersectionsSlice)
	return hitIndex >= 0 && intersectionsSlice[hitIndex].t < distance
}
