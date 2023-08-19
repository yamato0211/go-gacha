DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users`
(
    `id` int not null primary key auto_increment,
    `name` VARCHAR(20) unique not null,
    `token` VARCHAR(255) not null,
    `created_at` timestamp not null default current_timestamp,
    `updated_at` timestamp not null default current_timestamp on update current_timestamp
);