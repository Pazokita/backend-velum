package maps

// Simule une base de données avec des cartes
var mapList = []MapMetadata{
	{
		ID:       1,
		Title:    "Paris 1750",
		ImageURL: "http://localhost:8080/static/paris-1750.png",
		BBox:     [4]float64{2.31, 48.85, 2.37, 48.89},
		Opacity:  0.8,
	},
	{
		ID:       2,
		Title:    "Paris 1850",
		ImageURL: "http://localhost:8080/static/paris-1850.png",
		BBox:     [4]float64{2.29, 48.84, 2.38, 48.90},
		Opacity:  0.6,
	},
}

// GetAllMaps retourne toutes les cartes disponibles
func GetAllMaps() []MapMetadata {
	return mapList
}
