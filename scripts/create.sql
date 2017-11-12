-- CREATE temperature table

CREATE TABLE reading (
        id SERIAL NOT NULL PRIMARY KEY,
        temperature float(4) NOT NULL,
        humidity float(4) NOT NULL,
        added_at TIMESTAMP WITH TIME ZONE NOT NULL,
)