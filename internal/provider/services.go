package provider

import (
	"github.com/teswinfrastructure/terraform-provider-groupmanager/internal/services/groups"
	"github.com/teswinfrastructure/terraform-provider-groupmanager/internal/services/serviceprincipals"
	"github.com/teswinfrastructure/terraform-provider-groupmanager/internal/services/users"
)

func SupportedServices() []ServiceRegistration {
	return []ServiceRegistration{
		groups.Registration{},
		serviceprincipals.Registration{},
		users.Registration{},
	}
}
