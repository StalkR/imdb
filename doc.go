/*
Package imdb implements IMDb web API.

All operations require an http client such as:
	client := http.DefaultClient

To search a title:
	results, err := imdb.SearchTitle(client, "matrix")
	...
results is a slice of imdb.Title results with basic information (Name, URL, Year).

To get detailed information on a title:
	title, err := imdb.NewTitle(client, "tt0133093")
	...
Actors, Rating, Description and other fields are available.
*/
package imdb
