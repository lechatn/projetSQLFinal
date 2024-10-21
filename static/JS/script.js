document.querySelectorAll('.card').forEach(card => {
  card.addEventListener('click', function() {
      card.classList.toggle('flipped'); // Ajoute ou retire la classe 'flipped' au clic
  });
});

