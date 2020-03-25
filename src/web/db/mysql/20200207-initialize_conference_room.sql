
-- +migrate Up

INSERT INTO conference_room (name, usage_situation) VALUES ('1A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('2A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('3A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('4A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('5A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('6A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('7A', 0);

INSERT INTO conference_room (name, usage_situation) VALUES ('1B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('2B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('3B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('4B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('5B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('6B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('7B', 0);

INSERT INTO conference_room (name, usage_situation) VALUES ('1C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('2C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('3C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('4C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('5C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('6C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('7C', 0);

INSERT INTO conference_room (name, usage_situation) VALUES ('1D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('2D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('3D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('4D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('5D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('6D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('7D', 0);

-- +migrate Down

DELETE FROM conference_room WHERE name='1A';
DELETE FROM conference_room WHERE name='2A';
DELETE FROM conference_room WHERE name='3A';
DELETE FROM conference_room WHERE name='4A';
DELETE FROM conference_room WHERE name='5A';
DELETE FROM conference_room WHERE name='6A';
DELETE FROM conference_room WHERE name='7A';

DELETE FROM conference_room WHERE name='1B';
DELETE FROM conference_room WHERE name='2B';
DELETE FROM conference_room WHERE name='3B';
DELETE FROM conference_room WHERE name='4B';
DELETE FROM conference_room WHERE name='5B';
DELETE FROM conference_room WHERE name='6B';
DELETE FROM conference_room WHERE name='7B';

DELETE FROM conference_room WHERE name='1C';
DELETE FROM conference_room WHERE name='2C';
DELETE FROM conference_room WHERE name='3C';
DELETE FROM conference_room WHERE name='4C';
DELETE FROM conference_room WHERE name='5C';
DELETE FROM conference_room WHERE name='6C';
DELETE FROM conference_room WHERE name='7C';

DELETE FROM conference_room WHERE name='1D';
DELETE FROM conference_room WHERE name='2D';
DELETE FROM conference_room WHERE name='3D';
DELETE FROM conference_room WHERE name='4D';
DELETE FROM conference_room WHERE name='5D';
DELETE FROM conference_room WHERE name='6D';
DELETE FROM conference_room WHERE name='7D';