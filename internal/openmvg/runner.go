package openmvg

import (
	"fmt"
	"os"
	"time"

	"github.com/soup666/openMVGO/internal/utils"
)

// Config for running the OpenMVG pipeline
type OpenMVGConfig struct {
	InputDir          string
	OutputDir         string
	MatchesDir        string
	ReconstructionDir string
	CameraDBFile      *string
}

// Create an OpenMVG config. Handles creating tempory directories for mvs and reconstruction
func NewOpenMVGConfig(inputDir string, outputDir string, cameraDBFile *string) OpenMVGConfig {

	return OpenMVGConfig{
		InputDir:     inputDir,
		OutputDir:    outputDir,
		CameraDBFile: cameraDBFile,
	}
}

type OpenMVGService interface {
	RunHealthCheck()
	SfMSequentialPipeline()
	RunSfMInitImageListing()
	RunSfMComputeFeatures()
	RunSfMPairGenerator()
	RunSfMComputeMatches()
	RunSfMGeometricFilter()
	RunSfMReconstruction()
	RunSfMComputeSfMDataColor()
	PopulateTmpDir()
}

type AppFileServiceImpl struct {
	Config OpenMVGConfig
}

func NewOpenMVGService(config OpenMVGConfig) AppFileServiceImpl {

	// Pre run folder checks
	if config.InputDir == "" || config.OutputDir == "" {
		utils.Check(fmt.Errorf("input and output directories must be specified"))
	}

	if err := utils.EnsureDir(config.InputDir); err != nil {
		utils.Check(fmt.Errorf("failed to ensure input directory: %w", err))
	}

	if err := utils.EnsureDir(config.OutputDir); err != nil {
		utils.Check(fmt.Errorf("failed to ensure output directory: %w", err))
	}

	return AppFileServiceImpl{
		Config: config,
	}
}

func (s *AppFileServiceImpl) PopulateTmpDir() {
	// Ensure the camera database file is set, if not download it
	if s.Config.CameraDBFile == nil || *s.Config.CameraDBFile == "" {
		f, err := utils.DownloadFile("https://raw.githubusercontent.com/openMVG/openMVG/refs/heads/develop/src/openMVG/exif/sensor_width_database/sensor_width_camera_database.txt")
		utils.Check(err)
		s.Config.CameraDBFile = &f
	}

	timestamp := time.Now().Unix()

	// Matches dir
	matchesDir, err := os.MkdirTemp("", fmt.Sprintf("%dmatches", timestamp))
	utils.Check(err)

	s.Config.MatchesDir = matchesDir

	// Reconstruction dir
	reconstructionDir, err := os.MkdirTemp("", fmt.Sprintf("%dreconstruction", timestamp))
	utils.Check(err)

	s.Config.ReconstructionDir = reconstructionDir

}

func (s *AppFileServiceImpl) SfMSequentialPipeline() {
	s.PopulateTmpDir()
	defer os.Remove(*s.Config.CameraDBFile)
	defer os.RemoveAll(s.Config.MatchesDir)
	defer os.RemoveAll(s.Config.ReconstructionDir)

	s.RunSfMInitImageListing()
	s.RunSfMComputeFeatures()
	s.RunSfMPairGenerator()
	s.RunSfMComputeMatches()
	s.RunSfMGeometricFilter()
	s.RunSfMReconstruction()
	s.RunSfMComputeSfMDataColor()
}

func (s *AppFileServiceImpl) RunHealthCheck() {
	utils.RunCommand("Tests", []string{})
}

func (s *AppFileServiceImpl) RunSfMInitImageListing() {
	args := []string{
		"-i", s.Config.InputDir,
		"-o", s.Config.MatchesDir,
		"-d", *s.Config.CameraDBFile,
		"-f", "2304", // 2304 is the focal length, adjust as necessary
	}

	utils.RunCommand("openMVG_main_SfMInit_ImageListing", args)
}

func (s *AppFileServiceImpl) RunSfMComputeFeatures() {
	args := []string{
		"-i", s.Config.MatchesDir + "/sfm_data.json",
		"-o", s.Config.MatchesDir,
		"-m", "SIFT",
	}

	utils.RunCommand("openMVG_main_ComputeFeatures", args)
}

func (s *AppFileServiceImpl) RunSfMPairGenerator() {
	args := []string{
		"-i", s.Config.MatchesDir + "/sfm_data.json",
		"-o", s.Config.MatchesDir + "/pairs.bin",
	}

	utils.RunCommand("openMVG_main_PairGenerator", args)
}

func (s *AppFileServiceImpl) RunSfMComputeMatches() {
	args := []string{
		"-i", s.Config.MatchesDir + "/sfm_data.json",
		"-p", s.Config.MatchesDir + "/pairs.bin",
		"-o", s.Config.MatchesDir + "/matches.putative.bin",
	}

	utils.RunCommand("openMVG_main_ComputeMatches", args)
}

func (s *AppFileServiceImpl) RunSfMGeometricFilter() {
	args := []string{
		"-i", s.Config.MatchesDir + "/sfm_data.json",
		"-m", s.Config.MatchesDir + "/matches.putative.bin",
		"-g", "f",
		"-o", s.Config.MatchesDir + "/matches.f.bin",
	}

	utils.RunCommand("openMVG_main_GeometricFilter", args)
}

func (s *AppFileServiceImpl) RunSfMReconstruction() {
	args := []string{
		"--sfm_engine", "INCREMENTAL",
		"--input_file", s.Config.MatchesDir + "/sfm_data.json",
		"--match_dir", s.Config.MatchesDir,
		"--output_dir", s.Config.ReconstructionDir,
	}

	utils.RunCommand("openMVG_main_SfM", args)
}

func (s *AppFileServiceImpl) RunSfMComputeSfMDataColor() {
	args := []string{
		"-i", s.Config.ReconstructionDir + "/sfm_data.bin",
		"-o", s.Config.ReconstructionDir + "/colorized.ply",
	}

	utils.RunCommand("openMVG_main_ComputeSfM_DataColor", args)
	utils.CopyFile(s.Config.ReconstructionDir+"/colorized.ply", s.Config.OutputDir+"colorized.ply")
	utils.CopyFile(s.Config.ReconstructionDir+"/sfm_data.bin", s.Config.OutputDir+"sfm_data.bin")
}
