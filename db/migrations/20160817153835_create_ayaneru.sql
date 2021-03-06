
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS ayanerus (
id int(11) NOT NULL AUTO_INCREMENT,
source varchar(255) NOT NULL DEFAULT "",
url text NOT NULL,
created_at datetime DEFAULT NULL,
updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (id),
UNIQUE INDEX index_on_url (url(255)))
ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE ayanerus;
