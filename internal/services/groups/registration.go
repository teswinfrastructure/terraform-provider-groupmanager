package groups

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Groups"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"Groups",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"groupmanager_group": groupDataSource(),
		// "groupmanager_groups": groupsDataSource(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"groupmanager_group": groupResource(),
		// "groupmanager_group_member": groupMemberResource(),
	}
}
