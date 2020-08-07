CREATE TABLE supers (
    id           serial,
    uuid         integer UNIQUE,
    "type"       varchar(8),
    "name"       varchar(100),
    full_name    varchar(250),
    intelligence integer,
    "power"      integer,
    occupation   varchar(100),
    "image"      varchar(100)
);
