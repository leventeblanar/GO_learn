package model

type Cikk struct {
	ID				int
	Nev				string
	Cikkszam		string
	CikkTipusId		int
}

type CikkTipus struct {
	ID				int
	Nev				string
}