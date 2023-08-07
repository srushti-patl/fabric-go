/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.9
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// ServiceProfileStateEnum : Equinix assigned state.
type ServiceProfileStateEnum string

// List of ServiceProfileStateEnum
const (
	ACTIVE_ServiceProfileStateEnum           ServiceProfileStateEnum = "ACTIVE"
	PENDING_APPROVAL_ServiceProfileStateEnum ServiceProfileStateEnum = "PENDING_APPROVAL"
	DELETED_ServiceProfileStateEnum          ServiceProfileStateEnum = "DELETED"
	REJECTED_ServiceProfileStateEnum         ServiceProfileStateEnum = "REJECTED"
)
