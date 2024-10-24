CREATE TABLE employes (
    idEmployes INTEGER PRIMARY KEY AUTOINCREMENT, 
    name VARCHAR(30) NOT NULL,
    firstname VARCHAR(30) NOT NULL,
    birthdate DATE NOT NULL,
    mail VARCHAR(50) NOT NULL,
    city VARCHAR(30) NOT NULL,
    idDepartement INTEGER NOT NULL,
    idPost INTEGER NOT NULL,
    salary INTEGER NOT NULL,
    FOREIGN KEY (idDepartement) REFERENCES departement(idDepartement)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (idPost) REFERENCES post(idPost)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE departement (
    idDepartement INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE post (
    idPost INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE project (
    idProject INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL,
    responsable INTEGER,
    FOREIGN KEY (responsable) REFERENCES employes(idEmployes)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE employes_project (
    idEmployes INTEGER NOT NULL,
    idProject INTEGER NOT NULL,
    PRIMARY KEY (idEmployes, idProject),
    FOREIGN KEY (idEmployes) REFERENCES employes(idEmployes)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (idProject) REFERENCES project(idProject)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE hierarchy (
    idEmployes INTEGER NOT NULL,
    idSuperior INTEGER NOT NULL,
    PRIMARY KEY (idEmployes, idSuperior),
    FOREIGN KEY (idEmployes) REFERENCES employes(idEmployes)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (idSuperior) REFERENCES employes(idEmployes)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

INSERT INTO departement (name) 
VALUES 
('Informatique'), 
('Ressources Humaines'), 
('Marketing'), 
('Finances'), 
('Ventes'), 
('Support Client'), 
('Logistique'), 
('Recherche & Développement');

INSERT INTO post (name) 
VALUES 
('Développeur'), 
('Manager'), 
('Analyste'), 
('Technicien'), 
('Comptable'), 
('Directeur'), 
('Assistant'), 
('Commercial'), 
('Chef de Projet'), 
('Consultant');

INSERT INTO employes (name, firstname, birthdate, mail, city, idDepartement, idPost, salary) 
VALUES 
('Dupont', 'Jean', '1985-05-12', 'jean.dupont@mail.com', 'Paris', 1, 1, 3000),
('Martin', 'Sophie', '1990-09-23', 'sophie.martin@mail.com', 'Lyon', 2, 2, 4000),
('Durand', 'Luc', '1975-02-14', 'luc.durand@mail.com', 'Marseille', 3, 3, 3500),
('Bernard', 'Claire', '1988-11-22', 'claire.bernard@mail.com', 'Nice', 1, 2, 4200),
('Lefevre', 'Pierre', '1995-08-19', 'pierre.lefevre@mail.com', 'Bordeaux', 2, 4, 2500),
('Moreau', 'Juliette', '1993-07-02', 'juliette.moreau@mail.com', 'Lille', 3, 3, 3300),
('Roux', 'Michel', '1979-10-13', 'michel.roux@mail.com', 'Toulouse', 4, 5, 5000),
('Gauthier', 'Camille', '1981-01-26', 'camille.gauthier@mail.com', 'Paris', 5, 6, 6000),
('Girard', 'Marc', '1992-03-30', 'marc.girard@mail.com', 'Lyon', 6, 7, 2700),
('Lopez', 'Nathalie', '1983-12-17', 'nathalie.lopez@mail.com', 'Marseille', 7, 8, 2900),
('Muller', 'Louis', '1998-09-06', 'louis.muller@mail.com', 'Nice', 8, 9, 4500),
('Fournier', 'Alice', '1976-05-24', 'alice.fournier@mail.com', 'Bordeaux', 1, 1, 3100),
('Mercier', 'Paul', '1989-02-21', 'paul.mercier@mail.com', 'Toulouse', 2, 2, 3800),
('Blanc', 'Monique', '1994-06-15', 'monique.blanc@mail.com', 'Paris', 3, 3, 3400),
('Dupuis', 'Hugo', '1987-09-01', 'hugo.dupuis@mail.com', 'Lille', 4, 4, 2700),
('Faure', 'Eva', '1991-10-29', 'eva.faure@mail.com', 'Lyon', 5, 5, 4800),
('Andre', 'Mathieu', '1986-03-11', 'mathieu.andre@mail.com', 'Nice', 6, 6, 5900),
('Bonnet', 'Charlotte', '1978-07-07', 'charlotte.bonnet@mail.com', 'Bordeaux', 7, 7, 3100),
('Dupre', 'Valentin', '1996-04-12', 'valentin.dupre@mail.com', 'Toulouse', 8, 8, 4300),
('Schmitt', 'Laura', '1997-11-20', 'laura.schmitt@mail.com', 'Paris', 1, 9, 4700);

INSERT INTO project (name, responsable) 
VALUES 
('Projet Hangman', 1),
('Projet Groupie', 2),
('Projet Java', 5),
('Projet Python', 6),
('Projet SQL', 9),
('Projet Forum', 11),
('Projet JS', 13);

INSERT INTO employes_project (idEmployes, idProject) 
VALUES 
(1, 1),
(1, 2),
(2, 3),
(2, 4),
(3, 5),
(3, 6),
(4, 7),
(4, 1),
(5, 2),
(5, 3),
(6, 4),
(6, 5),
(7, 6),
(7, 7),
(8, 1),
(8, 2),
(9, 3),
(9, 4),
(10, 5),
(10, 6),
(11, 7),
(11, 1),
(12, 2),
(12, 3),
(13, 4),
(13, 5),
(14, 6),
(14, 7),
(15, 1),
(15, 2),
(16, 3),
(16, 4),
(17, 5),
(17, 6),
(18, 7),
(18, 1),
(19, 2),
(19, 3),
(20, 4),
(20, 5);


INSERT INTO hierarchy (idEmployes, idSuperior) 
VALUES
(1, 8),  -- Jean Dupont (3000) -> Camille Gauthier (6000)
(2, 17), -- Sophie Martin (4000) -> Mathieu Andre (5900)
(3, 17), -- Luc Durand (3500) -> Mathieu Andre (5900)
(4, 17), -- Claire Bernard (4200) -> Mathieu Andre (5900)
(5, 8),  -- Pierre Lefevre (2500) -> Camille Gauthier (6000)
(6, 8),  -- Juliette Moreau (3300) -> Camille Gauthier (6000)
(7, 8),  -- Michel Roux (5000) -> Camille Gauthier (6000)
(9, 8),  -- Marc Girard (2700) -> Camille Gauthier (6000)
(10, 8), -- Nathalie Lopez (2900) -> Camille Gauthier (6000)
(11, 17), -- Louis Muller (4500) -> Mathieu Andre (5900)
(12, 8), -- Alice Fournier (3100) -> Camille Gauthier (6000)
(13, 17), -- Paul Mercier (3800) -> Mathieu Andre (5900)
(14, 8), -- Monique Blanc (3400) -> Camille Gauthier (6000)
(15, 8), -- Hugo Dupuis (2700) -> Camille Gauthier (6000)
(16, 17), -- Eva Faure (4800) -> Mathieu Andre (5900)
(18, 8), -- Charlotte Bonnet (3100) -> Camille Gauthier (6000)
(19, 17), -- Valentin Dupre (4300) -> Mathieu Andre (5900)
(20, 17); -- Laura Schmitt (4700) -> Mathieu Andre (5900)
