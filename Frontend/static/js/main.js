
const artistsContainer = document.getElementById("artists-container");

// Vérifier que la zone des artistes existe //
if (artistsContainer) {
    console.log("La zone des artistes a été trouvée !");
}

// Créer une carte pour un artiste //
function createArtistCard(artist) {
    const card = document.createElement("article");
    card.classList.add("artist-card");

    const image = document.createElement("img");
    image.src = artist.image;
    image.alt = artist.name;

    const info = document.createElement("div");
    info.classList.add("artist-info");

    const name = document.createElement("h3");
    name.textContent = artist.name;

    info.appendChild(name);
    card.appendChild(image);
    card.appendChild(info);

    return card;
}


function displayArtists(artists) {
    artistsContainer.innerHTML = "";

    artists.forEach(function (artist) {
        const card = createArtistCard(artist);
        artistsContainer.appendChild(card);
    });
}


// Récupérer les éléments de recherche
const searchInput = document.getElementById("searchInput");
const resultCount = document.getElementById("resultCount");

// Récupérer toutes les cartes déjà affichées par Go
const artistCards = document.querySelectorAll(".artist-card");

// Vérifier que les éléments existent
if (searchInput && resultCount) {

    // Afficher le nombre total d'artistes
    resultCount.textContent = artistCards.length + " artiste(s)";

    // Écouter les saisies dans la barre de recherche
    searchInput.addEventListener("input", function () {

        const searchValue = searchInput.value.toLowerCase().trim();

        let visibleArtists = 0;

        artistCards.forEach(function (card) {

            const artistData = card.dataset.search.toLowerCase();

            if (artistData.includes(searchValue)) {
                card.style.display = "";
                visibleArtists++;
            } else {
                card.style.display = "none";
            }

        });

        resultCount.textContent =
            visibleArtists + " artiste(s) trouvé(s)";
    });
}