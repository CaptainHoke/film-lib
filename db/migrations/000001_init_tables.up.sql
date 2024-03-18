CREATE TABLE IF NOT EXISTS actors
(
    actor_id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    sex CHAR(1) NOT NULL,
    birthdate DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS films
(
    film_id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT CHECK ( char_length(description) <= 1000 ),
    release_date DATE NOT NULL,
    rating DECIMAL(3, 1) NOT NULL
);

CREATE TABLE actor_film
(
    actor_id INT,
    film_id INT,
    PRIMARY KEY (actor_id, film_id),
    FOREIGN KEY (actor_id) REFERENCES actors(actor_id) ON DELETE CASCADE,
    FOREIGN KEY (film_id) REFERENCES films(film_id) ON DELETE CASCADE
);
