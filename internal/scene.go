package internal

type PointLight struct {
	Position  Tuple
	Intensity Color
}

type Scene struct {
	pointLights []*PointLight
	camera      *Camera
	rootGroup   *Group
}

func newScene(pointLights []PointLight, camera Camera) *Scene {
	scene := &Scene{
		pointLights: make([]*PointLight, len(pointLights)),
		camera:      &camera,
		rootGroup:   newGroup(),
	}
	for i, pointLight := range pointLights {
		scene.pointLights[i] = &pointLight
	}
	return scene
}
