---
subcategory: "Groups"
---

# Resource: groupmanager_group

Manages only owned group within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `Group.Create`, `GroupMembers.Read.All` and `User.Read.All`

## Example Usage

*Basic example*

```terraform
data "groupmanager_client_config" "current" {}

resource "groupmanager_group" "example" {
  display_name     = "example"
  owners           = [data.groupmanager_client_config.current.object_id]
  security_enabled = true
}
```

*Group with dynamic membership*

```terraform
data "groupmanager_client_config" "current" {}

resource "groupmanager_group" "example" {
  display_name     = "MyGroup"
  owners           = [data.groupmanager_client_config.current.object_id]
  security_enabled = true
  types            = ["DynamicMembership"]

  dynamic_membership {
    enabled = true
    rule    = "user.department -eq \"IT\""
  }
}
```

## Argument Reference

The following arguments are supported:

* `administrative_unit_ids` - (Optional) The object IDs of administrative units in which the group is a member. If specified, new groups will be created in the scope of the first administrative unit and added to the others. If empty, new groups will be created at the tenant level.
* `behaviors` - (Optional) A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeMembersToCalendarEventsDisabled`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
* `description` - (Optional) The description for the group.
* `display_name` - (Required) The display name for the group.
* `dynamic_membership` - (Optional) A `dynamic_membership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
* `mail_enabled` - (Optional) Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
* `mail_nickname` - (Optional) The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
* `members` - (Optional) A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamic_membership` block.
* `owners` - (Optional) A set of object IDs of principals that will be granted ownership of the group. Supported object types are users or service principals. By default, the principal being used to execute Terraform is assigned as the sole owner. Groups cannot be created with no owners or have all their owners removed.

-> **Group Ownership**  It's recommended to always specify one or more group owners, including the principal being used to execute Terraform, such as in the example above. When removing group owners, if a user principal has been assigned ownership, the last user cannot be removed as an owner. Microsoft 365 groups are required to always have at least one owner which *must be a user* (i.e. not a service principal).

* `prevent_duplicate_names` - (Optional) If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
* `provisioning_options` - (Optional) A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
* `security_enabled` - (Optional) Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A Microsoft 365 group can be security enabled *and* mail enabled (see the `types` property).
* `theme` - (Optional) The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
* `types` - (Optional) A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mail_enabled` is true. Changing this forces a new resource to be created.

-> **Supported Group Types** At present, only security groups and Microsoft 365 groups can be created or managed with this resource. Distribution groups and mail-enabled security groups are not supported. Microsoft 365 groups can be security-enabled.

* `visibility` - (Optional) The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.

-> **Group Name Uniqueness** Group names are not unique within Azure Active Directory. Use the `prevent_duplicate_names` argument to check for existing groups if you want to avoid name collisions.

---

`dynamic_membership` block supports the following:

* `enabled` - (Required) Whether rule processing is "On" (true) or "Paused" (false).
* `rule` - (Required) The rule that determines membership of this group. For more information, see official documentation on [membership rules syntax](https://docs.microsoft.com/en-gb/azure/active-directory/enterprise-users/groups-dynamic-membership).

~> **Dynamic Group Memberships** Remember to include `DynamicMembership` in the set of `types` for the group when configuring a dynamic membership rule. Dynamic membership is a premium feature which requires an Azure Active Directory P1 or P2 license.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `mail` - The SMTP address for the group.
* `object_id` - The object ID of the group.
* `onpremises_domain_name` - The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_netbios_name` - The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_sam_account_name` - The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_security_identifier` - The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
* `onpremises_sync_enabled` - Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
* `preferred_language` - The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
* `proxy_addresses` - List of email addresses for the group that direct to the same group mailbox.

## Import

Groups can be imported using their object ID, e.g.

```shell
terraform import groupmanager_group.my_group 00000000-0000-0000-0000-000000000000
```
