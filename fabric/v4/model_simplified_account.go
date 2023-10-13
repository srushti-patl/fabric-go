/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.10
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

type SimplifiedAccount struct {
	// Account number
	AccountNumber int64 `json:"accountNumber,omitempty"`
	// Account name
	AccountName string `json:"accountName,omitempty"`
	// Customer organization identifier
	OrgId int64 `json:"orgId,omitempty"`
	// Customer organization name
	OrganizationName string `json:"organizationName,omitempty"`
	// Global organization identifier
	GlobalOrgId string `json:"globalOrgId,omitempty"`
	// Global organization name
	GlobalOrganizationName string `json:"globalOrganizationName,omitempty"`
	// Account ucmId
	UcmId string `json:"ucmId,omitempty"`
	// Account name
	GlobalCustId string `json:"globalCustId,omitempty"`
	// Reseller account number
	ResellerAccountNumber int64 `json:"resellerAccountNumber,omitempty"`
	// Reseller account name
	ResellerAccountName string `json:"resellerAccountName,omitempty"`
	// Reseller account ucmId
	ResellerUcmId string `json:"resellerUcmId,omitempty"`
	// Reseller customer organization identifier
	ResellerOrgId int64 `json:"resellerOrgId,omitempty"`
}
