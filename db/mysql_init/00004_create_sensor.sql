CREATE TABLE Sensor(
    id INT,
    acceleration_x Double,
    acceleration_y Double,
    acceleration_z Double,
    FOREIGN KEY (id) REFERENCES User(id)
)DEFAULT CHARSET=utf8;