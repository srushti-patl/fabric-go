/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Customer physical Port
type PortDemarcationPoint struct {
	// Port cabinet unique space id
	CabinetUniqueSpaceId string `json:"cabinetUniqueSpaceId,omitempty"`
	// Port cage unique space id
	CageUniqueSpaceId string `json:"cageUniqueSpaceId,omitempty"`
	// Port patch panel
	PatchPanel string `json:"patchPanel,omitempty"`
	// Port patch panel
	PatchPanelName string `json:"patchPanelName,omitempty"`
	// Port patch panel port A
	PatchPanelPortA string `json:"patchPanelPortA,omitempty"`
	// Port patch panel port B
	PatchPanelPortB string `json:"patchPanelPortB,omitempty"`
	// Port connector type
	ConnectorType string `json:"connectorType,omitempty"`
	// Port ibx identifier
	Ibx string `json:"ibx,omitempty"`
	// Port reservation identifier
	PortReservationId string `json:"portReservationId,omitempty"`
	// Port group identifier
	PortGroup string `json:"portGroup,omitempty"`
	// Port identifier
	CorrelationId string `json:"correlationId,omitempty"`
}
