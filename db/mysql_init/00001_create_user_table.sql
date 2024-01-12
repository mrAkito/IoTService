CREATE TABLE User(
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255),
    collab_service VARCHAR(255),
    login Boolean,
    birthday DATE,
    payment_method VARCHAR(255),
    phone_number VARCHAR(255),
    kind Boolean
)DEFAULT CHARSET=utf8;