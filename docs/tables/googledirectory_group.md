# Table: googledirectory_group

Query information about groups defined in the Google Workspace directory.

## Examples

### Basic info

```sql
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group;
```

### Get group by ID

```sql
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group
where
  id = '02ce457p6conzyd';
```

### Get group by email

```sql
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group
where
  email = 'scranton@dundermifflin.com';
```

### List top 5 groups by member count

```sql
select
  name,
  direct_members_count
from
  googledirectory_group
order by
  direct_members_count desc
limit 5;
```

### List all groups and their members

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

### List groups using the [query filter](https://developers.google.com/admin-sdk/directory/v1/guides/search-groups)

```sql
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group
where
  query = 'email:steampipe*';
```
