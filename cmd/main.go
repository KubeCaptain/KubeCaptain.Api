package main

import (
	"KubeCaption.Api/internal/handlers/core"
	caption "KubeCaption.Api/pkg/captain"
)

func main() {

	caption.NewCaptain().Mount("v1", core.NewPodHandler()).Launch()

}


