INSERT INTO Author
(author_id, author_fio)
VALUES
(1, 'J.D. Salinger'),
(2, 'F. Scott. Fitzgerald'),
(3, 'Jane Austen'),
(4, 'Scott Hanselman'),
(5, 'Jason N. GayloUSA'),
(6, 'Pranav Rastogi'),
(7, 'Todd Miranda'),
(8, 'Christian Wenz');

INSERT INTO book
(book_id, book_name)
VALUES
(1, 'The Catcher in the Rye'),
(2, 'Nine Stories'),
(3, 'Franny and Zooey'),
(4, 'The Great Gatsby'),
(5, 'Tender id the Night'),
(6, 'Pride and Prejudice'),
(7, 'Professional ASP.NET 4.5 in C# and VB');

INSERT INTO publisher
(publisher_id, publisher_name)
VALUES
(1, 'Moscow Label'),
(2, 'new york life'),
(3, 'Sidney library');


INSERT INTO book_author
(book_id, author_id)
VALUES
(1, 1),
(2, 1),
(3, 1),
(4, 2),
(5, 2),
(6, 3),
(7, 4),
(7, 5),
(7, 6),
(7, 7),
(7, 8);



INSERT INTO book_publisher
(book_id, publisher_id)
VALUES
(1, 1),
(2, 1),
(3, 3),
(4, 3),
(5, 2),
(6, 2),
(7, 1),
(7, 2),
(7, 3);

