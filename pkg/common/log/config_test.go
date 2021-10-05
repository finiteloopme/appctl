package log

import (
	"os"
	"testing"

	"github.com/finiteloopme/appctl/pkg/apis/core/v1alpha1"
)

func TestInfo(t *testing.T) {
	os.Setenv("LOG_FORMAT", v1alpha1.LogFormatJSON.String())
	//TODO: no need to test this function?
}
