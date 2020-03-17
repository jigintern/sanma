CREATE TABLE crawl_targets
(
  target_id INT UNIQUE AUTO_INCREMENT,
  target_url VARCHAR(128) NOT NULL,
  user_id VARCHAR(128) NOT NULL,
  PRIMARY KEY(target_id),
  FOREIGN KEY(user_id) REFERENCES users (user_id)
)
ENGINE=InnoDB;
