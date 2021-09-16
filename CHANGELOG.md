## v0.0.3 [2021-09-16]

_What's new?_

- Added: Additional optional key columns and better filtering capabilities to `googledirectory_group`, `googledirectory_group_member`, and `googledirectory_user` tables ([#20](https://github.com/turbot/steampipe-plugin-googledirectory/pull/20))

_Enhancements_

- Updated: Improve context cancellation handling in all tables ([#20](https://github.com/turbot/steampipe-plugin-googledirectory/pull/20))

_Bug fixes_

- Fixed: Remove check for credentials in `GOOGLE_APPLICATION_CREDENTIALS` environment variable to align with Google's authentication methods ([#20](https://github.com/turbot/steampipe-plugin-googledirectory/pull/20))

## v0.0.2 [2021-09-01]

_What's new?_

- New tables added
  - [googledirectory_domain_alias](https://hub.steampipe.io/plugins/turbot/googledirectory/tables/googledirectory_domain_alias) ([#11](https://github.com/turbot/steampipe-plugin-googledirectory/pull/11))
  - [googledirectory_role_assignment](https://hub.steampipe.io/plugins/turbot/googledirectory/tables/googledirectory_role_assignment) ([#12](https://github.com/turbot/steampipe-plugin-googledirectory/pull/12))

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.5.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v150--2021-08-06) and `google.golang.org/api v0.54.0` ([#17](https://github.com/turbot/steampipe-plugin-googledirectory/pull/17))

_Bug fixes_

- Fixed typos in all table function names ([#6](https://github.com/turbot/steampipe-plugin-googledirectory/pull/6))

## v0.0.1 [2021-08-12]

_What's new?_

- New tables added

  - [googledirectory_domain](https://hub.steampipe.io/plugins/turbot/googledirectory/tables/googledirectory_domain)
  - [googledirectory_group](https://hub.steampipe.io/plugins/turbot/googledirectory/tables/googledirectory_group)
  - [googledirectory_group_member](https://hub.steampipe.io/plugins/turbot/googledirectory/tables/googledirectory_group_member)
  - [googledirectory_org_unit](https://hub.steampipe.io/plugins/turbot/googledirectory/tables/googledirectory_org_unit)
  - [googledirectory_privilege](https://hub.steampipe.io/plugins/turbot/googledirectory/tables/googledirectory_privilege)
  - [googledirectory_role](https://hub.steampipe.io/plugins/turbot/googledirectory/tables/googledirectory_role)
  - [googledirectory_user](https://hub.steampipe.io/plugins/turbot/googledirectory/tables/googledirectory_user)
