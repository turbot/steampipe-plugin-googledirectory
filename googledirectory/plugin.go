/*
Package googledirectory implements a steampipe plugin for googledirectory.

This plugin provides data that Steampipe uses to present foreign
tables that represent Google Directory resources.
*/
package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-googledirectory"

// Plugin creates this (googledirectory) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"googledirectory_domain":          tableGoogleDirectoryDomain(ctx),
			"googledirectory_domain_alias":    tableGoogleDirectoryDomainAlias(ctx),
			"googledirectory_group":           tableGoogleDirectoryGroup(ctx),
			"googledirectory_group_member":    tableGoogleDirectoryGroupMember(ctx),
			"googledirectory_org_unit":        tableGoogleDirectoryOrgUnit(ctx),
			"googledirectory_privilege":       tableGoogleDirectoryPrivilege(ctx),
			"googledirectory_role":            tableGoogleDirectoryRole(ctx),
			"googledirectory_role_assignment": tableGoogleDirectoryRoleAssignment(ctx),
			"googledirectory_user":            tableGoogleDirectoryUser(ctx),
		},
	}

	return p
}
