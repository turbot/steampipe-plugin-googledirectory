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

### Get role assigments by ID

```sql
select
  role_assignment_id,
  role_id,
  assigned_to,
  scope_type
from
  googledirectory_role_assignment
where
  role_id = '02ce457p6conzyd';
```
