drop table if exists Users;

create table if not exists Users(
	Username varchar(255) primary key,
	Password varchar(255) not null,
	FirstName varchar(255),
	LastName varchar(255),
	Email varchar(255),
	CreatedAt timestamp default current_timestamp
);