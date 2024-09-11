package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

//// TABLE DEFINITION

func tableGoogleDirectoryGroupMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_group_member",
		Description: "Group members defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			ParentHydrate: listDirectoryGroups,
			Hydrate:       listDirectoryGroupMembers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "group_id",
					Require: plugin.Optional,
				},
				{
					Name:    "role",
					Require: plugin.Optional,
				},
			},
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"group_id", "id"}),
			Hydrate:    getDirectoryGroupMember,
		},
		Columns: []*plugin.Column{
			{
				Name:        "group_id",
				Description: "Specifies the ID of the group, the user belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupId"),
			},
			{
				Name:        "id",
				Description: "The unique ID of the group member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "email",
				Description: "Specifies the member's email address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Email"),
			},
			{
				Name:        "role",
				Description: "Specifies the role of the member in a group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Role"),
			},
			{
				Name:        "status",
				Description: "Specifies the status of the member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Status"),
			},
			{
				Name:        "delivery_settings",
				Description: "Defines mail delivery preferences of member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DeliverySettings"),
			},
			{
				Name:        "etag",
				Description: "A hash of the metadata, used to ensure there were no concurrent modifications to the resource when attempting an update.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Etag"),
			},
			{
				Name:        "kind",
				Description: "The type of the API resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Kind"),
			},
			{
				Name:        "type",
				Description: "The type of group member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Type"),
			},
		},
	}
}

type MemberInfo struct {
	GroupId string
	*admin.Member
}

//// LIST FUNCTION

func listDirectoryGroupMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var group *admin.Group
	if h.Item != nil {
		group = h.Item.(*admin.Group)
	}

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	groupID := d.EqualsQuals["group_id"].GetStringValue()

	// Minimize the API call if group ID is specified in where clause
	if groupID != group.Id {
		return nil, nil
	}

	var role string
	if d.EqualsQuals["role"] != nil {
		role = d.EqualsQuals["role"].GetStringValue()
	}

	// By default, API can return maximum 200 records in a single page
	maxResult := int64(200)

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < maxResult {
			maxResult = *limit
		}
	}

	resp := service.Members.List(group.Id).Roles(role).MaxResults(maxResult)
	if err := resp.Pages(ctx, func(page *admin.Members) error {
		for _, member := range page.Members {
			d.StreamListItem(ctx, &MemberInfo{group.Id, member})

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if plugin.IsCancelled(ctx) {
				page.NextPageToken = ""
				break
			}
		}
		return nil
	}); err != nil {
		// Return nil, if given group is not present
		if err.(*googleapi.Error).Code == 404 {
			return nil, nil
		}
		return nil, err
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

func getDirectoryGroupMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryGroupMember")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	groupID := d.EqualsQuals["group_id"].GetStringValue()
	memberID := d.EqualsQuals["id"].GetStringValue()

	// Return nil, if no input provided
	if groupID == "" || memberID == "" {
		return nil, nil
	}

	resp, err := service.Members.Get(groupID, memberID).Do()
	if err != nil {
		return nil, err
	}

	return &MemberInfo{groupID, resp}, nil
}
