# Table: googledirectory_domain_alias

Query information about domain aliases defined in the Google Workspace directory.

## Examples

### Basic info

```sql
select
  domain_alias_name,
  creation_time,
  verified
from
  googledirectory_domain_alias;
```

### List unverified domain aliases

```sql
select
  domain_alias_name,
  creation_time,
  verified
from
  googledirectory_domain_alias
where
  not verified;
```

### List domain aliases by parent domain

```sql
select
  domain_alias_name,
  parent_domain_name,
  creation_time,
  verified
from
  googledirectory_domain_alias
where
  parent_domain_name = 'domain.com';
```
