package glfw

import (
	"math"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
)

const (
	scaleEnvKey = "FYNE_SCALE"
)

func calculateScale(user, system float32) float32 {
	if user < 0 {
		user = 1.0
	}

	if system == 0 {
		system = 1.0
	}

	raw := system * user
	return float32(math.Round(float64(raw*10.0))) / 10.0
}

func userScale() float32 {
	env := os.Getenv(scaleEnvKey)

	if env != "" && env != "auto" {
		scale, err := strconv.ParseFloat(env, 32)
		if err == nil && scale != 0 {
			return float32(scale)
		}
		fyne.LogError("Error reading scale", err)
	}

	if env != "auto" {
		if setting := fyne.CurrentApp().Settings().Scale(); setting > 0 {
			return setting
		}
	}

	return 1.0 // user preference for auto is now passed as 1 so the system auto is picked up
}
