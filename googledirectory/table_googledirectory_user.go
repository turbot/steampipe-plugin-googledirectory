package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	admin "google.golang.org/api/admin/directory/v1"
)

//// TABLE DEFINITION

func tableGoogleDirectroryUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_user",
		Description: "Users defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryUsers,
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
			KeyColumns: plugin.AnyColumn([]string{"id", "primary_email"}),
			Hydrate:    getDirectoryUser,
		},
		Columns: []*plugin.Column{
			{
				Name:        "full_name",
				Description: "The user's full name formed by concatenating the first and last name values.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name.FullName"),
			},
			{
				Name:        "id",
				Description: "The unique ID for the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "primary_email",
				Description: "Specifies the user's primary email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "customer_id",
				Description: "The customer ID to retrieve all account users.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "Specifies user's G-Suite account creation time.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "is_admin",
				Description: "Indicates whether an user have super administrator privileges, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_delegated_admin",
				Description: "Indicates whether the user is a delegated administrator, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "suspended",
				Description: "Indicates whether an user is suspended, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "agreed_to_terms",
				Description: "Indicates whether the user has completed an initial login and accepted the Terms of Service agreement, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "archived",
				Description: "Indicates whether an user is archived, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "change_password_at_next_login",
				Description: "Indicates if the user is forced to change their password at next login.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "deletion_time",
				Description: "Specifies user's deletion time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DeletionTime").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "etag",
				Description: "A hash of the metadata, used to ensure there were no concurrent modifications to the resource when attempting an update.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "family_name",
				Description: "The user's last name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name.FamilyName"),
			},
			{
				Name:        "gender",
				Description: "The user's gender.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "given_name",
				Description: "The user's first name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name.GivenName"),
			},
			{
				Name:        "hash_function",
				Description: "Specifies the hash format of the password property.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "include_in_global_address_list",
				Description: "Indicates whether the user's profile is visible in the Google Workspace global address list when the contact sharing feature is enabled for the domain.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "ip_whitelisted",
				Description: "Indicates whether the user's IP address is whitelisted, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_enforced_in_2sv",
				Description: "Indicates whether the 2-step verification enforced, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("IsEnforcedIn2Sv"),
			},
			{
				Name:        "is_enrolled_in_2sv",
				Description: "Indicates whether an user is enrolled in 2-step verification, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("IsEnrolledIn2Sv"),
			},
			{
				Name:        "is_mailbox_setup",
				Description: "Indicates whether the user's Google mailbox is created, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "kind",
				Description: "The type of the API resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_login_time",
				Description: "Specifies user's last login time.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "org_unit_path",
				Description: "The full path of the parent organization associated with the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "recovery_email",
				Description: "Specifies the recovery email of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "recovery_phone",
				Description: "Specifies the recovery phone of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "suspension_reason",
				Description: "Specifies the reason a user account is suspended either by the administrator or by Google at the time of suspension.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "thumbnail_photo_etag",
				Description: "ETag of the user's photo.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "thumbnail_photo_url",
				Description: "Photo Url of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "query",
				Description: "Filter string to [filter](https://developers.google.com/admin-sdk/directory/v1/guides/search-users) users.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("query"),
			},
			{
				Name:        "addresses",
				Description: "A list of the user's addresses.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "aliases",
				Description: "A list of the user's alias email addresses.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "custom_schemas",
				Description: "Custom fields of the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "emails",
				Description: "A list of the user's email addresses.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "external_ids",
				Description: "A list of external IDs for the user, such as an employee or network ID.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "ims",
				Description: "The user's Instant Messenger (IM) accounts.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "keywords",
				Description: "The user's keywords.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "languages",
				Description: "The user's languages.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "locations",
				Description: "The user's locations.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "non_editable_aliases",
				Description: "A list of the user's non-editable alias email addresses.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "notes",
				Description: "Notes for the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "organizations",
				Description: "A list of organizations the user belongs to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "phones",
				Description: "A list of the user's phone numbers.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "posix_accounts",
				Description: "A list of POSIX account information for the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "relations",
				Description: "A list of the user's relationships to other users.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "ssh_public_keys",
				Description: "A list of SSH public keys.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "websites",
				Description: "The user's websites.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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

	var query string
	if d.KeyColumnQuals["query"] != nil {
		query = d.KeyColumnQuals["query"].GetStringValue()
	}

	resp := service.Users.List().Customer(customerID).Query(query)
	if err := resp.Pages(ctx, func(page *admin.Users) error {
		for _, user := range page.Users {
			d.StreamListItem(ctx, user)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

func getDirectoryUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryUser")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetStringValue()
	primaryEmail := d.KeyColumnQuals["primary_email"].GetStringValue()

	// Return nil, if no input provided
	if id == "" && primaryEmail == "" {
		return nil, nil
	}

	var inputStr string
	if id == "" {
		inputStr = primaryEmail
	} else {
		inputStr = id
	}

	resp, err := service.Users.Get(inputStr).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
