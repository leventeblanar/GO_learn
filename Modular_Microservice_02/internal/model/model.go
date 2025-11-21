package model

type Track struct {
	TrackId 		int
	Name			string
	Milliseconds	int
	AlbumId			int
	MediaTypeId		int
	UnitPrice		float64
}

type Album struct {
	AlbumId			int
	Title			string
	Tracks			[]Track
}

type AlbumStats struct {
	AlbumId			int
	Title			string
	TrackCount		int
	TotalDuration	float64
	AverageDuration	float64
}