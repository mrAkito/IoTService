CREATE TABLE Car(
    id INT,
    number VARCHAR(255),
    equipment VARCHAR(255),
    credit VARCHAR(255),
    FOREIGN KEY (id) REFERENCES User(id)
)DEFAULT CHARSET=utf8;