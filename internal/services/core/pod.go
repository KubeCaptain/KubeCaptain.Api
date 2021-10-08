package core

import (
	"KubeCaption.Api/internal/dto/core"
	"KubeCaption.Api/pkg/kube"
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodListByNamespace(namespace string) (podList []core.PodDTO, err error) {

	ctx := context.Background()
	opts := metav1.ListOptions{}
	client := kube.NewClient()
	list, err := client.CoreV1().Pods(namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}

	for _, item := range list.Items {
		podList = append(podList, core.PodDTO{
			Name:        item.Name,
			Namespace:   item.Namespace,
			Labels:      item.Labels,
			Annotations: item.Annotations,
			PodIP:       item.Status.PodIP,
			HostIP:      item.Status.HostIP,
			Image:       item.Spec.Containers[0].Image,
		})
	}

	return
}

func GetPodByPodName(namespace, podName string) (pod core.PodDTO, err error) {

	ctx := context.Background()
	opts := metav1.GetOptions{}
	client := kube.NewClient()

	item, err := client.CoreV1().Pods(namespace).Get(ctx, podName, opts)
	if err != nil {
		return pod, err
	}

	return core.PodDTO{
		Name:        item.Name,
		Namespace:   item.Namespace,
		Labels:      item.Labels,
		Annotations: item.Annotations,
		PodIP:       item.Status.PodIP,
		HostIP:      item.Status.HostIP,
		Image:       item.Spec.Containers[0].Image,
	}, nil

}

func DeletePodByPodName(namespace, podName string) error {

	ctx := context.Background()
	opts := metav1.DeleteOptions{}
	client := kube.NewClient()

	return client.CoreV1().Pods(namespace).Delete(ctx, podName, opts)

}

func ApplyPod(pod *core.PodTemplate) (err error) {

	client := kube.NewClient()

	// 1. the pod is exist
	_, err = client.CoreV1().Pods(pod.Namespace).Get(context.Background(), pod.Name, metav1.GetOptions{})

	if err == nil {
		_ = client.CoreV1().Pods(pod.Namespace).Delete(context.Background(), pod.Name, metav1.DeleteOptions{})
	}

	obj := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  pod.Name,
					Image: pod.Image,
				},
			},
		},
	}

	ctx := context.Background()
	opts := metav1.CreateOptions{}
	_, err = client.CoreV1().Pods(pod.Namespace).Create(ctx, obj, opts)

	return
}
