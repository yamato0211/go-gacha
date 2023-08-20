DROP TABLE IF EXISTS `characters`;

CREATE TABLE IF NOT EXISTS `characters` (
    `id` int not null primary key auto_increment,
    `name` VARCHAR(255) unique not null,
    `description` text,
    `hp` int not null,
    `cost` int not null,
    `power` int not null,
    `speed` int not null,
    `rarity` int not null,
    `created_at` timestamp not null default current_timestamp,
    `updated_at` timestamp not null default current_timestamp on update current_timestamp
);