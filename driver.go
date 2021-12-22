package driver_gozero

import (
	"github.com/yedf/dtmdriver"
	_ "github.com/yedf/dtmdriver-gozero"
)

func init() {
	dtmdriver.Use("dtm-driver-gozero")
}
