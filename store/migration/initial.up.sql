
CREATE TABLE IF NOT EXISTS  ddos_db (
 uuid varchar NOT NULL,
 url varchar(200) DEFAULT NULL,
 number_requests int,
 duration int,
 PRIMARY KEY (uuid)
);
