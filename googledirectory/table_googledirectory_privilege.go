package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableGoogleDirectoryPrivilege(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_privilege",
		Description: "Privileges defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryPrivileges,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
			},
			ShouldIgnoreError: isNotFoundError([]string{"403", "404"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "privilege_name",
				Description: "The name of the privilege.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "service_name",
				Description: "The name of the service this privilege is for.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "service_id",
				Description: "The obfuscated ID of the service this privilege is for.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_ou_scopable",
				Description: "Indicates if the privilege can be restricted to an organization unit.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "customer_id",
				Description: "The customer ID to retrieve all privileges for a customer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("customer_id"),
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
				Name:        "child_privileges",
				Description: "A list of child privileges. Privileges for a service form a tree. Each privilege can have a list of child privileges; this list is empty for a leaf privilege.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryPrivileges(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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

	resp, err := service.Privileges.List(customerID).Do()
	if err != nil {
		return nil, err
	}

	for _, role := range resp.Items {
		d.StreamListItem(ctx, role)
	}

	return nil, err
}
