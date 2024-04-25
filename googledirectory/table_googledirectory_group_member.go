package googledirectory

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/turbot/go-kit/helpers"
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
			// ParentHydrate: listDirectoryGroups,
			Hydrate: listDirectoryGroupMembers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "group_id",
					Require: plugin.Optional,
				},
				{
					Name:    "role",
					Require: plugin.Optional,
				},
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "query",
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
				Transform:   transform.FromQual("group_id"),
			},
			{
				Name:        "name",
				Description: "The group's display name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "customer_id",
				Description: "The customer ID to retrieve all account groups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("customer_id"),
			},
			{
				Name:        "query",
				Description: "Filter string to [filter](https://developers.google.com/admin-sdk/directory/v1/guides/search-groups) groups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("query"),
			},
			{
				Name:        "id",
				Description: "The unique ID of the group member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "Specifies the member's email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role",
				Description: "Specifies the role of the member in a group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "Specifies the status of the member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "delivery_settings",
				Description: "Defines mail delivery preferences of member.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getDirectoryGroupMember,
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
				Name:        "type",
				Description: "The type of group member.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryGroupMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var groups []string

	groups, err := listDirectoryGroupForMembers(ctx, d, h)
	if err != nil {
		return nil, nil
	}

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	groupID := d.EqualsQuals["group_id"].GetStringValue()

	// Minimize the API call if group ID is specified in where clause
	if helpers.StringSliceContains(groups, groupID) {
		groups = []string{groupID}
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

	// make parallel API call
	var wg sync.WaitGroup
	serviceCh := make(chan *admin.Members, 2000)
	errorCh := make(chan error, 2000)

	for _, g := range groups {
		wg.Add(1)
		go listDirectoryGroupMembersByGroupIdAsync(g, role, maxResult, service.Members, &wg, serviceCh, errorCh, ctx)
	}

	// wait for all services to be processed
	wg.Wait()

	// NOTE: close channel before ranging over results
	close(serviceCh)
	close(errorCh)

	for err := range errorCh {
		// return the first error
		return nil, err
	}

	for result := range serviceCh {
		for _, service := range result.Members {
			d.StreamListItem(ctx, service)

			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, err
}

func listDirectoryGroupMembersByGroupIdAsync(groupId string, role string, maxResult int64, client *admin.MembersService, wg *sync.WaitGroup, serviceCh chan *admin.Members, errorCh chan error, ctx context.Context) {
	defer wg.Done()
	resp := client.List(groupId).Roles(role).MaxResults(maxResult)
	if err := resp.Pages(ctx, func(page *admin.Members) error {
		serviceCh <- page
		return nil
	}); err != nil {
		// Return nil, if given group is not present
		if err.(*googleapi.Error).Code != 404 {
			errorCh <- err
		}
	}
}

func listDirectoryGroupForMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) ([]string, error) {

	var groups []string
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	var queryFilter, query string
	var filter []string

	if d.EqualsQuals["name"] != nil {
		filter = append(filter, fmt.Sprintf("name='%s'", d.EqualsQuals["name"].GetStringValue()))
	}

	if d.EqualsQuals["query"] != nil {
		queryFilter = d.EqualsQuals["query"].GetStringValue()
	}

	if queryFilter != "" {
		query = queryFilter
	} else if len(filter) > 0 {
		query = strings.Join(filter, " ")
	}


	// Since, query parameter can't be empty, set default param name:**, to return all groups
	if query == "" {
		query = "name:**"
	}

	// Set default value to my_customer, to represent current account
	customerID := "my_customer"
	if d.EqualsQuals["customer_id"] != nil {
		customerID = d.EqualsQuals["customer_id"].GetStringValue()
	}

	// By default, API can return maximum 200 records in a single page
	maxResult := int64(200)

	resp := service.Groups.List().Customer(customerID).Query(query).MaxResults(maxResult)
	if err := resp.Pages(ctx, func(page *admin.Groups) error {
		for _, g := range page.Groups {
			groups = append(groups, g.Id)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return groups, nil
}

//// HYDRATE FUNCTIONS

func getDirectoryGroupMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryGroupMember")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	var groupID, memberID string
	if h.Item != nil {
		data := h.Item.(*admin.Member)
		groupID = d.EqualsQuals["group_id"].GetStringValue()
		memberID = data.Id
	} else {
		groupID = d.EqualsQuals["group_id"].GetStringValue()
		memberID = d.EqualsQuals["id"].GetStringValue()
	}

	// Return nil, if no input provided
	if groupID == "" || memberID == "" {
		return nil, nil
	}

	resp, err := service.Members.Get(groupID, memberID).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
