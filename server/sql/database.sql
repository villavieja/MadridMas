
CREATE TABLE IF NOT EXISTS incidents (
  id INT(11) NOT NULL AUTO_INCREMENT,
  latitude DOUBLE,
  longitude DOUBLE,
  creation_date DATE DEFAULT NULL,
  description VARCHAR(200) DEFAULT NULL,
  status int,
  PRIMARY KEY (id)
) ENGINE=InnoDB;
