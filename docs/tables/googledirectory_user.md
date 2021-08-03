# Table: googledirectory_user

Query information about users defined in the Google Workspace.

## Examples

### Basic info

```sql
select
  full_name,
  id,
  primary_email,
  creation_time,
  is_delegated_admin,
  customer_id
from
  googledirectory_user;
```

### Get user by ID

```sql
select
  full_name,
  id,
  primary_email,
  creation_time,
  is_delegated_admin,
  customer_id
from
  googledirectory_user
where
  id = '119982672925259996273';
```

### Get user by primary email

```sql
select
  full_name,
  id,
  primary_email,
  creation_time,
  is_delegated_admin,
  customer_id
from
  googledirectory_user
where
  primary_email = 'mscott@dundermifflin.com';
```

### List administrators

```sql
select
  id,
  full_name,
  primary_email,
  is_admin,
  is_delegated_admin
from
  googledirectory_user
where
  is_admin
  or is_delegated_admin;
```

### Users without two step verification

```sql
select
  id,
  full_name,
  primary_email,
  is_enrolled_in_2sv,
  is_enforced_in_2sv
from
  googledirectory_user
where
  not is_enrolled_in_2sv
  or not is_enforced_in_2sv;
```

### Users who have not logged in for more than 30 days

```sql
select
  id,
  full_name,
  primary_email,
  last_login_time
from
  googledirectory_user
where
  last_login_time < current_timestamp - interval '30 days';
```
