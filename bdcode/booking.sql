CREATE TABLE book
(
	book_id int,
	book_name text NOT NULL,
	
	CONSTRAINT PK_book_book_id PRIMARY KEY(book_id)

);

CREATE TABLE author
(
	author_id int,
	author_fio text NOT NULL,
	
	CONSTRAINT PK_author_author_id PRIMARY KEY(author_id)

);

CREATE TABLE publisher
(
	publisher_id int,
	publisher_name text NOT NULL,
	
	CONSTRAINT PK_publisher_publisher_id PRIMARY KEY(publisher_id)

);

CREATE TABLE book_author (
    book_id int NOT NULL,
    author_id int NOT NULL
);

CREATE TABLE book_publisher (
    book_id int NOT NULL,
    publisher_id int NOT NULL
);


ALTER TABLE ONLY book_author
    ADD CONSTRAINT pk_book_author PRIMARY KEY (book_id, author_id);
	
ALTER TABLE ONLY book_publisher
    ADD CONSTRAINT pk_book_publisher PRIMARY KEY (book_id, publisher_id);
	
ALTER TABLE ONLY book_author
    ADD CONSTRAINT fk_book_author_author FOREIGN KEY (author_id) REFERENCES author;
	
ALTER TABLE ONLY book_author
    ADD CONSTRAINT fk_book_author_book FOREIGN KEY (book_id) REFERENCES book;

ALTER TABLE ONLY book_publisher
    ADD CONSTRAINT fk_book_publisher_publisher FOREIGN KEY (publisher_id) REFERENCES publisher;
	
ALTER TABLE ONLY book_publisher
    ADD CONSTRAINT fk_book_publisher_book FOREIGN KEY (book_id) REFERENCES book;
	





