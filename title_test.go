package imdb

import "testing"

func TestTitle(t *testing.T) {
	_, err := NewTitle(client, "wrong")
	if err != ErrInvalidID {
		t.Errorf("NewTitle(wrong) = %v; want ErrInvalidId", err)
	}

	for _, tt := range []struct {
		ID   string
		want Title
	}{
		{
			ID: "tt0073845",
			want: Title{
				ID:       "tt0073845",
				URL:      "https://www.imdb.com/title/tt0073845",
				Name:     "L'uomo che sfidò l'organizzazione",
				Type:     "Movie",
				Year:     1975,
				Rating:   "5.2",
				Duration: "1h27m",
				Directors: []Name{
					Name{ID: "nm0340894", URL: "https://www.imdb.com/name/nm0340894", FullName: "Sergio Grieco"},
				},
				Writers: []Name{
					Name{ID: "nm0340894", URL: "https://www.imdb.com/name/nm0340894", FullName: "Sergio Grieco"},
					Name{ID: "nm0739308", URL: "https://www.imdb.com/name/nm0739308", FullName: "Rafael Romero Marchent"},
				},
				Actors: []Name{
					Name{ID: "nm0001886", URL: "https://www.imdb.com/name/nm0001886", FullName: "Howard Ross"},
					Name{ID: "nm0775810", URL: "https://www.imdb.com/name/nm0775810", FullName: "Karin Schubert"},
					Name{ID: "nm0237835", URL: "https://www.imdb.com/name/nm0237835", FullName: "Jean-Claude Dreyfus"},
					Name{ID: "nm0674183", URL: "https://www.imdb.com/name/nm0674183", FullName: "Nadine Perles"},
				},
				Genres:        []string{"Crime", "Drama"},
				Languages:     []string{"Italian"},
				Nationalities: []string{"Italy", "France", "Spain"},
				Description:   "L'uomo che sfidò l'organizzazione is a movie starring Howard Ross, Karin Schubert, and Jean-Claude Dreyfus. An airport employee switches a pack of drugs for baking soda and absconds to Barcelona, meanwhile the drug-runners are on...",
				Poster:        Media{ID: "rm143003904", TitleID: "tt0073845", URL: "https://www.imdb.com/title/tt0073845/mediaviewer/rm143003904", ContentURL: "https://m.media-amazon.com/images/M/MV5BOWVmYzRiZTAtOTRkMS00MDQxLTljNTItM2E2NjhmZGFmZDNkXkEyXkFqcGdeQXVyNTI0NjExNjA@._V1_UX182_CR0,0,182,268_AL_.jpg"},
				AKA:           []string{"Antimetopos me tin mafia", "El hombre que desafió a la organización", "L'homme qui défia l'organisation", "L'uomo che sfidò l'organizzazione", "One Man Against the Organization"},
			},
		},
		{
			ID: "tt0437804",
			want: Title{
				ID:       "tt0437803", // ID redirect.
				URL:      "https://www.imdb.com/title/tt0437803",
				Name:     "Alien Siege",
				Type:     "Movie",
				Year:     2005,
				Rating:   "3.5",
				Duration: "1h30m",
				Directors: []Name{
					Name{ID: "nm0821100", URL: "https://www.imdb.com/name/nm0821100", FullName: "Robert Stadd"},
				},
				Writers: []Name{
					Name{ID: "nm1305705", URL: "https://www.imdb.com/name/nm1305705", FullName: "Bill Lundy"},
					Name{ID: "nm0757507", URL: "https://www.imdb.com/name/nm0757507", FullName: "Paul Salamoff"},
					Name{ID: "nm0821100", URL: "https://www.imdb.com/name/nm0821100", FullName: "Robert Stadd"},
					Name{ID: "nm0884235", URL: "https://www.imdb.com/name/nm0884235", FullName: "Ian Valentine"},
				},
				Actors: []Name{
					Name{ID: "nm0424635", URL: "https://www.imdb.com/name/nm0424635", FullName: "Brad Johnson"},
					Name{ID: "nm1658541", URL: "https://www.imdb.com/name/nm1658541", FullName: "Erin Ross"},
					Name{ID: "nm0250169", URL: "https://www.imdb.com/name/nm0250169", FullName: "Lilas Lane"},
					Name{ID: "nm0003405", URL: "https://www.imdb.com/name/nm0003405", FullName: "Nathan Anderson"},
				},
				Genres:        []string{"Action", "Adventure", "Drama", "Sci-Fi", "Thriller"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "Alien Siege is a TV movie starring Brad Johnson, Erin Ross, and Lilas Lane. Earth is attacked by the Kulkus, a hostile breed infected by a lethal virus and needing human blood to develop an antidote. Earth's governments negotiate...",
				Poster:        Media{ID: "rm3287190272", TitleID: "tt0437803", URL: "https://www.imdb.com/title/tt0437803/mediaviewer/rm3287190272", ContentURL: "https://m.media-amazon.com/images/M/MV5BMTk1MTA4NDMwMF5BMl5BanBnXkFtZTcwNTM2MTIzMg@@._V1_UY268_CR4,0,182,268_AL_.jpg"},
				AKA:           []string{"A Föld ostroma", "Alien Blood", "Alien Siege", "Alien Siege - Tod aus dem All", "Etat de siège", "O Perigo Alienígena", "Obca krew"},
			},
		},
		{
			ID: "tt1179034",
			want: Title{
				ID:       "tt1179034",
				URL:      "https://www.imdb.com/title/tt1179034",
				Name:     "From Paris with Love",
				Type:     "Movie",
				Year:     2010,
				Rating:   "6.5",
				Duration: "1h32m",
				Directors: []Name{
					Name{ID: "nm0603628", URL: "https://www.imdb.com/name/nm0603628", FullName: "Pierre Morel"},
				},
				Writers: []Name{
					Name{ID: "nm0367867", URL: "https://www.imdb.com/name/nm0367867", FullName: "Adi Hasak"},
					Name{ID: "nm0000108", URL: "https://www.imdb.com/name/nm0000108", FullName: "Luc Besson"},
				},
				Actors: []Name{
					Name{ID: "nm0000237", URL: "https://www.imdb.com/name/nm0000237", FullName: "John Travolta"},
					Name{ID: "nm0001667", URL: "https://www.imdb.com/name/nm0001667", FullName: "Jonathan Rhys Meyers"},
					Name{ID: "nm0810738", URL: "https://www.imdb.com/name/nm0810738", FullName: "Kasia Smutniak"},
					Name{ID: "nm0243983", URL: "https://www.imdb.com/name/nm0243983", FullName: "Richard Durden"},
				},
				Genres:        []string{"Action", "Crime", "Thriller"},
				Languages:     []string{"English", "French", "Mandarin", "German"},
				Nationalities: []string{"France"},
				Description:   "From Paris with Love is a movie starring John Travolta, Jonathan Rhys Meyers, and Kasia Smutniak. In Paris, a young employee in the office of the US Ambassador hooks up with an American spy looking to stop a terrorist attack in the...",
				Poster:        Media{ID: "rm30481408", TitleID: "tt1179034", URL: "https://www.imdb.com/title/tt1179034/mediaviewer/rm30481408", ContentURL: "https://m.media-amazon.com/images/M/MV5BODAwMDFjNjktMWY2Mi00MmVhLWI0MjYtNzg4OTI0NzA5YzBjXkEyXkFqcGdeQXVyNTIzOTk5ODM@._V1_UX182_CR0,0,182,268_AL_.jpg"},
				AKA:           []string{"Apo to Parisi me agapi", "Armastusega Pariisist", "Bez soucitu", "Bons baisers de Paris", "De Paris com Amor", "Desde París con amor", "Din Paris, cu dragoste", "Dupla Implacável", "From Paris with Love", "Iz Pariza s ljubavlju", "Iz Pariza z ljubeznijo", "Linkejimai is Paryziaus", "MiParis be'ahava", "Paris'ten sevgilerle", "París en la Mira", "París en la mira", "Pozdrowienia z Paryza", "Párizsból szeretettel", "Sangre y amor en París", "Από το Παρίσι με αγάπη", "З Парижу з любов'ю", "Из Парижа с любовью", "От Париж с любов"},
			},
		},
		{
			ID: "tt0133093",
			want: Title{
				ID:       "tt0133093",
				URL:      "https://www.imdb.com/title/tt0133093",
				Name:     "The Matrix",
				Type:     "Movie",
				Year:     1999,
				Rating:   "8.7",
				Duration: "2h16m",
				Directors: []Name{
					Name{ID: "nm0905154", URL: "https://www.imdb.com/name/nm0905154", FullName: "Lana Wachowski"},
					Name{ID: "nm0905152", URL: "https://www.imdb.com/name/nm0905152", FullName: "Lilly Wachowski"},
				},
				Writers: []Name{
					Name{ID: "nm0905152", URL: "https://www.imdb.com/name/nm0905152", FullName: "Lilly Wachowski"},
					Name{ID: "nm0905154", URL: "https://www.imdb.com/name/nm0905154", FullName: "Lana Wachowski"},
				},
				Actors: []Name{
					Name{ID: "nm0000206", URL: "https://www.imdb.com/name/nm0000206", FullName: "Keanu Reeves"},
					Name{ID: "nm0000401", URL: "https://www.imdb.com/name/nm0000401", FullName: "Laurence Fishburne"},
					Name{ID: "nm0005251", URL: "https://www.imdb.com/name/nm0005251", FullName: "Carrie-Anne Moss"},
					Name{ID: "nm0915989", URL: "https://www.imdb.com/name/nm0915989", FullName: "Hugo Weaving"},
				},
				Genres:        []string{"Action", "Sci-Fi"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "The Matrix is a movie starring Keanu Reeves, Laurence Fishburne, and Carrie-Anne Moss. A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.",
				Poster:        Media{ID: "rm525547776", TitleID: "tt0133093", URL: "https://www.imdb.com/title/tt0133093/mediaviewer/rm525547776", ContentURL: "https://m.media-amazon.com/images/M/MV5BNzQzOTk3OTAtNDQ0Zi00ZTVkLWI0MTEtMDllZjNkYzNjNTc4L2ltYWdlXkEyXkFqcGdeQXVyNjU0OTQ0OTY@._V1_UX182_CR0,0,182,268_AL_.jpg"},
				AKA:           []string{"Fylkið", "La matrice", "La matriz", "Maatriks", "Matorikkusu", "Matrica", "Matriks", "Matrikss", "Matrix", "Mátrix", "The Matrix", "Матрикс", "Матрица", "Матрицата", "Матриця"},
			},
		},
		{
			ID: "tt0291830",
			want: Title{
				ID:       "tt0291830",
				URL:      "https://www.imdb.com/title/tt0291830",
				Name:     "Corps à coeur",
				Type:     "Movie",
				Year:     1981,
				Duration: "26m",
				Directors: []Name{
					Name{ID: "nm1029036", URL: "https://www.imdb.com/name/nm1029036", FullName: "Jean-François Després"},
					Name{ID: "nm1003289", URL: "https://www.imdb.com/name/nm1003289", FullName: "Alain Godon"},
				},
				Genres:        []string{"Documentary", "Short"},
				Nationalities: []string{"Canada"},
			},
		},
		{
			ID: "tt1965639",
			want: Title{
				ID:            "tt1965639",
				URL:           "https://www.imdb.com/title/tt1965639",
				Name:          "El clima y las cuatro estaciones",
				Type:          "TVSeries",
				Year:          1994,
				Genres:        []string{"Documentary"},
				Languages:     []string{"Spanish"},
				Nationalities: []string{"Spain"},
			},
		},
		{
			ID: "tt0086677",
			want: Title{
				ID:       "tt0086677",
				URL:      "https://www.imdb.com/title/tt0086677",
				Name:     "Brothers",
				Type:     "TVSeries",
				Year:     1984,
				Rating:   "8.0",
				Duration: "22m",
				Writers: []Name{
					Name{ID: "nm0515953", URL: "https://www.imdb.com/name/nm0515953", FullName: "David Lloyd"},
					Name{ID: "nm0031246", URL: "https://www.imdb.com/name/nm0031246", FullName: "Greg Antonacci"},
				},
				Actors: []Name{
					Name{ID: "nm0907107", URL: "https://www.imdb.com/name/nm0907107", FullName: "Robert Walden"},
					Name{ID: "nm0716618", URL: "https://www.imdb.com/name/nm0716618", FullName: "Paul Regina"},
					Name{ID: "nm0535933", URL: "https://www.imdb.com/name/nm0535933", FullName: "Brandon Maggart"},
					Name{ID: "nm0604699", URL: "https://www.imdb.com/name/nm0604699", FullName: "Hallie Todd"},
				},
				Genres:        []string{"Comedy"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "Brothers is a TV series starring Robert Walden, Paul Regina, and Brandon Maggart. Two conservative men support their younger brother when he comes out as gay, and help him navigate being openly homosexual in 1980s Philadelphia.",
				Poster:        Media{ID: "rm1328617472", TitleID: "tt0086677", URL: "https://www.imdb.com/title/tt0086677/mediaviewer/rm1328617472", ContentURL: "https://m.media-amazon.com/images/M/MV5BNDM3N2IxY2UtMDc4MS00ZDM5LWJjODUtMjQ5ODg2YWIxMTE4XkEyXkFqcGdeQXVyMDYxMTUwNg@@._V1_UY268_CR87,0,182,268_AL_.jpg"},
				AKA:           []string{"Braća", "Brothers", "Unter Brüdern"},
			},
		},
		{
			ID: "tt1371159",
			want: Title{
				ID:     "tt1371159",
				URL:    "https://www.imdb.com/title/tt1371159",
				Name:   "Iron Man 2",
				Type:   "VideoGame",
				Year:   2010,
				Rating: "6.2",
				Directors: []Name{
					Name{ID: "nm4157448", URL: "https://www.imdb.com/name/nm4157448", FullName: "Michael McCormick"},
					Name{ID: "nm4157551", URL: "https://www.imdb.com/name/nm4157551", FullName: "Robert Taylor"},
				},
				Writers: []Name{
					Name{ID: "nm3881781", URL: "https://www.imdb.com/name/nm3881781", FullName: "Matt Fraction"},
					Name{ID: "nm1411347", URL: "https://www.imdb.com/name/nm1411347", FullName: "Don Heck"},
				},
				Actors: []Name{
					Name{ID: "nm0000332", URL: "https://www.imdb.com/name/nm0000332", FullName: "Don Cheadle"},
					Name{ID: "nm0519680", URL: "https://www.imdb.com/name/nm0519680", FullName: "Eric Loomis"},
					Name{ID: "nm0000168", URL: "https://www.imdb.com/name/nm0000168", FullName: "Samuel L. Jackson"},
					Name{ID: "nm0482851", URL: "https://www.imdb.com/name/nm0482851", FullName: "Phil LaMarr"},
				},
				Genres:        []string{"Action", "Adventure", "Crime", "Fantasy", "Sci-Fi", "Thriller"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "Iron Man 2 is a video game starring Don Cheadle, Eric Loomis, and Samuel L. Jackson. Video game based upon the film of the same name.",
				Poster:        Media{ID: "rm1884363776", TitleID: "tt1371159", URL: "https://www.imdb.com/title/tt1371159/mediaviewer/rm1884363776", ContentURL: "https://m.media-amazon.com/images/M/MV5BMWNkYWIyZjctMGFmZi00NGYyLWIzY2QtNGEwNDQ2MDgwOWZiXkEyXkFqcGdeQXVyNTk5Nzg0MDE@._V1_UY268_CR17,0,182,268_AL_.jpg"},
				AKA:           []string{"Iron Man 2"},
			},
		},
		{
			ID: "tt0423866",
			want: Title{
				ID:       "tt0423866",
				URL:      "https://www.imdb.com/title/tt0423866",
				Name:     "Bin-jip",
				Type:     "Movie",
				Year:     2004,
				Rating:   "8.0",
				Duration: "1h28m",
				Directors: []Name{
					Name{ID: "nm1104118", URL: "https://www.imdb.com/name/nm1104118", FullName: "Ki-duk Kim"},
				},
				Writers: []Name{
					Name{ID: "nm1104118", URL: "https://www.imdb.com/name/nm1104118", FullName: "Ki-duk Kim"},
				},
				Actors: []Name{
					{ID: "nm1165901", URL: "https://www.imdb.com/name/nm1165901", FullName: "Seung-Yun Lee"},
					{ID: "nm1030819", URL: "https://www.imdb.com/name/nm1030819", FullName: "Hee Jae"},
					{ID: "nm1891528", URL: "https://www.imdb.com/name/nm1891528", FullName: "Hyuk-ho Kwon"},
					{ID: "nm1873389", URL: "https://www.imdb.com/name/nm1873389", FullName: "Jeong-ho Choi"},
				},
				Genres:        []string{"Crime", "Drama", "Romance"},
				Languages:     []string{"Korean"},
				Nationalities: []string{"South Korea", "Japan"},
				Description:   "Bin-jip is a movie starring Seung-Yun Lee, Hee Jae, and Hyuk-ho Kwon. A transient young man breaks into empty homes to partake of the vacationing residents' lives for a few days.",
				Poster:        Media{ID: "rm880057600", TitleID: "tt0423866", URL: "https://www.imdb.com/title/tt0423866/mediaviewer/rm880057600", ContentURL: "https://m.media-amazon.com/images/M/MV5BMTM1ODIwNzM5OV5BMl5BanBnXkFtZTcwNjk5MDkyMQ@@._V1_UX182_CR0,0,182,268_AL_.jpg"},
				AKA:           []string{"3-Iron", "3-iron", "Bin jib", "Bin-Jip - Der Schattenmann", "Bin-Jip - Leere Häuser", "Bin-jip", "Bin-jip - tomme huse", "Bos ev", "Bosh Evler", "Casa Vazia", "El espíritu de la pasión", "Empty Houses", "Ferro 3", "Ferro 3 - La casa vuota", "Hierro 3", "Hierro-3", "Järn 3:an", "Järntrean", "Khanehaye Khali", "Lehargish B'Bayit", "Locataires", "Lopakodó lelkek", "Menaj in trei", "Olomonahoi mazi", "Provalnik", "Pusty dom", "Rautakolmonen", "Tomme hus", "Tusti namai", "Utsusemi", "Ολομόναχοι Μαζί", "Порожнiй будинок", "Пустой дом", "Стик Nо 3"},
			},
		},
	} {
		got, err := NewTitle(client, tt.ID)
		if err != nil {
			t.Errorf("NewTitle(%s) error: %v", tt.ID, err)
		} else {
			if err := diffStruct(*got, tt.want); err != nil {
				t.Errorf("NewTitle(%s): %v", tt.ID, err)
			}
		}
	}
}
