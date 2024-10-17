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
INSERT INTO employes (name, firstname, birthdate, mail, city, idDepartement, idPost, salary) VALUES ('Doe', 'John', '1980-01-01', 'john.doe@gmail.com', 'Paris', 1, 1, 50000);