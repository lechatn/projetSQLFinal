document.querySelectorAll('.card').forEach(card => {
  card.addEventListener('click', function() {
      card.classList.toggle('flipped'); // Ajoute ou retire la classe 'flipped' au clic
  });
});

function Displayform() {
  var formulaire = document.getElementById("form");
  if (formulaire.style.display == "none") {
    formulaire.style.display = "block";
    document.getElementById("name").focus(); // Met le focus sur le premier champ du formulaire
  } else {
    formulaire.style.display = "none";
  }  
}