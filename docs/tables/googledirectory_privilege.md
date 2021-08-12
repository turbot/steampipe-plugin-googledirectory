# Table: googledirectory_privilege

Query information about privileges defined in the Google Workspace directory.

## Examples

### Basic info

```sql
select
  privilege_name,
  service_name,
  service_id,
  is_ou_scopable
from
  googledirectory_privilege;
```

### List privileges by service

```sql
select
  service_name,
  count(*)
from
  googledirectory_privilege
group by
  service_name
order by
  count desc;
```

### List privileges for each role

```sql
select
  r.role_name as role_name,
  p.service_name as service_name,
  p.privilege_name as privilege_name
from
  googledirectory_role as r,
  jsonb_array_elements(r.role_privileges) as rp,
  googledirectory_privilege as p
where
  rp ->> 'serviceId' = p.service_id
  and rp ->> 'privilegeName' = p.privilege_name
order by
  role_name,
  service_name,
  privilege_name;
```
