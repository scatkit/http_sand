alter table users
rename column name to user_name;

alter table users
add column is_edited boolean;

alter table users
drop column password;


