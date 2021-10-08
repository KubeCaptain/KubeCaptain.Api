package core

type PodDTO struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	HostIP      string            `json:"host_ip"`
	PodIP       string            `json:"pod_ip"`
	Image       string            `json:"image"`
}

type PodTemplate struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Image     string `json:"image"`
}
