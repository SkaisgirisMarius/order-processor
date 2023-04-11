CREATE TABLE orders (
                        id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        proxy_count bigint(20) NOT NULL CHECK (proxy_count <= 100),
                        name varchar(30) DEFAULT NULL,
                        created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
