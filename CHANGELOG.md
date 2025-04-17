## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#109](https://github.com/turbot/steampipe-plugin-googledirectory/pull/109))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#109](https://github.com/turbot/steampipe-plugin-googledirectory/pull/109))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#103](https://github.com/turbot/steampipe-plugin-googledirectory/pull/103))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#103](https://github.com/turbot/steampipe-plugin-googledirectory/pull/103))

## v0.8.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#85](https://github.com/turbot/steampipe-plugin-googledirectory/pull/85))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#85](https://github.com/turbot/steampipe-plugin-googledirectory/pull/85))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-googledirectory/blob/main/docs/LICENSE). ([#85](https://github.com/turbot/steampipe-plugin-googledirectory/pull/85))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#84](https://github.com/turbot/steampipe-plugin-googledirectory/pull/84))

## v0.7.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#64](https://github.com/turbot/steampipe-plugin-googledirectory/pull/64))

## v0.7.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#60](https://github.com/turbot/steampipe-plugin-googledirectory/pull/60))
- Recompiled plugin with Go version `1.21`. ([#60](https://github.com/turbot/steampipe-plugin-googledirectory/pull/60))

## v0.6.0 [2023-08-31]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v551-2023-07-26). ([#50](https://github.com/turbot/steampipe-plugin-googledirectory/pull/50))
- Recompiled plugin with `google.golang.org/api v0.138.0`. ([#52](https://github.com/turbot/steampipe-plugin-googledirectory/pull/52))
- Recompiled plugin with `github.com/aws/aws-sdk-go v1.34.0`. ([#47](https://github.com/turbot/steampipe-plugin-googledirectory/pull/47))
- Recompiled plugin with `golang.org/x/net v0.7.0`. ([#49](https://github.com/turbot/steampipe-plugin-googledirectory/pull/49))
- Recompiled plugin with `github.com/turbot/go-kit v0.7.0`. ([#51](https://github.com/turbot/steampipe-plugin-googledirectory/pull/51))

## v0.5.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#44](https://github.com/turbot/steampipe-plugin-googledirectory/pull/44))

## v0.4.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#39](https://github.com/turbot/steampipe-plugin-googledirectory/pull/39))
- Recompiled plugin with Go version `1.19`. ([#39](https://github.com/turbot/steampipe-plugin-googledirectory/pull/39))

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
