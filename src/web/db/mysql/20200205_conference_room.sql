-- +migrate Up
CREATE TABLE conference_room (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  name CHAR(30) NOT NULL,
  usage_situation boolean NOT NULL,
  PRIMARY KEY (id)
);


-- +migrate Down
DROP TABLE conference_room;