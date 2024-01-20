CREATE TABLE news
(
    id UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    title    VARCHAR(512) NOT NULL CHECK ( title <> '' ),
    description VARCHAR NOT NULL CHECK ( description <> '' ),
    photo uuid,
    published_by uuid,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);