package Logger

import (
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	Debugf("Debug !")
	Infof("Info !")
	Warnf("Warning !")
	Errorf("Error !")
}
