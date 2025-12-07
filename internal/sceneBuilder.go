package internal

import "math"

type SceneBuilder struct {
	width, height int
	scene         *Scene
}

func buildScene(width, height int) *SceneBuilder {
	scene := &Scene{
		pointLights: make([]*PointLight, 0),
		camera:      newCamera(width, height, math.Pi/2, Point(0, 0, -5), Point(0, 0, 0), Vector(0, 1, 0)),
		rootGroup:   newGroup(),
	}
	return &SceneBuilder{width: width, height: height, scene: scene}
}

func (sb *SceneBuilder) addPointLight(position Tuple, intensity Color) *SceneBuilder {
	sb.scene.pointLights = append(sb.scene.pointLights, &PointLight{
		Position:  position,
		Intensity: intensity,
	})
	return sb
}

func (sb *SceneBuilder) withCamera(fov float64, from, lookAt, up Tuple) *SceneBuilder {
	sb.scene.camera = newCamera(sb.width, sb.height, fov, from, lookAt, up)
	return sb
}

func (sb *SceneBuilder) add(sceneObjectBuilder *SceneObjectBuilder) *SceneBuilder {
	sb.scene.rootGroup.add(sceneObjectBuilder.build())
	return sb
}

func (sb *SceneBuilder) build() *Scene {
	return sb.scene
}

type SceneObjectBuilder struct {
	sceneObject SceneObject
}

func (sceneObjectBuilder *SceneObjectBuilder) add(child *SceneObjectBuilder) *SceneObjectBuilder {
	switch sceneObjectBuilder.sceneObject.(type) {
	case *Group:
		sceneObjectBuilder.sceneObject.(*Group).add(child.build())
	default:
		panic("can only add children to groups")
	}
	return sceneObjectBuilder
}

func (sceneObjectBuilder *SceneObjectBuilder) build() SceneObject {
	return sceneObjectBuilder.sceneObject
}

func (sceneObjectBuilder *SceneObjectBuilder) material(material Material) *SceneObjectBuilder {
	sceneObjectBuilder.sceneObject.getCommonState().material = &material
	return sceneObjectBuilder
}

func (sceneObjectBuilder *SceneObjectBuilder) transform(transformation *Matrix4x4) *SceneObjectBuilder {
	state := sceneObjectBuilder.sceneObject.getCommonState()
	t := state.transformation
	state.transformation = transformation.mul4x4(t)
	return sceneObjectBuilder
}

func (sceneObjectBuilder *SceneObjectBuilder) scaleXYZ(x, y, z float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(Scaling(x, y, z))
}

func (sceneObjectBuilder *SceneObjectBuilder) scale(s float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(Scaling(s, s, s))
}

func (sceneObjectBuilder *SceneObjectBuilder) scaleX(x float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(Scaling(x, 1, 1))
}

func (sceneObjectBuilder *SceneObjectBuilder) scaleY(y float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(Scaling(1, y, 1))
}

func (sceneObjectBuilder *SceneObjectBuilder) scaleZ(z float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(Scaling(1, 1, z))
}

func (sceneObjectBuilder *SceneObjectBuilder) translate(x, y, z float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(Translation(x, y, z))
}

func (sceneObjectBuilder *SceneObjectBuilder) translateX(x float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(Translation(x, 0, 0))
}

func (sceneObjectBuilder *SceneObjectBuilder) translateY(y float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(Translation(0, y, 0))
}

func (sceneObjectBuilder *SceneObjectBuilder) translateZ(z float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(Translation(0, 0, z))
}

func (sceneObjectBuilder *SceneObjectBuilder) rotateX(theta float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(RotationX(theta))
}

func (sceneObjectBuilder *SceneObjectBuilder) rotateY(theta float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(RotationY(theta))
}

func (sceneObjectBuilder *SceneObjectBuilder) rotateZ(theta float64) *SceneObjectBuilder {
	return sceneObjectBuilder.transform(RotationZ(theta))
}

func object(sceneObject SceneObject) *SceneObjectBuilder {
	return &SceneObjectBuilder{sceneObject: sceneObject}
}

func group() *SceneObjectBuilder {
	return &SceneObjectBuilder{sceneObject: newGroup()}
}
