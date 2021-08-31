package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	admin "google.golang.org/api/admin/directory/v1"
)

//// TABLE DEFINITION

func tableGoogleDirectoryRoleAssignment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_role_assignment",
		Description: "Role Assignments defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryRoleAssignments,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
				{
					Name:    "role_id",
					Require: plugin.Optional,
				},
				{
					Name:    "user_key",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("role_assignment_id"),
			Hydrate:    getDirectoryRoleAssignment,
		},
		Columns: []*plugin.Column{
			{
				Name:        "role_assignment_id",
				Description: "The unique ID for the role assignment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role_id",
				Description: "The unique ID for the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "assigned_to",
				Description: "The unique ID of the user this role is assigned to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "customer_id",
				Description: "The customer ID to retrieve all account roles.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("customer_id"),
			},
			{
				Name:        "user_key",
				Description: "The user's primary email address, alias email address, or unique user ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("user_key"),
			},
			{
				Name:        "etag",
				Description: "A hash of the metadata, used to ensure there were no concurrent modifications to the resource when attempting an update.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kind",
				Description: "The type of the API resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "org_unit_id",
				Description: "If the role is restricted to an organization unit, this contains the ID for the organization unit the exercise of this role is restricted to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "scope_type",
				Description: "The scope in which this role is assigned.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryRoleAssignments(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Set default value to my_customer, to represent current account
	customerID := "my_customer"
	if d.KeyColumnQuals["customer_id"] != nil {
		customerID = d.KeyColumnQuals["customer_id"].GetStringValue()
	}

	var roleId string
	if d.KeyColumnQuals["role_id"] != nil {
		roleId = d.KeyColumnQuals["role_id"].GetStringValue()
	}

	resp := service.RoleAssignments.List(customerID).RoleId(roleId)
	if d.KeyColumnQuals["user_key"] != nil {
		resp.UserKey(d.KeyColumnQuals["user_key"].GetStringValue())
	}
	if err := resp.Pages(ctx, func(page *admin.RoleAssignments) error {
		for _, assignment := range page.Items {
			d.StreamListItem(ctx, assignment)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

func getDirectoryRoleAssignment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryRoleAssignment")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	roleAssignmentId := d.KeyColumnQuals["role_assignment_id"].GetStringValue()

	// Return nil, if no input provided
	if roleAssignmentId == "" {
		return nil, nil
	}

	resp, err := service.RoleAssignments.Get("my_customer", roleAssignmentId).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
