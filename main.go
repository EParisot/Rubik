package main

import (
	"github.com/divan/three"
	"github.com/gopherjs/gopherjs/js"
)

type Env struct {
	renderer *three.WebGLRenderer
	scene    *three.Scene
	camera   *three.PerspectiveCamera
}

func (env *Env) graphics() {
	width := js.Global.Get("innerWidth").Float()
	height := js.Global.Get("innerHeight").Float()

	renderer := three.NewWebGLRenderer()
	env.renderer = &renderer
	env.renderer.SetSize(width, height, true)
	js.Global.Get("document").Get("body").Call("appendChild", env.renderer.Get("domElement"))

	// setup camera and scene
	camera := three.NewPerspectiveCamera(70, width/height, 1, 1000)
	env.camera = &camera
	env.camera.Position.Set(400, 200, 500)

	env.scene = three.NewScene()

	// lights
	light := three.NewDirectionalLight(three.NewColor("white"), 1)
	light.Position.Set(0, 256, 256)
	env.scene.Add(light)

	// material
	params := three.NewMaterialParameters()
	params.Color = three.NewColor("white")
	mat := three.NewMeshLambertMaterial(params)

	// cube object
	cw := 50.
	margin := 2
	for z := 0; z < 3; z++ {
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				geom := three.NewBoxGeometry(&three.BoxGeometryParameters{
					Width:  cw,
					Height: cw,
					Depth:  cw,
				})
				mesh := three.NewMesh(geom, mat)

				mesh.Position.Set(float64(x)*cw+float64(x*margin),
					float64(y)*cw+float64(y*margin),
					float64(z)*cw+float64(z*margin))

				env.scene.Add(mesh)
			}
		}
	}
	env.renderer.Render(env.scene, *env.camera)
}

func (env *Env) animate() {
	js.Global.Call("requestAnimationFrame", env.animate)

	//mesh.Rotation.Set("y", mesh.Rotation.Get("y").Float()+0.01)
	env.renderer.Render(env.scene, *env.camera)
}

func main() {

	env := Env{}
	// init graphics
	env.graphics()
	// TODO handle inputs

	// start animation
	//env.animate()

}
