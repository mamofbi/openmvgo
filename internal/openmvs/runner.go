package openmvs

import (
	"fmt"

	"github.com/soup666/openMVGO/internal/utils"
)

/*
log.Println("# 3 DensifyPointCloud", "scene.mvs", "-o", "scene_dense.mvs", "-w", mvsPath, "--max-threads", "1")
log.Println("# 4 ReconstructMesh", "scene_dense.mvs", "-o", "scene_mesh.ply", "-w", mvsPath)
log.Println("# 5 RefineMesh", "scene.mvs", "-m", "scene_mesh.ply", "-o", "scene_dense_mesh_refine.mvs", "-w", mvsPath, "--scales", "1", "--max-face-area", "16", "--max-threads", "1")
log.Println("# 6 TextureMesh", "scene_dense.mvs", "-m", "scene_dense_mesh_refine.ply", "-o", "scene_dense_mesh_refine_texture.mvs", "-w", mvsPath, "--export-type", "obj")
*/

type OpenMVSConfig struct {
	InputDir string
}

func NewOpenMVSConfig(inputDir string) *OpenMVSConfig {
	return &OpenMVSConfig{
		InputDir: inputDir,
	}
}

type OpenMVSService interface {
	RunPipeline()
	RunDensifyPointCloud()
	RunReconstructMesh()
	RunRefineMesh()
	RunTextureMesh()
}

type OpenMVSServiceImpl struct {
	Config *OpenMVSConfig
}

func NewOpenMVSService(config *OpenMVSConfig) OpenMVSServiceImpl {
	if config.InputDir == "" {
		utils.Check(fmt.Errorf("input directory must be specified"))
	}

	if err := utils.EnsureDir(config.InputDir); err != nil {
		utils.Check(fmt.Errorf("failed to ensure input directory: %w", err))
	}

	return OpenMVSServiceImpl{
		Config: config,
	}
}

func (s OpenMVSServiceImpl) RunPipeline() {
	s.RunDensifyPointCloud()
	s.RunReconstructMesh()
	s.RunRefineMesh()
	s.RunTextureMesh()
}

func (s OpenMVSServiceImpl) RunDensifyPointCloud() {
	mvsPath := s.Config.InputDir

	if err := utils.RunCommand("DensifyPointCloud", []string{"scene.mvs", "-o", "scene_dense.mvs", "-w", mvsPath, "--max-threads", "1"}); err != nil {
		utils.Check(fmt.Errorf("failed to run DensifyPointCloud: %w", err))
	}
}

func (s OpenMVSServiceImpl) RunReconstructMesh() {
	mvsPath := s.Config.InputDir

	if err := utils.RunCommand("ReconstructMesh", []string{"scene_dense.mvs", "-o", "scene_mesh.ply", "-w", mvsPath}); err != nil {
		utils.Check(fmt.Errorf("failed to run ReconstructMesh: %w", err))
	}
}

func (s OpenMVSServiceImpl) RunRefineMesh() {
	mvsPath := s.Config.InputDir

	if err := utils.RunCommand("RefineMesh", []string{"scene.mvs", "-m", "scene_mesh.ply", "-o", "scene_dense_mesh_refine.mvs", "-w", mvsPath, "--scales", "1", "--max-face-area", "16", "--max-threads", "1"}); err != nil {
		utils.Check(fmt.Errorf("failed to run RefineMesh: %w", err))
	}
}

func (s OpenMVSServiceImpl) RunTextureMesh() {
	mvsPath := s.Config.InputDir

	if err := utils.RunCommand("TextureMesh", []string{"scene_dense.mvs", "-m", "scene_dense_mesh_refine.ply", "-o", "scene_dense_mesh_refine_texture.mvs", "-w", mvsPath, "--export-type", "obj"}); err != nil {
		utils.Check(fmt.Errorf("failed to run TextureMesh: %w", err))
	}
}
