CREATE TABLE posts (
  post_id serial PRIMARY KEY,
  user_id VARCHAR(40) NOT NULL,
  title VARCHAR(200) UNIQUE NOT NULL,
  content TEXT
);

CREATE TABLE tags (
  tag_id serial PRIMARY KEY,
  user_id VARCHAR(40) NOT NULL,
  name VARCHAR(50) NOT NULL,
  CONSTRAINT tag_is_unique_for_user UNIQUE (name, user_id)
);

CREATE TABLE posts_tags (
  post_id int REFERENCES posts (post_id),
  tag_id int REFERENCES tags (tag_id),
  CONSTRAINT post_tag_pkey PRIMARY KEY (post_id, tag_id)
);