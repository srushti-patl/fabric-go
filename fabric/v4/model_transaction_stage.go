/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

import (
	"time"
)

type TransactionStage struct {
	// for internal users
	External string `json:"external,omitempty"`
	// stage type
	Type_ string `json:"type,omitempty"`
	// state
	State string `json:"state,omitempty"`
	// Created by Date and Time
	InitiatedDateTime time.Time `json:"initiatedDateTime,omitempty"`
	// Created by Date and Time
	CompletedDateTime time.Time `json:"completedDateTime,omitempty"`
	// Connection name
	Duration string `json:"duration,omitempty"`
	// for internal users
	DurationSlo string `json:"durationSlo,omitempty"`
}
