CREATE TABLE Estimation(
    id INT,
    driver_esti VARCHAR(255),
    FOREIGN KEY (id) REFERENCES User(id)
)DEFAULT CHARSET=utf8;