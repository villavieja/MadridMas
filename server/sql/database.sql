
CREATE TABLE IF NOT EXISTS incidents (
  id INT(11) NOT NULL AUTO_INCREMENT,
  title VARCHAR(200) DEFAULT NULL,
  description VARCHAR(200) DEFAULT NULL,
  /* anonymous incidents are associated to user 0 */
  fk_user int NOT NULL,
  latitude DOUBLE,
  longitude DOUBLE,
  creation_date DATE DEFAULT NULL,
  status int,
  PRIMARY KEY (id)
) ENGINE=InnoDB;


CREATE TABLE IF NOT EXISTS users (
  id_user INT(11) NOT NULL AUTO_INCREMENT,
  email VARCHAR(200) DEFAULT NULL,
  password VARCHAR(200) DEFAULT NULL,
  PRIMARY KEY (id_user)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS photos (
  id_photo INT(11) NOT NULL AUTO_INCREMENT,
  path VARCHAR(200) DEFAULT NULL,
  PRIMARY KEY (id_photo)
) ENGINE=InnoDB;


CREATE TABLE IF NOT EXISTS incidents_photos (
  fk_incident INT(11) NOT NULL,
  fk_photo INT(11) NOT NULL,
  PRIMARY KEY (fk_incident,fk_photo)
)ENGINE=InnoDB;
