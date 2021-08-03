# Table: googledirectory_domain

Query information about domains defined in the Google Workspace directory.

## Examples

### Basic info

```sql
select
  domain_name,
  creation_time,
  is_primary
from
  googledirectory_domain;
```

### List unverified domains

```sql
select
  domain_name,
  creation_time,
  verified
from
  googledirectory_domain
where
  not verified;
```
