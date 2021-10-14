package main

import (
	"KubeCaption.Api/internal/handlers/core"
	"KubeCaption.Api/internal/middleware"
	caption "KubeCaption.Api/pkg/captain"
	"KubeCaption.Api/pkg/jwt"
	"fmt"
)

func main() {
	token, _ :=jwt.GenerateToken(12, "ben.c", "ku123test", "www.kubecaptain.com")
	fmt.Println(token)
	caption.NewCaptain().Attach(
		middleware.JWT(),
	).Mount("v1",
		core.NewPodHandler(),
	).Launch()
}


