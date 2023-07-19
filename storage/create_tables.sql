drop table if exists Categories;

create table if not exists Categories(
	ID	int auto_increment primary key,
	Name varchar(255) not null unique,
	Description varchar(1024)
	
);

drop table if exists Manufacturer;

create table if not exists Manufacturers(
	ID	int auto_increment primary key,
	Name varchar(255) not null unique
);

drop table if exists Products;

create table if not exists Products(
	ID	int auto_increment primary key,
	Name varchar(255) not null,
	Description varchar(1024),
	WeightInKG float,
	PiecesPerPackage int,
	Image varchar(1024),
	Manufacturer int,
	Category int,
	foreign key (Manufacturer) references Manufacturers(ID),
	foreign key (Category) references Categories(ID)
);

drop table if exists Users;

create table if not exists Users(
	Username varchar(255) primary key,
	Password varchar(255) not null,
	FirstName varchar(255),
	LastName varchar(255),
	Email varchar(255),
	CreatedAt timestamp default current_timestamp
);