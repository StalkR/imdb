package imdb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeason(t *testing.T) {
	assert := assert.New(t)
	for _, tt := range []struct {
		season int
		want   Season
	}{
		{
			season: 7,
			want: Season{
				Episodes:        []Episode{{ID: "tt1000552", Type: "tvEpisode", Season: "7", Episode: "1", TitleText: "I Fuckin' Miss Cory and Trevor", ReleaseDate: ReleaseDate{Month: 4, Day: 8, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BMWIwZWExNGUtNTJjNC00N2Q2LWFjYjItMzhhNzc2NmNlMzBhXkEyXkFqcGc@._V1_.jpg", MaxHeight: 589, MaxWidth: 1080, Caption: "Robb Wells in I Fuckin' Miss Cory and Trevor (2007)"}, Plot: "Ricky, Julian and Bubbles are stealing meat from a grocery store in an effort to make money. Julian however, inspired by the words of his 3rd grade teacher, decides to quit stealing and seek legal gainful employment, and gets a job delivering pizzas. The lack of manpower puts a cramp in Ricky&#39;s meat-stealing operation, and he begins to curse Corey and Trevor&#39;s absence. Sarah later tells the camera that Corey and Trevor are staying in a mental hospital to recover from the nervous breakdowns they suffered as a result of the constant abuse Ricky has visited upon them over the years. Ricky decides to steal a massive quantity of meat from a delicatessen, and hatches a plan that would have Bubbles hiding in the freezer until the store closes. The deli turns out to be open later than he remembers however, and Bubbles remains locked in the freezer for three hours. When the store closes, Ricky orders a pizza and when Julian arrives, the two abandon the plan to steal meat and rescue Bubbles from the freezer.", AggregateRating: 8.1, VoteCount: 482, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt1000552/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt1006252", Type: "tvEpisode", Season: "7", Episode: "2", TitleText: "I Banged Lucy and Knocked Her Up... No Big Deal", ReleaseDate: ReleaseDate{Month: 4, Day: 15, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BNDJmZDMwNGYtODc5OS00ZTQ0LTk2MGQtZjNjYzEzMTNmNDk1XkEyXkFqcGc@._V1_.jpg", MaxHeight: 589, MaxWidth: 1080, Caption: "Jonathan Torrens and John Paul Tremblay in I Banged Lucy and Knocked Her Up... No Big Deal (2007)"}, Plot: "J-Roc has set up an in-park store selling luggage stolen from the airport - Julian and Lucy are involved. George and Ted start doing surveillance on them. Ricky has concerns about who the father of Lucys baby is - Lucy finally comes clean.", AggregateRating: 7.9, VoteCount: 444, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt1006252/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt1006254", Type: "tvEpisode", Season: "7", Episode: "3", TitleText: "Three Good Men Are Dead", ReleaseDate: ReleaseDate{Month: 4, Day: 22, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BMmQyODA3OTMtMzQ0Ni00YzAyLWI0MGItYWQ4YTMxODc3OWI3XkEyXkFqcGc@._V1_.jpg", MaxHeight: 584, MaxWidth: 1080, Caption: "Mike Smith in Three Good Men Are Dead (2007)"}, Plot: "George and Ted attack Lahey for not filing the warrants they got for the Boys. Randy and Phil catch them which forces George and Teds hand. Lahey recruits the Boys to help him take George and Ted down by staging a beating on Ray.", AggregateRating: 8.9, VoteCount: 568, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt1006254/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt0992775", Type: "tvEpisode", Season: "7", Episode: "4", TitleText: "Friends of the Road", ReleaseDate: ReleaseDate{Month: 4, Day: 29, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BZWFhNGJmZDgtZDFiMi00NTliLTkwNzYtMzQ5NjFmNjhmZDJiXkEyXkFqcGc@._V1_.jpg", MaxHeight: 584, MaxWidth: 1080, Caption: "John Paul Tremblay and Robb Wells in Friends of the Road (2007)"}, Plot: "Ricky and Julian have to take Bubbles to the train convention after Ray gets picked up for prostitution in his truck. The Boys make a new friend in Sebastian Bach. Bubbles comes home from the convention with an expense present.", AggregateRating: 8.4, VoteCount: 461, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0992775/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt1006253", Type: "tvEpisode", Season: "7", Episode: "5", TitleText: "The Mustard Tiger", ReleaseDate: ReleaseDate{Month: 5, Day: 6, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BMDlhZGViNTEtYjRlMS00OTYxLWE1YmUtM2E4ODFlODhhYzNjXkEyXkFqcGc@._V1_.jpg", MaxHeight: 588, MaxWidth: 1080, Caption: "John Paul Tremblay and Mike Smith in The Mustard Tiger (2007)"}, Plot: "Julian has a plan for transporting marijuana over the border - model train track. Jacob and two of his friends go missing when the Boys send them to lay the track. Phil and Thomas Collins press the Boys for information.", AggregateRating: 8.2, VoteCount: 445, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt1006253/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt1026371", Type: "tvEpisode", Season: "7", Episode: "6", TitleText: "We Can't Call People Without Wings Angels So We Call Them Friends", ReleaseDate: ReleaseDate{Month: 5, Day: 13, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BZDI4Nzc3MTMtNjM0OC00ZDJiLTgzMzctNTZjZGExOTM4MGY3XkEyXkFqcGc@._V1_.jpg", MaxHeight: 582, MaxWidth: 1080, Caption: "John Paul Tremblay, Mike Smith, and Robb Wells in We Can't Call People Without Wings Angels So We Call Them Friends (2007)"}, Plot: "The boys try to find Jacob and the crew, who have been missing for some time.", AggregateRating: 8.2, VoteCount: 444, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt1026371/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt1026368", Type: "tvEpisode", Season: "7", Episode: "7", TitleText: "Jump the Cheeseburger", ReleaseDate: ReleaseDate{Month: 5, Day: 20, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BZGUwMGVmNGYtZWJmMS00MTVjLTg0YjUtODMyMWRkMTEzODYzXkEyXkFqcGc@._V1_.jpg", MaxHeight: 585, MaxWidth: 1080, Caption: "Sam Tarasco in Jump the Cheeseburger (2007)"}, Plot: "Phil and Randy are preparing for the opening of the Dirty Burger.", AggregateRating: 7.9, VoteCount: 453, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt1026368/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt1026369", Type: "tvEpisode", Season: "7", Episode: "8", TitleText: "Let the Liquor do the Thinking", ReleaseDate: ReleaseDate{Month: 5, Day: 27, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BNGQ1MjljZTAtNmVmMS00NTAwLWJkODktYzZiN2Q0YWIyZjhhXkEyXkFqcGc@._V1_.jpg", MaxHeight: 587, MaxWidth: 1080, Caption: "Mike Smith in Let the Liquor do the Thinking (2007)"}, Plot: "Ray sends Bubbles a wanted poster from the US with his face on it; Bubbles starts getting paranoid. Sebastian Bach informs the Boys of a forest fire threatening the operation. Lahey uses the Swayze express as leverage against the Boys.", AggregateRating: 8.2, VoteCount: 456, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt1026369/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt1026367", Type: "tvEpisode", Season: "7", Episode: "9", TitleText: "Going Off the Rails on a Swayze Train", ReleaseDate: ReleaseDate{Month: 6, Day: 3, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BYWY2Mzc1ZTUtODA2Ny00YTkzLTk1MGItM2ZkNWUwNTIyZTRiXkEyXkFqcGc@._V1_.jpg", MaxHeight: 592, MaxWidth: 1080, Caption: "Mike Smith and Conky in Going Off the Rails on a Swayze Train (2007)"}, Plot: "Forest Rangers close in on the Boys. Lahey and Randy pursue the Swayze express, and recruit Jacob to wear a wire to catch the Boys in the act. Conkys return helps Bubbles sanity but causes nothing but trouble with Ricky.", AggregateRating: 8.3, VoteCount: 445, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt1026367/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt1026366", Type: "tvEpisode", Season: "7", Episode: "10", TitleText: "A Shit River Runs Through It", ReleaseDate: ReleaseDate{Month: 6, Day: 10, Year: 2007, Typename: "ReleaseDate"}, ReleaseYear: 2007, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BMTE5MWVlNGYtNWViMS00NWE4LWE3NjgtMDgzMjcyMDQ0M2YxXkEyXkFqcGc@._V1_.jpg", MaxHeight: 581, MaxWidth: 1080, Caption: "John Dunsworth and Patrick Roach in A Shit River Runs Through It (2007)"}, Plot: "Lahey and Randy, the Forest Rangers, the ATF, the DEA, and the FBI all close in on the Boys, and Bubbles and Conky make the situation worse for everyone. Can the Boys emerge unscathed?", AggregateRating: 9.1, VoteCount: 695, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt1026366/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}},
				Total:           10,
				HasNextPage:     false,
				EndCursor:       "dHQxMDI2MzY2",
				HasRatedEpisode: true,
			},
		},
		{
			season: 5,
			want: Season{
				Episodes:        []Episode{{ID: "tt0732907", Type: "tvEpisode", Season: "5", Episode: "1", TitleText: "Give Peace a Chance", ReleaseDate: ReleaseDate{Month: 4, Day: 17, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BNzMyNTlhZGUtOGQxYy00ZWNhLWJmZmEtZmEyYTU3YmYxNWI1XkEyXkFqcGc@._V1_.jpg", MaxHeight: 739, MaxWidth: 986, Caption: "John Paul Tremblay, Mike Smith, and Robb Wells in Give Peace a Chance (2005)"}, Plot: "The boys return from a stint in jail to discover that Cory and Trevor have squandered all of the pot profits; Lahey&#39;s time in a mental institution seems to have done him some good.", AggregateRating: 8.3, VoteCount: 540, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732907/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, {ID: "tt0732928", Type: "tvEpisode", Season: "5", Episode: "2", TitleText: "The Shit Puppets", ReleaseDate: ReleaseDate{Month: 4, Day: 24, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BZDczMTBiNmYtZjdhMi00ZDM5LTkxODYtZDdmMjE1NzE2MmM5XkEyXkFqcGc@._V1_.jpg", MaxHeight: 737, MaxWidth: 985, Caption: "John Paul Tremblay, Mike Smith, and Robb Wells in The Shit Puppets (2005)"}, Plot: "Julian takes Cory and Trevor to work as a way for them to pay back the pot-harvest money they blew. The boys are surprised to discover that the work involves using guns.", AggregateRating: 8.6, VoteCount: 541, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732928/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt0732924", Type: "tvEpisode", Season: "5", Episode: "3", TitleText: "The Fuckin' Way She Goes", ReleaseDate: ReleaseDate{Month: 5, Day: 1, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BM2YxYTI4NTYtNmJhOS00MGQ0LTg1OGEtZmU5OGI4ODlmOGQ3XkEyXkFqcGc@._V1_.jpg", MaxHeight: 762, MaxWidth: 1016, Caption: "Robb Wells in The Fuckin' Way She Goes (2005)"}, Plot: "Julian, Ricky, and the boys hide 100 kilos of hash in plain sight. Something unexpected gets burned up due to a celebratory barbecue.", AggregateRating: 8.5, VoteCount: 575, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732924/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt0732935", Type: "tvEpisode", Season: "5", Episode: "4", TitleText: "You Got to Blame the Thing Up Here", ReleaseDate: ReleaseDate{Month: 5, Day: 8, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BODAxMWYyYzQtZWU1Yi00ZjExLTg3NWYtZDMzMWZjOTU0YmRmXkEyXkFqcGc@._V1_.jpg", MaxHeight: 770, MaxWidth: 1027, Caption: "John Paul Tremblay in You Got to Blame the Thing Up Here (2005)"}, Plot: "Bubbles takes the blame for something Ricky did. Julian uses gumball machines to finance a money-making scheme.", AggregateRating: 8.3, VoteCount: 510, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732935/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt0732913", Type: "tvEpisode", Season: "5", Episode: "5", TitleText: "Mr. Lahey Is a Fuckin' Drunk, and He Always Will Be", ReleaseDate: ReleaseDate{Month: 5, Day: 15, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BZWRlMGNjZDMtNjEyNi00YTNhLWE2ZDUtMTU5NTg5ZmJhNjlkXkEyXkFqcGc@._V1_.jpg", MaxHeight: 788, MaxWidth: 1053, Caption: "Mike Smith and Robb Wells in Mr. Lahey Is a Fuckin' Drunk, and He Always Will Be (2005)"}, Plot: "Ricky&#39;s feeling a tiny bit bad about the fire he caused. Randy&#39;s worried because Lahey is constantly talking about not drinking.", AggregateRating: 8.6, VoteCount: 545, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732913/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt0732904", Type: "tvEpisode", Season: "5", Episode: "6", TitleText: "Don't Cross the Shitline", ReleaseDate: ReleaseDate{Month: 5, Day: 22, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BYTI1ZTcwYzMtOWNlMS00OTQ5LWJkZTYtYzhkMmRiZTA5YmM3XkEyXkFqcGc@._V1_.jpg", MaxHeight: 1080, MaxWidth: 1444, Caption: "Barrie Dunn, John Dunsworth, Mike Smith, Robb Wells, and Patrick Roach in Trailer Park Boys (2001)"}, Plot: "J-Roc&#39;s porn filming leads to unexpected complications.", AggregateRating: 8.3, VoteCount: 521, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732904/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt0732929", Type: "tvEpisode", Season: "5", Episode: "7", TitleText: "The Winds of Shit", ReleaseDate: ReleaseDate{Month: 5, Day: 29, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BYWFkNGViNmUtMTNiOS00ZDY1LThkYmItYTI5ZjBkZWNkODg2XkEyXkFqcGc@._V1_.jpg", MaxHeight: 773, MaxWidth: 1033, Caption: "John Paul Tremblay, Mike Smith, and Robb Wells in The Winds of Shit (2005)"}, Plot: "Lahey thinks up a plan that will put Julian back in jail.", AggregateRating: 8.5, VoteCount: 538, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732929/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt0732905", Type: "tvEpisode", Season: "5", Episode: "8", TitleText: "Dressed All Over & Zesty Mordant", ReleaseDate: ReleaseDate{Month: 6, Day: 5, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BNmFkNTYzNDYtNDYzOC00OWVmLWE5MmItNzVmOTA3ZTk5MjJlXkEyXkFqcGc@._V1_.jpg", MaxHeight: 1080, MaxWidth: 1444, Caption: "John Paul Tremblay and Mike Smith in Trailer Park Boys (2001)"}, Plot: "The boys prepare to ship hash in shopping carts, under Lahey&#39;s secret surveillance.", AggregateRating: 8.2, VoteCount: 503, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732905/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt0732908", Type: "tvEpisode", Season: "5", Episode: "9", TitleText: "I Am the Liquor", ReleaseDate: ReleaseDate{Month: 6, Day: 12, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BYzdhMmUwZTEtMGY2Zi00ZTQ4LTg3Y2QtMjZlYjI4ZmIxY2NkXkEyXkFqcGc@._V1_.jpg", MaxHeight: 782, MaxWidth: 1044, Caption: "Bernard Robichaud, Nobu Adilman, and Mio Adilman in I Am the Liquor (2005)"}, Plot: "Julian makes payments towards his trailer; he wants it back before bailing Ray out of jail. Ricky isn&#39;t happy about this when he finds out. Lahey bails out Cyrus, Dennis and Terry so that they will seek to get their hash back from the Boys.", AggregateRating: 8.9, VoteCount: 607, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732908/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}, Episode{ID: "tt0732927", Type: "tvEpisode", Season: "5", Episode: "10", TitleText: "The Shit Blizzard", ReleaseDate: ReleaseDate{Month: 6, Day: 19, Year: 2005, Typename: "ReleaseDate"}, ReleaseYear: 2005, Image: Image{URL: "https://m.media-amazon.com/images/M/MV5BOTM5ODAyNWUtMmJmNC00NTRmLTkxOTktNDQ5ZTJhOWM1NTA1XkEyXkFqcGc@._V1_.jpg", MaxHeight: 589, MaxWidth: 1080, Caption: "John Dunsworth and Patrick Roach in The Shit Blizzard (2005)"}, Plot: "Cyrus, Dennis, and Terry get released from jail and return to Sunnyvale to try to get their hash from the boys.", AggregateRating: 9, VoteCount: 613, CanRate: true, ContributionUrl: "https://contribute.imdb.com/image/tt0732927/add?bus=imdb&return_url=https%3A%2F%2Fwww.imdb.com%2Fclose_me&site=web"}},
				Total:           10,
				HasNextPage:     false,
				EndCursor:       "dHQwNzMyOTI3",
				HasRatedEpisode: true,
			},
		},
	} {
		got, err := NewSeason(client, "tt0290988", tt.season)
		if err != nil {
			t.Errorf("NewSeason(\"tt0290988\", %d) error: %v", tt.season, err)
			continue
		}
		assert.Equal(tt.want, *got, "NewSeason(\"tt0290988\", %d) error: %v", tt.season, err)
	}
}
