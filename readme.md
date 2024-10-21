// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table employes {
  idEmployes uuid [primary key]
  name varchar(30)
  firstname varchar(30)
  birthdate date
  mail varchar(50)
  city varchar(50)
  idDepartement uuid 
  idPost uuid
  salary integer
}

Table departement {
  idDepartement uuid [primary key]
  type varchar(50)
  responsable uuid
}

Table postes {
  idPost uuid [primary key]
  name varchar(50)
}

Table projet {
  idProject uuid [primary key]
  projectName varchar(50)
  responsable uuid
}

Table employes_project {
  idEmployes uuid
  idProject uuid
}

Table hierarchy {
  idEmploye uuid
  idManager uuid
}


Ref : employes.idDepartement > departement.idDepartement
Ref : employes.idPost > postes.idPost
Ref : departement.responsable > employes.idEmployes
Ref : employes_project.idEmployes > employes.idEmployes
Ref : employes_project.idProject > projet.idProject
Ref : projet.responsable > employes.idEmployes
Ref : hierarchy.idEmploye > employes.idEmployes
Ref : hierarchy.idManager > employes.idEmployes



<div class="content">
        {{range .}}
            <div>
                <h2>Id Employé : {{.IdEmployes}}</h2>
                <p>Nom : {{.Name}}</p>
                <p>Prénom : {{.Firstname}}</p>
                <p>Date de naissance : {{.Birthdate}}</p>
                <p>Email : {{.Mail}}</p>
                <p>Ville : {{.City}}</p>
                <p>ID Département : {{.IdDepartement}}</p>
                <p>ID Poste : {{.IdPost}}</p>
                <p>Salaire : {{.Salary}}</p>
            </div>
        {{end}}
    </div>
    <h1>Voici les employés de l'entreprise :</h1>