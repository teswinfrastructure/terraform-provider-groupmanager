package client

import (
	"terraform-provider-groupmanager/internal/common"

	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AdministrativeUnitsClient *msgraph.AdministrativeUnitsClient
	DirectoryObjectsClient    *msgraph.DirectoryObjectsClient
	GroupsClient              *msgraph.GroupsClient
}

func NewClient(o *common.ClientOptions) *Client {
	administrativeUnitsClient := msgraph.NewAdministrativeUnitsClient()
	o.ConfigureClient(&administrativeUnitsClient.BaseClient)

	// SDK uses wrong endpoint for v1.0 API, see https://github.com/manicminer/hamilton/issues/222
	administrativeUnitsClient.BaseClient.ApiVersion = msgraph.VersionBeta

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient()
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	groupsClient := msgraph.NewGroupsClient()
	o.ConfigureClient(&groupsClient.BaseClient)

	groupsClient.BaseClient.ApiVersion = msgraph.VersionBeta

	return &Client{
		AdministrativeUnitsClient: administrativeUnitsClient,
		DirectoryObjectsClient:    directoryObjectsClient,
		GroupsClient:              groupsClient,
	}
}
