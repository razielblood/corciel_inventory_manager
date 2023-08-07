drop table if exists ProductPresentation;
drop table if exists Products;
drop table if exists Brands;
drop table if exists Manufacturers;
drop table if exists Categories;

create table if not exists Categories(
	ID	int auto_increment primary key,
	Name varchar(255) not null unique,
	Description varchar(1024)
	
);

create table if not exists Manufacturers(
	ID	int auto_increment primary key,
	Name varchar(255) not null unique
);

create table if not exists Brands(
	ID int auto_increment primary key,
	Name varchar(255) not null unique,
	Manufacturer int,
	foreign key (Manufacturer) references Manufacturers(ID)
);

create table if not exists Products(
	ID	int auto_increment primary key,
	Name varchar(255) not null,
	Description varchar(1024),
	Image varchar(1024),
	Brand int,
	Category int,
	foreign key (Brand) references Brands(ID),
	foreign key (Category) references Categories(ID)
);

create table if not exists ProductPresentation(
	Product int not null,
	ID int primary key,
	WeightInKG float,
	PiecesPerPackage int,
	foreign key (Product) references Products(ID)
);
