package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	admin "google.golang.org/api/admin/directory/v1"
)

//// TABLE DEFINITION

func tableGoogleDirectoryGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_group",
		Description: "Groups defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryGroups,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
				{
					Name:    "query",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "email"}),
			Hydrate:    getDirectoryGroup,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The group's display name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The unique ID of a group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "Specifies the group's email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "admin_created",
				Description: "Indicates whether the group is created by an administrator, or by an user.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "customer_id",
				Description: "The customer ID to retrieve all account groups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("customer_id"),
			},
			{
				Name:        "description",
				Description: "An extended description to help users determine the purpose of a group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "direct_members_count",
				Description: "The number of users that are direct members of the group.",
				Type:        proto.ColumnType_INT,
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
				Name:        "query",
				Description: "Filter string to [filter](https://developers.google.com/admin-sdk/directory/v1/guides/search-groups) groups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("query"),
			},
			{
				Name:        "aliases",
				Description: "A list of the group's alias email addresses.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "non_editable_aliases",
				Description: "A list of the group's non-editable alias email addresses that are outside of the account's primary domain or subdomains.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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

	query := d.KeyColumnQuals["query"].GetStringValue()
	// Since, query parameter can't be empty, set default param name:**, to return all groups
	if query == "" {
		query = "name:**"
	}

	resp := service.Groups.List().Customer(customerID).Query(query)
	if err := resp.Pages(ctx, func(page *admin.Groups) error {
		for _, group := range page.Groups {
			d.StreamListItem(ctx, group)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

func getDirectoryGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryGroup")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetStringValue()
	email := d.KeyColumnQuals["email"].GetStringValue()

	// Return nil, if no input provided
	if id == "" && email == "" {
		return nil, nil
	}

	var inputStr string
	if id == "" {
		inputStr = email
	} else {
		inputStr = id
	}

	resp, err := service.Groups.Get(inputStr).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
