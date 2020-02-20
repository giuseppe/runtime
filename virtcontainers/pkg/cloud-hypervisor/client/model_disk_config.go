/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// DiskConfig struct for DiskConfig
type DiskConfig struct {
	Path string `json:"path"`
	Readonly bool `json:"readonly,omitempty"`
	Iommu bool `json:"iommu,omitempty"`
	NumQueues int32 `json:"num_queues,omitempty"`
	QueueSize int32 `json:"queue_size,omitempty"`
	VhostUser bool `json:"vhost_user,omitempty"`
	VhostSocket string `json:"vhost_socket,omitempty"`
	Wce bool `json:"wce,omitempty"`
}
