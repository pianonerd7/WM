CREATE TABLE ANC (
Word varchar(255) NOT NULL,
Lemma varchar (255) NOT NULL,
POS varchar(10) NOT NULL,
Frequency int NOT NULL,
Id int AUTO_INCREMENT,
PRIMARY KEY (Id)
);

DROP TABLE ANC;
LOAD DATA LOCAL INFILE '~/gocode/src/syntacticsub/ANC-all-count.txt' INTO TABLE ANC;