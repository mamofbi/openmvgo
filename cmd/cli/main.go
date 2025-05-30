package main

import "github.com/soup666/openMVGO/internal/openmvg"

func main() {
	openmvg := openmvg.NewOpenMVGService(openmvg.OpenMVGConfig{
		OpenMVGBinDir: "/usr/local/bin",
		BinDir:        "./bin",
		CameraDBFile:  "sensor_width_camera_database.txt",
		ImageWidth:    "2304",
	})

	openmvg.RunHealthCheck()
}
