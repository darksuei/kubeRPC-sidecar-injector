package handlers

import (
	"net/http"

	internalAdmission "github.com/darksuei/kubeRPC-sidecar-injector/internal/domain/admission"
	"github.com/darksuei/kubeRPC-sidecar-injector/internal/infrastructure/helpers"
	"github.com/gin-gonic/gin"
	admissionv1 "k8s.io/api/admission/v1"
)

func Mutate(c *gin.Context) {
	admissionReview, err := helpers.ValidateRequest[admissionv1.AdmissionReview](c)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	response, err := internalAdmission.Mutate(*admissionReview)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	admissionReviewResponse := admissionv1.AdmissionReview{
		TypeMeta: admissionReview.TypeMeta,
		Response: response,
	}

	admissionReviewResponse.Response.UID = admissionReview.Request.UID

	c.JSON(http.StatusOK, admissionReviewResponse)
}