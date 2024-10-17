CREATE TABLE employes (
    idEmployes CHAR(36) PRIMARY KEY, 
    name VARCHAR(30) NOT NULL,
    firstname VARCHAR(30) NOT NULL,
    birthdate DATE NOT NULL,
    mail VARCHAR(50) NOT NULL,
    city VARCHAR(30) NOT NULL,
    idDepartement CHAR(36) NOT NULL,
    idPost CHAR(36) NOT NULL,
    salary INTEGER NOT NULL,
    FOREIGN KEY (idDepartement) REFERENCES departement(idDepartement)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (idPost) REFERENCES post(idPost)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)

CREATE TABLE departement (
    idDepartement CHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
)

CREATE TABLE post (
    idPost CHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
)

CREATE TABLE project (
    idProject CHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    responsable CHAR(36) NOT NULL,
    FOREIGN KEY (responsable) REFERENCES employes(idEmployes)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)

CREATE TABLE employes_project (
    idEmployes CHAR(36) NOT NULL,
    idProject CHAR(36) NOT NULL,
    PRIMARY KEY (idEmployes, idProject),
    FOREIGN KEY (idEmployes) REFERENCES employes(idEmployes)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (idProject) REFERENCES project(idProject)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)

CREATE TABLE hierarchy (
    idEmployes CHAR(36) NOT NULL,
    idSuperior CHAR(36) NOT NULL,
    PRIMARY KEY (idEmployes, idSuperior),
    FOREIGN KEY (idEmployes) REFERENCES employes(idEmployes)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (idSuperior) REFERENCES employes(idEmployes)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)

