DROP TABLE IF EXISTS `gacha_contents`;

CREATE TABLE IF NOT EXISTS `gacha_contents` (
    `character_id` INT NOT NULL,
    `probability` int NOT NULL,
    FOREIGN KEY (character_id) REFERENCES characters(id),
    PRIMARY KEY (character_id)
);