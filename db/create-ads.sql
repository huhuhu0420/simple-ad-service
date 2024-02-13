CREATE TABLE ads (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL CHECK (end_at > start_at)
);

CREATE TABLE ad_ages (
    ad_id INTEGER REFERENCES ads(id),
    age_start INTEGER CHECK (age_start >= 1 AND age_start <= 99),
    age_end INTEGER CHECK (age_end >= 1 AND age_end <= 99 AND age_end >= age_start),
    PRIMARY KEY (ad_id)
);

CREATE TABLE countries (
    code CHAR(2) PRIMARY KEY
);
CREATE TABLE ad_countries (
    ad_id INTEGER REFERENCES ads(id),
    country_code CHAR(2) REFERENCES countries(code),
    PRIMARY KEY (ad_id, country_code)
);

CREATE TABLE ad_platforms (
    ad_id INTEGER REFERENCES ads(id),
    platform VARCHAR(10) CHECK (platform IN ('Android', 'iOS', 'Web')),
    PRIMARY KEY (ad_id, platform)
);

CREATE TABLE ad_genders (
    ad_id INTEGER REFERENCES ads(id),
    gender CHAR(1) CHECK (gender IN ('M', 'F')),
    PRIMARY KEY (ad_id, gender)
);
