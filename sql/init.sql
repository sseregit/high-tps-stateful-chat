USE chatting;

CREATE TABLE IF NOT EXISTS room (
    `id` bigint unsigned primary key NOT NULL auto_increment,
    `name` varchar(255) not null  unique,
    `createAt` datetime default current_timestamp,
    `updateAt` datetime default current_timestamp on update current_timestamp
);

CREATE TABLE IF NOT EXISTS chat (
    `id` bigint unsigned primary key NOT NULL auto_increment,
    `room` varchar(255) not null,
    `name` varchar(255) not null,
    `message` varchar(255) not null,
    `when` datetime default current_timestamp
);

CREATE TABLE IF NOT EXISTS serverInfo (
    `ip` varchar(255) primary key NOT NULL,
    `available` bool not null
);