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

CREATE TABLE seats (
  id VARCHAR(32) PRIMARY KEY,
  flight_id VARCHAR(32) NOT NULL,
  class_seat VARCHAR(50) NOT NULL,
  price_seat FLOAT NOT NULL,
  number_seat INT NOT NULL,
  available BOOL NOT NULL DEFAULT FALSE,
  FOREIGN KEY (flight_id) REFERENCES flights(id)
);

INSERT INTO seats (id, flight_id,class_seat,price_seat,number_seat,available)
VALUES ('1', '1', 'BUSINESS', 30.0, 1, false);
INSERT INTO seats (id, flight_id,class_seat,price_seat,number_seat,available)
VALUES ('2', '1', 'ECONOMIC', 30.0, 10, false);
INSERT INTO seats (id, flight_id,class_seat,price_seat,number_seat,available)
VALUES ('3', '1', 'FREE', 30.0, 20, false);



CREATE TABLE baggages (
  id VARCHAR(32) PRIMARY KEY,
  type_baggage VARCHAR(50) NOT NULL,
  price FLOAT NOT NULL
);

INSERT INTO baggages VALUES ('1', 'ZERO', 0);
INSERT INTO baggages VALUES ('2', 'PLUS', 100);

CREATE TABLE bookings (
  id VARCHAR(32) PRIMARY KEY,
  flight_id VARCHAR(32) NOT NULL,
  number_passenger INT NOT NULL,
  total_price FLOAT NOT NULL,
  FOREIGN KEY (flight_id) REFERENCES flights(id)
);

CREATE TABLE booking_detail (
  id VARCHAR(32) PRIMARY KEY,
  booking_id VARCHAR(32) NOT NULL,
  user_id VARCHAR(32) NOT NULL,
  name_passenger VARCHAR(255) NOT NULL,
  lastname_passenger VARCHAR(255) NOT NULL,
  doc_passenger VARCHAR(20) NOT NULL,
  seat_id VARCHAR(32) NOT NULL,
  baggage_id  VARCHAR(32) NOT NULL,
  FOREIGN KEY (booking_id) REFERENCES bookings(id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (seat_id) REFERENCES seats(id),
  FOREIGN KEY (baggage_id) REFERENCES baggages(id)
);

CREATE TABLE booking_code (
  id VARCHAR(32) PRIMARY KEY,
  booking_id VARCHAR(32) NOT NULL,
  code_booking VARCHAR(10) NOT NULL,
  tickets VARCHAR(100) NOT NULL,
  FOREIGN KEY (booking_id) REFERENCES bookings(id)
);
