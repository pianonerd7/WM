CREATE TABLE ANC (
Word varchar(255) NOT NULL,
Lemma varchar (255) NOT NULL,
POS varchar(10) NOT NULL,
Frequency int NOT NULL,
Id int AUTO_INCREMENT,
PRIMARY KEY (Id)
);

CREATE TABLE USERS (
email varchar(20) NOT NULL,
secret varbinary(30) NOT NULL,
PRIMARY KEY (email)
); 

DROP TABLE ANC;
DROP TABLE USERS;
LOAD DATA LOCAL INFILE '~/gocode/src/syntacticsub/ANC-all-count.txt' INTO TABLE ANC;