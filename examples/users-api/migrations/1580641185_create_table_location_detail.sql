DROP TABLE IF EXISTS tbl_users;

CREATE TABLE `tbl_users` (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    province_id int NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id )
);


INSERT INTO `tbl_users` (name, province_id) values ('Lorem', 1);
INSERT INTO `tbl_users` (name, province_id) values ('Ipsum', 2);
