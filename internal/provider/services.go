package provider

import (
	"terraform-provider-groupmanager/internal/services/groups"
	"terraform-provider-groupmanager/internal/services/serviceprincipals"
	"terraform-provider-groupmanager/internal/services/users"
)

func SupportedServices() []ServiceRegistration {
	return []ServiceRegistration{
		groups.Registration{},
		serviceprincipals.Registration{},
		users.Registration{},
	}
}
