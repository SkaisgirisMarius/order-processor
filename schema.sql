CREATE TABLE orders (
    id INT(11) NOT NULL AUTO_INCREMENT,
    proxy_count INT(11) NOT NULL CHECK (proxy_count <= 100),
    name VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB CHARSET=utf8mb4;
