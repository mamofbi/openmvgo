package openmvg

import (
	"github.com/soup666/openMVGO/internal/utils"
)

type OpenMVGConfig struct {
	OpenMVGBinDir string
	BinDir        string
	CameraDBFile  string
	ImageWidth    string
}

type OpenMVGService interface {
	RunHealthCheck()
}

type AppFileServiceImpl struct {
	config OpenMVGConfig
}

func NewOpenMVGService(config OpenMVGConfig) OpenMVGService {
	return &AppFileServiceImpl{
		config: config,
	}
}

func (s *AppFileServiceImpl) RunHealthCheck() {
	utils.RunCommand("Tests", []string{})
}
