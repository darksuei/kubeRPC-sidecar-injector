package api

import (
	"encoding/json"

	config "github.com/darksuei/kubeRPC-sidecar-injector/config"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
)

func MutatePod(req admissionv1.AdmissionReview) admissionv1.AdmissionResponse {
    var pod corev1.Pod
    json.Unmarshal(req.Request.Object.Raw, &pod)

    inject, err := config.ReadEnv("ANNOTATION_INJECT")

    if err != nil {
        inject = config.DEFAULT_ANNOTATION_INJECT
    }

    // Check if the annotation exists
    if pod.Annotations[inject] == "true" {
        sidecarPodName, err := config.ReadEnv("SIDECAR_POD_NAME")
        
        if err != nil {
            sidecarPodName = config.DEFAULT_SIDECAR_POD_NAME
        }

        sidecarPodImage, err := config.ReadEnv("SIDECAR_POD_IMAGE")
        
        if err != nil {
            sidecarPodImage = config.DEFAULT_SIDECAR_POD_IMAGE
        }

        sidecar := corev1.Container{
            Name:  sidecarPodName,
            Image: sidecarPodImage,

            // Add any sidecar-specific settings here
        }
        pod.Spec.Containers = append(pod.Spec.Containers, sidecar)
    }

    // Return the mutated pod
    mutatedPod, _ := json.Marshal(pod)
    return admissionv1.AdmissionResponse{
        Allowed: true,
        Patch:   mutatedPod,
    }
}