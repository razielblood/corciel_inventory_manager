create table if not exists Categories(
	ID	int auto_increment primary key,
	Name varchar(255) not null,
	Description varchar(1024)
	
);

create table if not exists Manufacturers(
	ID	int auto_increment primary key,
	Name varchar(255) not null
);


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