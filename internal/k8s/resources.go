package k8s

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPods retrieves pods in a given namespace, utilizing cache if possible.
func GetPods(ctx context.Context, namespace string) ([]corev1.Pod, error) {
	cacheKey := fmt.Sprintf("pods:%s", namespace)
	if val, found := GetCachedResource(cacheKey); found {
		if pods, ok := val.([]corev1.Pod); ok {
			return pods, nil
		}
	}

	if GlobalClient == nil || GlobalClient.Clientset == nil {
		return nil, fmt.Errorf("kubernetes client not initialized")
	}

	podList, err := GlobalClient.Clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %w", err)
	}

	SetCachedResource(cacheKey, podList.Items, 1, 30*time.Second)
	return podList.Items, nil
}

// GetPod retrieves a specific pod in a given namespace.
func GetPod(ctx context.Context, namespace, name string) (*corev1.Pod, error) {
	if GlobalClient == nil || GlobalClient.Clientset == nil {
		return nil, fmt.Errorf("kubernetes client not initialized")
	}
	pod, err := GlobalClient.Clientset.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pod %s in %s: %w", name, namespace, err)
	}
	return pod, nil
}
