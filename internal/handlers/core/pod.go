package core

import (
	"KubeCaption.Api/internal/dto/core"
	service "KubeCaption.Api/internal/services/core"
	"KubeCaption.Api/pkg/captain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PodHandler struct{}

func NewPodHandler() *PodHandler {
	return &PodHandler{}
}

func (h *PodHandler) ListPod(c *gin.Context) {

	list, err := service.GetPodListByNamespace(c.Param("namespace"))
	if err != nil {
		c.JSON(http.StatusOK, err)
	}

	c.JSON(http.StatusOK, list)
}

func (h *PodHandler) GetPod(c *gin.Context) {

	list, err := service.GetPodByPodName(c.Param("namespace"), c.Param("podName"))
	if err != nil {
		c.JSON(http.StatusOK, err)
	}

	c.JSON(http.StatusOK, list)
}

func (h *PodHandler) DeletePod(c *gin.Context) {

	err := service.DeletePodByPodName(c.Param("namespace"), c.Param("podName"))
	// TODO
	// Define error message process logic
	//
	c.JSON(http.StatusOK, err)

}

func (h *PodHandler) ApplyPod(c *gin.Context) {

	pod := &core.PodTemplate{}

	if err := c.ShouldBindJSON(pod); err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	err := service.ApplyPod(pod)

	// TODO
	// Define error message process logic
	//
	c.JSON(http.StatusOK, err)

}

func (h *PodHandler) Build(e *captain.Captain) {
	e.Handle("GET", "/pods/:namespace", h.ListPod)
	e.Handle("GET", "/pods/:namespace/:podName", h.GetPod)
	e.Handle("DELETE", "/pods/:namespace/:podName", h.DeletePod)
	e.Handle("POST", "/pods", h.ApplyPod)
}
