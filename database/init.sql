DROP TABLE IF EXISTS maps;

CREATE TABLE maps (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    imageUrl TEXT NOT NULL,
    bbox TEXT NOT NULL,
    opacity REAL NOT NULL DEFAULT 1.0
);

INSERT INTO maps (title, imageUrl, bbox, opacity) VALUES
('Plan de Paris 1750', 'https://velum.app/maps/paris_1750.png', '2.331,48.856,2.364,48.870', 0.8),
('Carte de Rome 1748', 'https://velum.app/maps/rome_1748.png', '12.465,41.888,12.490,41.902', 0.75),
('Londres 1800', 'https://velum.app/maps/london_1800.png', '0.087,51.504,0.112,51.516', 0.9);