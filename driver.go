package driver_gozero

import (
	"github.com/yedf/dtmdriver"
	_ "github.com/yedf/dtmdriver-gozero"
)

func init() {
	err := dtmdriver.Use("dtm-driver-gozero")
	if err != nil {
		panic(err)
	}
}
