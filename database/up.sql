DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS flights;

-- create
CREATE TABLE flights (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(10) NOT NULL,
  airplane_id VARCHAR(32) NOT NULL,
  departure_date DATE NOT NULL,
  departure_time VARCHAR(10) NOT NULL,
  arrival_date DATE NOT NULL,
  arrival_time VARCHAR(10) NOT NULL,
  origin_city VARCHAR(255) NOT NULL,
  destination_city VARCHAR(255) NOT NULL,
  price FLOAT NOT NULL
);

INSERT INTO flights (id, name, airplane_id, departure_date, departure_time, arrival_date, arrival_time, origin_city, destination_city,price) VALUES ('1', 'AV2060', '123132', to_date('16/04/2024', 'DD/MM/YYYY'), '13:30', to_date('20/04/2024', 'DD/MM/YYYY'), '14:30', 'LIMA', 'CUZCO', 315.0);
 
INSERT INTO flights (id, name, airplane_id, departure_date, departure_time, arrival_date, arrival_time, origin_city, destination_city,price) VALUES ('2', 'AV2060', '123132', to_date('17/04/2024', 'DD/MM/YYYY'), '17:30', to_date('22/04/2024', 'DD/MM/YYYY'), '14:30', 'LIMA', 'CUZCO', 315.0);

INSERT INTO flights (id, name, airplane_id, departure_date, departure_time, arrival_date, arrival_time, origin_city, destination_city,price) VALUES ('3', 'AV2060', '123132', to_date('20/04/2024', 'DD/MM/YYYY'), '13:30', to_date('25/04/2024', 'DD/MM/YYYY'), '14:30', 'LIMA', 'CUZCO', 315.0);