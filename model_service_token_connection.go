/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.3
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Service Token Connection Type Information
type ServiceTokenConnection struct {
	// Type of Connection
	Type_ string `json:"type"`
	// Authorization to connect remotely
	AllowRemoteConnection bool `json:"allowRemoteConnection,omitempty"`
	// Connection bandwidth limit in Mbps
	BandwidthLimit int32 `json:"bandwidthLimit,omitempty"`
	// List of permitted bandwidths.
	SupportedBandwidths []int32           `json:"supportedBandwidths,omitempty"`
	ASide               *ServiceTokenSide `json:"aSide,omitempty"`
	ZSide               *ServiceTokenSide `json:"zSide,omitempty"`
}
