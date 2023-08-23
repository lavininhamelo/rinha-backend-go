CREATE TABLE persons (
                         id CHAR(36) DEFAULT (UUID()),
                         nickname VARCHAR(255) NOT NULL,
                         name VARCHAR(255) NOT NULL,
                         birthday VARCHAR(20),
                         stack TEXT,
                         PRIMARY KEY (id)
);