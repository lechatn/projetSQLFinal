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
('Projet Alpha', 1),
('Projet Beta', 2),
('Projet Gamma', 5),
('Projet Delta', 6),
('Projet Epsilon', 9),
('Projet Zeta', 11),
('Projet Eta', 13),
('Projet Iota', 17),
('Projet Kappa', 19);

INSERT INTO employes_project (idEmployes, idProject) 
VALUES 
(1, 1), 
(2, 1), 
(3, 2), 
(4, 2), 
(5, 3), 
(6, 3), 
(7, 4), 
(8, 4), 
(9, 5), 
(10, 5), 
(11, 6), 
(12, 6), 
(13, 7), 
(14, 7), 
(15, 8), 
(16, 8), 
(17, 9), 
(18, 9), 
(19, 10), 
(20, 10);

INSERT INTO hierarchy (idEmployes, idSuperior) 
VALUES 
(2, 1), 
(3, 1), 
(4, 2), 
(5, 3), 
(6, 3), 
(7, 4), 
(8, 4), 
(9, 5), 
(10, 6), 
(11, 7), 
(12, 8), 
(13, 9), 
(14, 10), 
(15, 11), 
(16, 12), 
(17, 13), 
(18, 14), 
(19, 15), 
(20, 16);