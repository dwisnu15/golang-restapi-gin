
CREATE TABLE users
(
    "username"            varchar PRIMARY KEY,
    "hashed_password"     varchar        NOT NULL,
    "full_name"           varchar        NOT NULL,
    "email"               varchar UNIQUE NOT NULL,
    "password_changed_at" timestamptz    NOT NULL DEFAULT ('0001-01-01 00:00:00Z'),
    "created_at"          timestamptz    NOT NULL DEFAULT (now())
);

CREATE FUNCTION insert_user(
    username_param character varying,
    password_param character varying,
    fullname_param character varying,
    email_param character varying) returns boolean
    language plpgsql
as
$$
BEGIN
    INSERT INTO users(username, hashed_password, full_name, email)
    VALUES (username_param, password_param, fullname_param, email_param);
    return true;
end;
$$;

CREATE FUNCTION get_user(username_param character varying)
    RETURNS TABLE
            (
                username            character varying,
                hashed_password     character varying,
                full_name           character varying,
                email               character varying,
                password_changed_at timestamptz,
                created_at          timestamptz
            )
    language plpgsql
as
$$
BEGIN
    return query
        SELECT *
        FROM users
        WHERE username = username_param
        LIMIT 1;
end;
$$;