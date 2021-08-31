# Table: googledirectory_role_assignment

Query information about role assignments defined in the Google Workspace directory.

## Examples

### Basic info

```sql
select
  role_assignment_id,
  role_id,
  assigned_to,
  scope_type
from
  googledirectory_role_assignment;
```

### Get role assignments by role ID

```sql
select
  role_assignment_id,
  role_id,
  assigned_to,
  scope_type
from
  googledirectory_role_assignment
where
  role_id = '522363132560015';
```

### Get role assignments by user

```sql
select
  assigned_role.role_assignment_id as role_assignment_id,
  r.role_name as role_name,
  u.full_name as user_name
from
  googledirectory_role_assignment as assigned_role,
  googledirectory_user as u,
  googledirectory_role as r
where
  assigned_role.user_key = u.id
  and assigned_role.role_id = r.role_id;
```
