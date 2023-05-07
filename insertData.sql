insert into articles (title, contents, username, nice, created_at) values ('Hello World', 'This is my first article.', 'John Doe', 0, now());

insert into articles (title, contents, username, nice) values ('Hello World 2', 'This is my second article.', 'John Doe', 0);

insert into comments (article_id, message, created_at) values (1, 'This is my first comment.', now());

insert into comments (article_id, message, created_at) values (1, 'This is my second comment.', now());

