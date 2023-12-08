/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// EPT service instance
type PrecisionTimeServiceRequest struct {
	// Indicate the entity is EPT service
	Type_                string                 `json:"type"`
	Name                 string                 `json:"name"`
	Description          string                 `json:"description,omitempty"`
	Package_             *PackageRequest        `json:"package"`
	Connections          []FabricConnectionUuid `json:"connections"`
	NetworkingIpv4       *Ipv4                  `json:"networkingIpv4"`
	AdvanceConfiguration *AdvanceConfiguration  `json:"advanceConfiguration,omitempty"`
	Project              *Project               `json:"project,omitempty"`
}
