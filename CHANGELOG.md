## v0.3.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#35](https://github.com/turbot/steampipe-plugin-googledirectory/pull/35))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#34](https://github.com/turbot/steampipe-plugin-googledirectory/pull/34))

## v0.2.1 [2022-04-14]

_Bug fixes_

- Fixed links in documentation for configuring OAuth client authentication.

## v0.2.0 [2022-01-31]

_What's new?_

- Added: The `credentials` argument can now be specified in the configuration file to pass in either the path to or the contents of a service account key file in JSON format ([#32](https://github.com/turbot/steampipe-plugin-googledirectory/pull/32))
- Added: The `token_path` argument can now be specified in the configuration file to authenticate using OAuth 2.0 ([#32](https://github.com/turbot/steampipe-plugin-googledirectory/pull/32))

_Deprecated_

- The `credential_file` argument in the configuration file is now deprecated and will be removed in the next major version. We recommend using the `credentials` argument instead, which can take the same file path as the `credential_file` argument. ([#32](https://github.com/turbot/steampipe-plugin-googledirectory/pull/32))

## v0.1.0 [2021-12-08]

_Enhancements_

- Recompiled plugin with Go version 1.17 ([#28](https://github.com/turbot/steampipe-plugin-googledirectory/pull/28))
- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#27](https://github.com/turbot/steampipe-plugin-googledirectory/pull/27))

## v0.0.4 [2021-10-20]

_Bug fixes_

- Fixed: All tables now return the service API disabled error directly instead of returning empty rows

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
