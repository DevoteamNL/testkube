/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// pod request body
type PodRequest struct {
	Resources *PodResourcesRequest `json:"resources,omitempty"`
	// pod template extensions
	PodTemplate string `json:"podTemplate,omitempty"`
	// name of the template resource
	PodTemplateReference string `json:"podTemplateReference,omitempty"`
}
