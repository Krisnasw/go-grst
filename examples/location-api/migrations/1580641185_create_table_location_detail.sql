DROP TABLE IF EXISTS tbl_city;
DROP TABLE IF EXISTS tbl_province;

CREATE TABLE `tbl_province` (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id )
);

CREATE TABLE `tbl_city` (
    id int NOT NULL AUTO_INCREMENT,
    province_id int NOT NULL,
    name varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY ( id ),
    FOREIGN KEY (province_id) REFERENCES tbl_province(id)
);

INSERT INTO `tbl_province` (name) values ('DKI Jakarta');
INSERT INTO `tbl_city` (province_id, name) values (1, 'Kab. Kep. Seribu');
INSERT INTO `tbl_city` (province_id, name) values (1, 'Jakarta Pusat');
INSERT INTO `tbl_city` (province_id, name) values (1, 'Jakarta Utara');
INSERT INTO `tbl_city` (province_id, name) values (1, 'Jakarta Barat');
INSERT INTO `tbl_city` (province_id, name) values (1, 'Jakarta Selatan');
INSERT INTO `tbl_city` (province_id, name) values (1, 'Jakarta Timur');

INSERT INTO `tbl_province` (name) values ('Jawa Barat');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Bogor');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Sukabumi');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Cianjur');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Bandung');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Garut');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Tasikmalaya');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Ciamis');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Kuningan');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Cirebon');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Majalengka');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Sumedang');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Indramayu');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Subang');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Purwakarta');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Karawang');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Bekasi');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Bandung Barat');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kab. Pangandaran');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kota Bogor');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kota Sukabumi');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kota Bandung');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kota Cirebon');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kota Bekasi');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kota Depok');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kota Cimahi');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kota Tasikmalaya');
INSERT INTO `tbl_city` (province_id, name) values (2, 'Kota Banjar');