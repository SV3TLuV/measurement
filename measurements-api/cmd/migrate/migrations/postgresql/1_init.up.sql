CREATE TABLE object_types
(
    object_type_id SERIAL      NOT NULL PRIMARY KEY,
    title          VARCHAR(64) NOT NULL CONSTRAINT object_types_title_unique UNIQUE
);


CREATE TABLE objects
(
    object_id   INTEGER      NOT NULL PRIMARY KEY,
    parent_id   INTEGER      REFERENCES objects ON DELETE CASCADE,
    type_id     INTEGER      NOT NULL REFERENCES object_types ON DELETE CASCADE,
    title       VARCHAR(128) NOT NULL,
    address     VARCHAR(128),
    lat         DOUBLE PRECISION,
    lon         DOUBLE PRECISION
);


CREATE TABLE post_infos
(
    object_id              INTEGER   NOT NULL PRIMARY KEY REFERENCES objects ON DELETE CASCADE,
    last_polling_date_time TIMESTAMP,
    is_listened            BOOLEAN NOT NULL DEFAULT FALSE
);


CREATE TABLE measurements
(
    measurement_id         INTEGER       NOT NULL PRIMARY KEY ,
    object_id              INTEGER       NOT NULL REFERENCES objects ON DELETE CASCADE,
    obj_name               VARCHAR(256),
    obj_num                VARCHAR(32),
    created                TIMESTAMP     NOT NULL,
    changed                TIMESTAMP,
    date_time              VARCHAR(48),
    real_date_time         TIMESTAMP,
    temp                   REAL,
    pressure               REAL,
    wind_dir               INTEGER,
    wind_dir_str           VARCHAR(16),
    wind_speed             REAL,
    humid                  REAL,
    water_vapor_elasticity REAL,
    atm_phenom             REAL,
    humid_int              REAL,
    temp_int               REAL,

    v_202917 REAL,
    m_202917 REAL,
    q_202917 VARCHAR(128),

    v_202918 REAL,
    m_202918 REAL,
    q_202918 VARCHAR(128),

    v_202919 REAL,
    m_202919 REAL,
    q_202919 VARCHAR(128),

    v_202920 REAL,
    m_202920 REAL,
    q_202920 VARCHAR(128),

    v_202921 REAL,
    m_202921 REAL,
    q_202921 VARCHAR(128),

    v_202932 REAL,
    m_202932 REAL,
    q_202932 VARCHAR(128),

    v_202935 REAL,
    m_202935 REAL,
    q_202935 VARCHAR(128),

    v_202924 REAL,
    m_202924 REAL,
    q_202924 VARCHAR(128),

    v_202925 REAL,
    m_202925 REAL,
    q_202925 VARCHAR(128),

    v_203565 REAL,
    m_203565 REAL,
    q_203565 VARCHAR(128),

    v_209190 REAL,
    m_209190 REAL,
    q_209190 VARCHAR(128),

    v_203570 REAL,
    m_203570 REAL,
    q_203570 VARCHAR(128),

    v_203551 REAL,
    m_203551 REAL,
    q_203551 VARCHAR(128),

    v_202936 REAL,
    m_202936 REAL,
    q_202936 VARCHAR(128),

    v_203569 REAL,
    m_203569 REAL,
    q_203569 VARCHAR(128),

    v_203557 REAL,
    m_203557 REAL,
    q_203557 VARCHAR(128),

    v_203568 REAL,
    m_203568 REAL,
    q_203568 VARCHAR(128),

    v_203559 REAL,
    m_203559 REAL,
    q_203559 VARCHAR(128),

    v_203577 REAL,
    m_203577 REAL,
    q_203577 VARCHAR(128),

    v_211082 REAL,
    m_211082 REAL,
    q_211082 VARCHAR(128),

    v_202931 REAL,
    m_202931 REAL,
    q_202931 VARCHAR(128)
);


CREATE TABLE columns
(
    column_id   SERIAL       NOT NULL PRIMARY KEY,
    title       VARCHAR(128) NOT NULL,
    short_title VARCHAR(128) NOT NULL,
    formula     VARCHAR(128),
    obj_field   VARCHAR(48)  NOT NULL CONSTRAINT columns_name_unique UNIQUE,
    code        VARCHAR(8)
);


CREATE TABLE qualities
(
    quality_id SERIAL       NOT NULL PRIMARY KEY,
    priority   INTEGER      NOT NULL,
    caption    VARCHAR(256) NOT NULL,
    color      VARCHAR(16)  NOT NULL,
    title      VARCHAR(32)  NOT NULL
);


CREATE TABLE roles
(
    role_id SERIAL      NOT NULL PRIMARY KEY,
    title   VARCHAR(48) NOT NULL CONSTRAINT roles_name_unique UNIQUE,
    name    VARCHAR(48) NOT NULL
);


CREATE TABLE permissions
(
    permission_id SERIAL       NOT NULL PRIMARY KEY,
    name          VARCHAR(128) NOT NULL,
    title         VARCHAR(128) NOT NULL
);


CREATE TABLE users
(
    user_id       SERIAL       NOT NULL PRIMARY KEY,
    login         VARCHAR(128) NOT NULL CONSTRAINT users_login_unique UNIQUE,
    password_hash VARCHAR(512) NOT NULL,
    role_id       INTEGER      NOT NULL REFERENCES roles ON DELETE CASCADE,
    is_blocked    BOOLEAN      NOT NULL DEFAULT FALSE
);


CREATE TABLE user_posts
(
    user_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    object_id INTEGER NOT NULL REFERENCES objects ON DELETE CASCADE,
    PRIMARY KEY (object_id, user_id)
);


CREATE TABLE user_permissions
(
    user_id       INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES permissions ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id)
);


CREATE TABLE user_columns
(
    user_id   INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    column_id INTEGER NOT NULL REFERENCES columns ON DELETE CASCADE,
    PRIMARY KEY (user_id, column_id)
);


CREATE TABLE sessions
(
    session_id    UUID         NOT NULL PRIMARY KEY,
    refresh_token VARCHAR(512) NOT NULL,
    user_id       INTEGER      NOT NULL REFERENCES users ON DELETE CASCADE,
    created       TIMESTAMP    NOT NULL DEFAULT timezone('utc'::text, now()),
    updated       TIMESTAMP
);