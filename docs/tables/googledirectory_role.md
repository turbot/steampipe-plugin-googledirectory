# Table: googledirectory_role

Query information about roles defined in the Google Workspace directory.

## Examples

### Basic info

```sql
select
  role_name,
  role_id,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role;
```

### Get role by ID

```sql
select
  role_name,
  role_id,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role
where
  role_id = '02ce457p6conzyd';
```

### List super admin roles

```sql
select
  role_id,
  role_name,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role
where
  is_super_admin_role;
```

### List system roles

```sql
select
  role_id,
  role_name,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role
where
  is_system_role;
```

### List privileges by role

```sql
select
  role_name,
  p ->> 'serviceId' as service_id,
  p ->> 'privilegeName' as privilege
from
  googledirectory_role as r,
  jsonb_array_elements(r.role_privileges) as p
order by
  role_name,
  service_id,
  privilege;
```
