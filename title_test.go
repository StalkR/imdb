package imdb

import (
	"fmt"
	"testing"
)

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
				URL:      "http://www.imdb.com/title/tt0073845",
				Name:     "L'uomo che sfidò l'organizzazione",
				Year:     1975,
				Rating:   "5.3",
				Duration: "87m",
				Directors: []Name{
					Name{ID: "nm0340894", URL: "http://www.imdb.com/name/nm0340894", FullName: "Sergio Grieco"},
				},
				Writers: []Name{
					Name{ID: "nm0340894", URL: "http://www.imdb.com/name/nm0340894", FullName: "Sergio Grieco"},
				},
				Actors: []Name{
					Name{ID: "nm0001886", URL: "http://www.imdb.com/name/nm0001886", FullName: "Howard Ross"},
					Name{ID: "nm0775810", URL: "http://www.imdb.com/name/nm0775810", FullName: "Karin Schubert"},
					Name{ID: "nm0237835", URL: "http://www.imdb.com/name/nm0237835", FullName: "Jean-Claude Dreyfus"},
					Name{ID: "nm0674183", URL: "http://www.imdb.com/name/nm0674183", FullName: "Nadine Perles"},
					Name{ID: "nm0197623", URL: "http://www.imdb.com/name/nm0197623", FullName: "Alberto Dalbés"},
					Name{ID: "nm0130952", URL: "http://www.imdb.com/name/nm0130952", FullName: "José Calvo"},
					Name{ID: "nm0000963", URL: "http://www.imdb.com/name/nm0000963", FullName: "Stephen Boyd"},
					Name{ID: "nm0157826", URL: "http://www.imdb.com/name/nm0157826", FullName: "José Luis Chinchilla"},
					Name{ID: "nm0054079", URL: "http://www.imdb.com/name/nm0054079", FullName: "Pietro Torrisi"},
					Name{ID: "nm0877114", URL: "http://www.imdb.com/name/nm0877114", FullName: "Luciana Turina"},
					Name{ID: "nm0328179", URL: "http://www.imdb.com/name/nm0328179", FullName: "Gaspar 'Indio' González"},
				},
				Genres:        []string{"Crime", "Drama"},
				Languages:     []string{"Italian"},
				Nationalities: []string{"Italy", "France", "Spain"},
				AKA:           []string{"Antimetopos me tin mafia", "El hombre que desafió a la organización", "L'homme qui défia l'organisation", "One Man Against the Organization"},
			},
		},
		{
			ID: "tt0437804",
			want: Title{
				ID:       "tt0437803", // ID redirect.
				URL:      "http://www.imdb.com/title/tt0437803",
				Name:     "Alien Siege",
				Type:     "TV Movie",
				Year:     2005,
				Rating:   "3.6",
				Duration: "90m",
				Directors: []Name{
					Name{ID: "nm0821100", URL: "http://www.imdb.com/name/nm0821100", FullName: "Robert Stadd"},
				},
				Writers: []Name{
					Name{ID: "nm1305705", URL: "http://www.imdb.com/name/nm1305705", FullName: "Bill Lundy"},
					Name{ID: "nm0757507", URL: "http://www.imdb.com/name/nm0757507", FullName: "Paul Salamoff"},
				},
				Actors: []Name{
					Name{ID: "nm0424635", URL: "http://www.imdb.com/name/nm0424635", FullName: "Brad Johnson"},
					Name{ID: "nm1658541", URL: "http://www.imdb.com/name/nm1658541", FullName: "Erin Ross"},
					Name{ID: "nm0250169", URL: "http://www.imdb.com/name/nm0250169", FullName: "Lilas Lane"},
					Name{ID: "nm0003405", URL: "http://www.imdb.com/name/nm0003405", FullName: "Nathan Anderson"},
					Name{ID: "nm1108894", URL: "http://www.imdb.com/name/nm1108894", FullName: "Michael Cory Davis"},
					Name{ID: "nm1855195", URL: "http://www.imdb.com/name/nm1855195", FullName: "Gregor Paslawsky"},
					Name{ID: "nm0048844", URL: "http://www.imdb.com/name/nm0048844", FullName: "Ray Baker"},
					Name{ID: "nm0001835", URL: "http://www.imdb.com/name/nm0001835", FullName: "Carl Weathers"},
					Name{ID: "nm1197031", URL: "http://www.imdb.com/name/nm1197031", FullName: "George Zlatarev"},
					Name{ID: "nm1791963", URL: "http://www.imdb.com/name/nm1791963", FullName: "Vladimir Nikolov"},
					Name{ID: "nm1284482", URL: "http://www.imdb.com/name/nm1284482", FullName: "Atanas Srebrev"},
					Name{ID: "nm1838387", URL: "http://www.imdb.com/name/nm1838387", FullName: "Carl Ed"},
					Name{ID: "nm1838452", URL: "http://www.imdb.com/name/nm1838452", FullName: "Ivan Ivanov"},
					Name{ID: "nm1854456", URL: "http://www.imdb.com/name/nm1854456", FullName: "Zarkov Binev"},
					Name{ID: "nm1208093", URL: "http://www.imdb.com/name/nm1208093", FullName: "Julian Vergov"},
				},
				Genres:        []string{"Sci-Fi"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "Earth is attacked by the Kulkus, a hostile breed infected by a lethal virus and needing human blood to develop an antidote. Earth's governments negotiate peace terms with the Kulku ...",
				Poster:        Media{ID: "rm3287190272", TitleID: "tt0437803", URL: "http://www.imdb.com/media/rm3287190272/tt0437803", ContentURL: "http://ia.media-imdb.com/images/M/MV5BMTk1MTA4NDMwMF5BMl5BanBnXkFtZTcwNTM2MTIzMg@@._V1_SY317_CR5,0,214,317_AL_.jpg"},
				AKA:           []string{"A Föld ostroma", "Alien Blood", "Alien Siege", "Alien Siege - Tod aus dem All", "Etat de siège", "O Perigo Alienígena", "Obca krew"},
			},
		},
		{
			ID: "tt1179034",
			want: Title{
				ID:       "tt1179034",
				URL:      "http://www.imdb.com/title/tt1179034",
				Name:     "From Paris with Love",
				Year:     2010,
				Rating:   "6.5",
				Duration: "92m",
				Directors: []Name{
					Name{ID: "nm0603628", URL: "http://www.imdb.com/name/nm0603628", FullName: "Pierre Morel"},
				},
				Writers: []Name{
					Name{ID: "nm0367867", URL: "http://www.imdb.com/name/nm0367867", FullName: "Adi Hasak"},
					Name{ID: "nm0000108", URL: "http://www.imdb.com/name/nm0000108", FullName: "Luc Besson"},
				},
				Actors: []Name{
					Name{ID: "nm0000237", URL: "http://www.imdb.com/name/nm0000237", FullName: "John Travolta"},
					Name{ID: "nm0001667", URL: "http://www.imdb.com/name/nm0001667", FullName: "Jonathan Rhys Meyers"},
					Name{ID: "nm0810738", URL: "http://www.imdb.com/name/nm0810738", FullName: "Kasia Smutniak"},
					Name{ID: "nm0243983", URL: "http://www.imdb.com/name/nm0243983", FullName: "Richard Durden"},
					Name{ID: "nm0082897", URL: "http://www.imdb.com/name/nm0082897", FullName: "Bing Yin"},
					Name{ID: "nm2362244", URL: "http://www.imdb.com/name/nm2362244", FullName: "Amber Rose Revah"},
					Name{ID: "nm1101713", URL: "http://www.imdb.com/name/nm1101713", FullName: "Eric Godon"},
					Name{ID: "nm2394847", URL: "http://www.imdb.com/name/nm2394847", FullName: "François Bredon"},
					Name{ID: "nm1861574", URL: "http://www.imdb.com/name/nm1861574", FullName: "Chems Dahmani"},
					Name{ID: "nm1410527", URL: "http://www.imdb.com/name/nm1410527", FullName: "Sami Darr"},
					Name{ID: "nm3783445", URL: "http://www.imdb.com/name/nm3783445", FullName: "Julien Hagnery"},
					Name{ID: "nm0830599", URL: "http://www.imdb.com/name/nm0830599", FullName: "Mostéfa Stiti"},
					Name{ID: "nm3266168", URL: "http://www.imdb.com/name/nm3266168", FullName: "Rebecca Dayan"},
					Name{ID: "nm2457148", URL: "http://www.imdb.com/name/nm2457148", FullName: "Michaël Vander-Meiren"},
					Name{ID: "nm1359275", URL: "http://www.imdb.com/name/nm1359275", FullName: "Didier Constant"},
				},
				Genres:        []string{"Action", "Thriller"},
				Languages:     []string{"English", "French", "Mandarin", "German"},
				Nationalities: []string{"France", "USA"},
				Description:   "In Paris, a young employee in the office of the US Ambassador hooks up with an American spy looking to stop a terrorist attack in the city.",
				Poster:        Media{ID: "rm2505674496", TitleID: "tt1179034", URL: "http://www.imdb.com/media/rm2505674496/tt1179034", ContentURL: "http://ia.media-imdb.com/images/M/MV5BNDUyMzExOTAyM15BMl5BanBnXkFtZTcwMTU0NjAyMw@@._V1_SX214_AL_.jpg"},
				AKA:           []string{"Apo to Parisi me agapi", "Armastusega Pariisist", "Bons baisers de Paris", "De Paris com Amor", "Desde París con amor", "Dupla Implacável", "From Paris with Love", "Iz Pariza s ljubavlju", "MiParis be'ahava", "Paris'ten sevgilerle", "París en la Mira", "París en la mira", "Pozdrowienia z Paryza", "Párizsból szeretettel", "Sangre y amor en París", "З Парижу з любов'ю", "Из Парижа с любовью"},
			},
		},
		{
			ID: "tt0133093",
			want: Title{
				ID:       "tt0133093",
				URL:      "http://www.imdb.com/title/tt0133093",
				Name:     "The Matrix",
				Year:     1999,
				Rating:   "8.7",
				Duration: "136m",
				Directors: []Name{
					Name{ID: "nm0905152", URL: "http://www.imdb.com/name/nm0905152", FullName: "Andy Wachowski"},
					Name{ID: "nm0905154", URL: "http://www.imdb.com/name/nm0905154", FullName: "Lana Wachowski"},
				},
				Writers: []Name{
					Name{ID: "nm0905152", URL: "http://www.imdb.com/name/nm0905152", FullName: "Andy Wachowski"},
					Name{ID: "nm0905154", URL: "http://www.imdb.com/name/nm0905154", FullName: "Lana Wachowski"},
				},
				Actors: []Name{
					Name{ID: "nm0000206", URL: "http://www.imdb.com/name/nm0000206", FullName: "Keanu Reeves"},
					Name{ID: "nm0000401", URL: "http://www.imdb.com/name/nm0000401", FullName: "Laurence Fishburne"},
					Name{ID: "nm0005251", URL: "http://www.imdb.com/name/nm0005251", FullName: "Carrie-Anne Moss"},
					Name{ID: "nm0915989", URL: "http://www.imdb.com/name/nm0915989", FullName: "Hugo Weaving"},
					Name{ID: "nm0287825", URL: "http://www.imdb.com/name/nm0287825", FullName: "Gloria Foster"},
					Name{ID: "nm0001592", URL: "http://www.imdb.com/name/nm0001592", FullName: "Joe Pantoliano"},
					Name{ID: "nm0159059", URL: "http://www.imdb.com/name/nm0159059", FullName: "Marcus Chong"},
					Name{ID: "nm0032810", URL: "http://www.imdb.com/name/nm0032810", FullName: "Julian Arahanga"},
					Name{ID: "nm0233391", URL: "http://www.imdb.com/name/nm0233391", FullName: "Matt Doran"},
					Name{ID: "nm0565883", URL: "http://www.imdb.com/name/nm0565883", FullName: "Belinda McClory"},
					Name{ID: "nm0662562", URL: "http://www.imdb.com/name/nm0662562", FullName: "Anthony Ray Parker"},
					Name{ID: "nm0323822", URL: "http://www.imdb.com/name/nm0323822", FullName: "Paul Goddard"},
					Name{ID: "nm0853079", URL: "http://www.imdb.com/name/nm0853079", FullName: "Robert Taylor"},
					Name{ID: "nm0040058", URL: "http://www.imdb.com/name/nm0040058", FullName: "David Aston"},
					Name{ID: "nm0336802", URL: "http://www.imdb.com/name/nm0336802", FullName: "Marc Aden Gray"},
				},
				Genres:        []string{"Action", "Sci-Fi"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA", "Australia"},
				Description:   "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.",
				Poster:        Media{ID: "rm461886464", TitleID: "tt0133093", URL: "http://www.imdb.com/media/rm461886464/tt0133093", ContentURL: "http://ia.media-imdb.com/images/M/MV5BMTkxNDYxOTA4M15BMl5BanBnXkFtZTgwNTk0NzQxMTE@._V1_SX214_AL_.jpg"},
				AKA:           []string{"La matrice", "Ma Tran", "Maatriks", "Matrica", "Matriks", "Matrix", "Mátrix", "The Matrix", "Матрикс", "Матрица", "Матрицата"},
			},
		},
		{
			ID: "tt0291830",
			want: Title{
				ID:       "tt0291830",
				URL:      "http://www.imdb.com/title/tt0291830",
				Name:     "Corps à coeur",
				Year:     1981,
				Duration: "26m",
				Directors: []Name{
					Name{ID: "nm1029036", URL: "http://www.imdb.com/name/nm1029036", FullName: "Jean-François Després"},
					Name{ID: "nm1003289", URL: "http://www.imdb.com/name/nm1003289", FullName: "Alain Godon"},
				},
				Genres:        []string{"Documentary", "Short"},
				Nationalities: []string{"Canada"},
			},
		},
		{
			ID: "tt1965639",
			want: Title{
				ID:            "tt1965639",
				URL:           "http://www.imdb.com/title/tt1965639",
				Name:          "El clima y las cuatro estaciones",
				Type:          "TV Series",
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
				URL:      "http://www.imdb.com/title/tt0086677",
				Name:     "Brothers",
				Type:     "TV Series",
				Year:     1984,
				Rating:   "8.0",
				Duration: "22m",
				Writers: []Name{
					Name{ID: "nm0515953", URL: "http://www.imdb.com/name/nm0515953", FullName: "David Lloyd"},
					Name{ID: "nm0031246", URL: "http://www.imdb.com/name/nm0031246", FullName: "Greg Antonacci"},
				},
				Actors: []Name{
					Name{ID: "nm0907107", URL: "http://www.imdb.com/name/nm0907107", FullName: "Robert Walden"},
					Name{ID: "nm0716618", URL: "http://www.imdb.com/name/nm0716618", FullName: "Paul Regina"},
					Name{ID: "nm0535933", URL: "http://www.imdb.com/name/nm0535933", FullName: "Brandon Maggart"},
					Name{ID: "nm0865177", URL: "http://www.imdb.com/name/nm0865177", FullName: "Hallie Todd"},
					Name{ID: "nm0533383", URL: "http://www.imdb.com/name/nm0533383", FullName: "Philip Charles MacKenzie"},
					Name{ID: "nm0726939", URL: "http://www.imdb.com/name/nm0726939", FullName: "Robin Riker"},
					Name{ID: "nm0664289", URL: "http://www.imdb.com/name/nm0664289", FullName: "Mary Ann Pascal"},
				},
				Genres:        []string{"Comedy"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "Joe Waters is an ex-place kicker for the Philadelphia Eagles. Now retired, he's opened up a restaurant. Lou is his older brother, a gruff construction worker. Both Joe and Lou receive the ...",
				AKA:           []string{"Brothers", "Unter Brüdern"},
			},
		},
		{
			ID: "tt1371159",
			want: Title{
				ID:     "tt1371159",
				URL:    "http://www.imdb.com/title/tt1371159",
				Name:   "Iron Man 2",
				Type:   "Video Game",
				Year:   2010,
				Rating: "6.2",
				Directors: []Name{
					Name{ID: "nm4157448", URL: "http://www.imdb.com/name/nm4157448", FullName: "Michael McCormick"},
					Name{ID: "nm4157551", URL: "http://www.imdb.com/name/nm4157551", FullName: "Robert Taylor"},
				},
				Writers: []Name{
					Name{ID: "nm3881781", URL: "http://www.imdb.com/name/nm3881781", FullName: "Matt Fraction"},
				},
				Actors: []Name{
					Name{ID: "nm0000332", URL: "http://www.imdb.com/name/nm0000332", FullName: "Don Cheadle"},
					Name{ID: "nm0519680", URL: "http://www.imdb.com/name/nm0519680", FullName: "Eric Loomis"},
					Name{ID: "nm0000168", URL: "http://www.imdb.com/name/nm0000168", FullName: "Samuel L. Jackson"},
					Name{ID: "nm0482851", URL: "http://www.imdb.com/name/nm0482851", FullName: "Phil LaMarr"},
					Name{ID: "nm0072829", URL: "http://www.imdb.com/name/nm0072829", FullName: "John Eric Bentley"},
					Name{ID: "nm0005243", URL: "http://www.imdb.com/name/nm0005243", FullName: "Meredith Monroe"},
					Name{ID: "nm0133047", URL: "http://www.imdb.com/name/nm0133047", FullName: "Catherine Campion"},
					Name{ID: "nm0224929", URL: "http://www.imdb.com/name/nm0224929", FullName: "Dimitri Diatchenko"},
					Name{ID: "nm0149677", URL: "http://www.imdb.com/name/nm0149677", FullName: "Andrew Chaikin"},
					Name{ID: "nm0101752", URL: "http://www.imdb.com/name/nm0101752", FullName: "Doug Boyd"},
					Name{ID: "nm0946382", URL: "http://www.imdb.com/name/nm0946382", FullName: "Cedric Yarbrough"},
					Name{ID: "nm0217246", URL: "http://www.imdb.com/name/nm0217246", FullName: "Denny Delk"},
					Name{ID: "nm0413996", URL: "http://www.imdb.com/name/nm0413996", FullName: "Roger Jackson"},
					Name{ID: "nm2406242", URL: "http://www.imdb.com/name/nm2406242", FullName: "Adam Harrington"},
					Name{ID: "nm3894439", URL: "http://www.imdb.com/name/nm3894439", FullName: "Ariel Goldberg"},
				},
				Genres:        []string{"Action", "Adventure", "Sci-Fi"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
			},
		},
		{
			ID: "tt3038708",
			want: Title{
				ID:   "tt3038708",
				URL:  "http://www.imdb.com/title/tt3038708",
				Name: "Iron Sky the Coming Race",
			},
		},
	} {
		got, err := NewTitle(client, tt.ID)
		if err != nil {
			t.Errorf("NewTitle(%s) error: %v", tt.ID, err)
		} else {
			diffStruct(t, fmt.Sprintf("NewTitle(%s)", tt.ID), tt.want, *got)
		}
	}
}
