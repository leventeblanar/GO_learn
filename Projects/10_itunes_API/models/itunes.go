package models

type ITunesResponse struct {
	ResultCount		int 		`json:"resultCount"`
	Results			[]Track		`json:"results"`
}

type Track struct {
	TrackName		string 		`json:"trackName"`
	ArtistName		string		`json:"artistName"`
	CollectionName	string		`json:"collectionName"`
}


