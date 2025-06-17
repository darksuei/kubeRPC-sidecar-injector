package admission

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	config "github.com/darksuei/kubeRPC-sidecar-injector/config"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Mutate(req admissionv1.AdmissionReview) (*admissionv1.AdmissionResponse, error) {
    var pod corev1.Pod

    if err := json.Unmarshal(req.Request.Object.Raw, &pod); err != nil {
        return nil, errors.New("failed to unmarshal raw request object")
    }

    inject, _ := config.ReadEnv("ANNOTATION_INJECT", config.DEFAULT_ANNOTATION_INJECT)

    // Check if the annotation exists
    if pod.Annotations[inject] != "true" {
        return &admissionv1.AdmissionResponse{
            UID:     req.Request.UID,
            Allowed: true,
        }, nil
    }
    
    name, _ := config.ReadEnv("SIDECAR_POD_NAME", config.DEFAULT_SIDECAR_POD_NAME)
    image, _ := config.ReadEnv("SIDECAR_POD_IMAGE", config.DEFAULT_SIDECAR_POD_IMAGE)
    port, _ := config.ReadEnv("SIDECAR_POD_PORT", config.DEFAULT_SIDECAR_POD_PORT)

    portInt64, err := strconv.ParseInt(port, 10, 32)
    if err != nil {
        return AdmissionReviewError(req, errors.New("failed to parse port")), nil
    }

    portInt32 := int32(portInt64)

    // Build the JSON‑patch
    patch := []map[string]interface{}{
        {
            "op":   "add",
            "path": "/spec/containers/-",
            "value": corev1.Container{
                Name:  name,
                Image: image,
                Ports: []corev1.ContainerPort{
                    {
                        ContainerPort: portInt32,
                        Name:          name + "-service",
                        Protocol:      corev1.ProtocolTCP,
                    },
                },
                ImagePullPolicy: corev1.PullIfNotPresent,
                // Todo: Missing env, liveness probe & readiness probe, resource limits.
            },
        },
    }
    patchBytes, _ := json.Marshal(patch)

    pt := admissionv1.PatchTypeJSONPatch

    return &admissionv1.AdmissionResponse{
        UID:       req.Request.UID,
        Allowed:   true,
        Patch:     patchBytes,
        PatchType: &pt,
    }, nil
}

func AdmissionReviewError(req admissionv1.AdmissionReview, err error) *admissionv1.AdmissionResponse {
    return &admissionv1.AdmissionResponse{
        UID:     req.Request.UID,
        Allowed: false,
        Result: &metav1.Status{
            Code:    http.StatusInternalServerError,
            Message: err.Error(),
        },
    }
}