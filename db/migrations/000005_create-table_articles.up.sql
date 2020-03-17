CREATE TABLE articles
(
  article_id INT UNIQUE AUTO_INCREMENT,
  author VARCHAR(128),
  article_status INT NOT NULL,
  article_type INT NOT NULL,
  article_url VARCHAR(128),
  title VARCHAR(128) NOT NULL,
  contents TEXT,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(article_id),
  FOREIGN KEY(author) REFERENCES users (user_id),
  FOREIGN KEY(article_status) REFERENCES a_status (status_id),
  FOREIGN KEY(article_type) REFERENCES a_type (type_id)
)
ENGINE=InnoDB;
