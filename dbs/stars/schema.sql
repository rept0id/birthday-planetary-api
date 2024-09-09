CREATE TABLE stars (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    ly FLOAT, -- Light Years
    ly_pm FLOAT, -- Light Years Plus Minus
);

CREATE TABLE stars_metakeys (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE star_meta (
    id SERIAL PRIMARY KEY,
    public BOOLEAN DEFAULT FALSE,
    star_id INTEGER REFERENCES stars(id),
    metakey_id INTEGER REFERENCES stars_metakeys(id),
    value VARCHAR(255) NOT NULL
);
