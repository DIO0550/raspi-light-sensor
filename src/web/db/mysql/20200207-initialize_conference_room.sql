
-- +migrate Up

INSERT INTO conference_room (name, usage_situation) VALUES ('1階A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('2階A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('3階A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('4階A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('5階A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('6階A', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('7階A', 0);

INSERT INTO conference_room (name, usage_situation) VALUES ('1階B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('2階B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('3階B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('4階B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('5階B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('6階B', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('7階B', 0);

INSERT INTO conference_room (name, usage_situation) VALUES ('1階C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('2階C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('3階C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('4階C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('5階C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('6階C', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('7階C', 0);

INSERT INTO conference_room (name, usage_situation) VALUES ('1階D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('2階D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('3階D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('4階D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('5階D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('6階D', 0);
INSERT INTO conference_room (name, usage_situation) VALUES ('7階D', 0);

-- +migrate Down

DELETE FROM conference_room WHERE name='1階A';
DELETE FROM conference_room WHERE name='2階A';
DELETE FROM conference_room WHERE name='3階A';
DELETE FROM conference_room WHERE name='4階A';
DELETE FROM conference_room WHERE name='5階A';
DELETE FROM conference_room WHERE name='6階A';
DELETE FROM conference_room WHERE name='7階A';

DELETE FROM conference_room WHERE name='1階B';
DELETE FROM conference_room WHERE name='2階B';
DELETE FROM conference_room WHERE name='3階B';
DELETE FROM conference_room WHERE name='4階B';
DELETE FROM conference_room WHERE name='5階B';
DELETE FROM conference_room WHERE name='6階B';
DELETE FROM conference_room WHERE name='7階B';

DELETE FROM conference_room WHERE name='1階C';
DELETE FROM conference_room WHERE name='2階C';
DELETE FROM conference_room WHERE name='3階C';
DELETE FROM conference_room WHERE name='4階C';
DELETE FROM conference_room WHERE name='5階C';
DELETE FROM conference_room WHERE name='6階C';
DELETE FROM conference_room WHERE name='7階C';

DELETE FROM conference_room WHERE name='1階D';
DELETE FROM conference_room WHERE name='2階D';
DELETE FROM conference_room WHERE name='3階D';
DELETE FROM conference_room WHERE name='4階D';
DELETE FROM conference_room WHERE name='5階D';
DELETE FROM conference_room WHERE name='6階D';
DELETE FROM conference_room WHERE name='7階D';