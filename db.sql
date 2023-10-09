create database todo_List;

create table tasks (
                       id serial primary key ,
                       title varchar(100) not null ,
                       description varchar(200),
                       status boolean default 'true',
                       created_at timestamp not null  default (now() at time zone 'utc-5')
);

insert into tasks (title, description)
values ('Разработка проекта To-Do List', 'написал моделки проекта');

select *from tasks;
drop table tasks;

