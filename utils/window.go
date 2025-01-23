package utils

import (
	"fmt"
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"unsafe"
)

func OpenWindow() {
	defer rl.CloseWindow()
	setFlags()

	rl.InitWindow(800, 600, "Hello world")

	fontTtf := rl.LoadFontEx("consola.ttf", 32, nil, 250)

	camera := setupCamera()

	cubePosition := rl.NewVector3(0.0, 0.0, 0.0)

	skyboxShader := createSkyboxShader()
	skyboxModel := createSkyboxModel()
	skyboxModel.Materials.Shader = skyboxShader
	skyboxTexture := createSkyBoxTexture()

	rl.SetMaterialTexture(skyboxModel.Materials, rl.MapCubemap, skyboxTexture)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree) // Update camera with free camera mode

		if rl.IsKeyDown(rl.KeyZ) {
			camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		// Start Skybox rendering
		// we are inside the cube, we need to disable backface culling
		rl.DisableBackfaceCulling()
		rl.DisableDepthMask()

		rl.DrawModel(skyboxModel, rl.NewVector3(0, 0, 0), 1.0, rl.White)

		// restore depth and backface culling
		rl.EnableBackfaceCulling()
		rl.EnableDepthMask()
		// End Skybox rendering

		rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawTextEx(fontTtf, "Test text", rl.Vector2{X: 20.0, Y: 100.0}, float32(fontTtf.BaseSize), 2, rl.Lime)

		drawGuiLayer()

		rl.EndDrawing()
	}

	rl.UnloadModel(skyboxModel)
	rl.UnloadTexture(skyboxTexture)
	rl.UnloadShader(skyboxShader)
	rl.UnloadFont(fontTtf)
}

func createSkyBoxTexture() rl.Texture2D {
	skyboxFileName := "skybox.png"
	var skyboxImg = rl.LoadImage(skyboxFileName)
	skyboxTexture := rl.LoadTextureCubemap(skyboxImg, rl.CubemapLayoutAutoDetect)

	return skyboxTexture
}

func createSkyboxShader() rl.Shader {
	skyboxShader := rl.LoadShader("skybox.vs", "skybox.fs")
	setShaderIntValue(skyboxShader, "environmentMap", rl.MapCubemap)

	return skyboxShader
}

func drawGuiLayer() {
	button := gui.Button(rl.NewRectangle(50, 150, 100, 40), "Click!")
	if button {
		fmt.Println("Button clicked!")
	}
}

func setFlags() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagWindowTopmost)
}

func setShaderIntValue(shader rl.Shader, name string, value int32) {
	rl.SetShaderValue(
		shader,
		rl.GetShaderLocation(shader, name),
		unsafe.Slice((*float32)(unsafe.Pointer(&value)), 4),
		rl.ShaderUniformInt,
	)
}

func setupCamera() rl.Camera {
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(1.0, 1.0, 10.0)
	camera.Target = rl.NewVector3(4.0, 1.0, 4.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	return camera
}

func createSkyboxModel() rl.Model {
	skyCube := rl.GenMeshCube(1.0, 1.0, 1.0)
	skybox := rl.LoadModelFromMesh(skyCube)

	return skybox
}
