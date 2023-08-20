DROP TABLE IF EXISTS `characters_users`;

CREATE TABLE IF NOT EXISTS `characters_users` (
    `user_id` int not null,
    `character_id` int not null,
    `count` int not null,
    foreign key (user_id) references users(id),
    foreign key (character_id) references characters(id),
    primary key (user_id, character_id)
);