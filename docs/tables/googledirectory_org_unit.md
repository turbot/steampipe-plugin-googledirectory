# Table: googledirectory_org_unit

Query information about organization units defined in the Google Workspace directory.

## Examples

### Basic info

```sql
select
  name,
  org_unit_id,
  org_unit_path,
  description
from
  googledirectory_org_unit;
```

### Get org unit by ID

```sql
select
  name,
  org_unit_id,
  org_unit_path,
  description
from
  googledirectory_org_unit
where
  org_unit_id = 'id:03pk8a4z4t34g1w';
```

### Get org unit by path

```sql
select
  name,
  org_unit_id,
  org_unit_path,
  description
from
  googledirectory_org_unit
where
  org_unit_path = '/DM';
```
