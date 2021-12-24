package driver_gozero

import (
	"github.com/dtm-labs/dtmdriver"
	_ "github.com/dtm-labs/dtmdriver-gozero"
)

func init() {
	err := dtmdriver.Use("dtm-driver-gozero")
	if err != nil {
		panic(err)
	}
}
