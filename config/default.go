package config

var DEFAULT_PORT string = "8080"

var DEFAULT_ANNOTATION_INJECT string = "kuberpc.suei.dev/inject"
var DEFAULT_ANNOTATION_APP_ID string = "kuberpc.suei.dev/app-id"
var DEFAULT_ANNOTATION_APP_PORT string = "kuberpc.suei.dev/app-port"

var DEFAULT_SIDECAR_POD_NAME string = "kuberpc-sidecar"
var DEFAULT_SIDECAR_POD_PORT string = "3500"
var DEFAULT_SIDECAR_POD_IMAGE string = "docker.io/darksueii/kuberpc-sidecar:0.0.0"