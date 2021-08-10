# Table: googledirectory_group_member

Query information about group members defined in the Google Workspace directory.

**Note:**

- A specific `group_id` must be defined in all queries to this table.

## Examples

### Basic info

```sql
select
  group_id,
  id,
  email,
  role
from
  googledirectory_group_member
where
  group_id = '01ksv4uv1gexk1h';
```

### List all owners of a group

```sql
select
  group_id,
  id,
  email,
  role
from
  googledirectory_group_member
where
  group_id = '01ksv4uv1gexk1h'
  and role = 'OWNER';
```

### Role composition for a group

```sql
select
  role,
  count(*)
from
  googledirectory_group_member
where
  group_id = '01ksv4uv1gexk1h'
group by role
order by
  count desc;
```

### List all groups with all members

```sql
select
  g.id as group_id,
  g.name as group_name,
  m.email as member_email
from
  googledirectory_group as g,
  googledirectory_group_member as m
where
  g.id = m.group_id
order by
  g.name,
  m.email;
```
