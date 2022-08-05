BEGIN;

CREATE TABLE IF NOT EXISTS genders
(
    id          BIGSERIAL PRIMARY KEY,
    gender_name VARCHAR NOT NULL
);

INSERT INTO genders (gender_name)
VALUES ('Male'),
       ('Female'),
       ('Not stated');

CREATE TABLE IF NOT EXISTS statuses
(
    id     BIGSERIAL PRIMARY KEY,
    status VARCHAR NOT NULL
);

INSERT INTO statuses (status)
VALUES ('Pending'),
       ('Accepted'),
       ('Rejected'),
       ('Blocked');


CREATE TABLE IF NOT EXISTS users
(
    id              uuid PRIMARY KEY,
    first_name      VARCHAR NOT NULL,
    middle_name     VARCHAR,
    last_name       VARCHAR,
    gender_id       INT              DEFAULT 3,
    hashed_password VARCHAR NOT NULL,
    email           VARCHAR NOT NULL,
    phone           VARCHAR,
    bio             TEXT,
    photo_url       VARCHAR,
    active_from     DATE    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    active_to       DATE,
    FOREIGN KEY (gender_id) REFERENCES genders (id)
);


CREATE TABLE IF NOT EXISTS user_friends
(
    id         BIGSERIAL PRIMARY KEY,
    source_id  uuid NOT NULL,
    target_id  uuid NOT NULL,
    status     INT  NOT NULL DEFAULT 0,
    created_at DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    message    TEXT,
    FOREIGN KEY (source_id) REFERENCES users (id),
    FOREIGN KEY (target_id) REFERENCES users (id),
    FOREIGN KEY (status) REFERENCES statuses (id)
);

CREATE TABLE IF NOT EXISTS messages
(
    id              BIGSERIAL PRIMARY KEY,
    source_id       uuid NOT NULL,
    target_id       uuid NOT NULL,
    message         TEXT,
    attachments_url VARCHAR[],
    created_at      DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATE,
    FOREIGN KEY (source_id) REFERENCES users (id),
    FOREIGN KEY (target_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS user_posts
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    uuid NOT NULL,
    message    text,
    created_at DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

COMMIT;