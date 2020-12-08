package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type ServerExtendedAuditingPolicyId struct {
	SubscriptionId              string
	ResourceGroup               string
	ServerName                  string
	ExtendedAuditingSettingName string
}

func NewServerExtendedAuditingPolicyID(subscriptionId, resourceGroup, serverName, extendedAuditingSettingName string) ServerExtendedAuditingPolicyId {
	return ServerExtendedAuditingPolicyId{
		SubscriptionId:              subscriptionId,
		ResourceGroup:               resourceGroup,
		ServerName:                  serverName,
		ExtendedAuditingSettingName: extendedAuditingSettingName,
	}
}

func (id ServerExtendedAuditingPolicyId) String() string {
	segments := []string{
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
		fmt.Sprintf("Server Name %q", id.ServerName),
		fmt.Sprintf("Extended Auditing Setting Name %q", id.ExtendedAuditingSettingName),
	}
	return strings.Join(segments, " / ")
}

func (id ServerExtendedAuditingPolicyId) ID(_ string) string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/extendedAuditingSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ServerName, id.ExtendedAuditingSettingName)
}

// ServerExtendedAuditingPolicyID parses a ServerExtendedAuditingPolicy ID into an ServerExtendedAuditingPolicyId struct
func ServerExtendedAuditingPolicyID(input string) (*ServerExtendedAuditingPolicyId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ServerExtendedAuditingPolicyId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.ServerName, err = id.PopSegment("servers"); err != nil {
		return nil, err
	}
	if resourceId.ExtendedAuditingSettingName, err = id.PopSegment("extendedAuditingSettings"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}