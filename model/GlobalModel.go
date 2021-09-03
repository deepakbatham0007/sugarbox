package model

type SearchName struct{
    FindItem string `json:"findItem"`
}

type SearchUserItems struct{
    Userid string `json:"Userid"`
}

type SearchUserResults struct{
	Name string `json:"Name"`
	Type string `json:"Type"`
	Director string `json:"Director"`
	Artist string `json:"Artist"`
	Rating int `json:"Rating"`
	Comment string `json:"Comment"`
}

type SearchRes struct{
	//Name sql.NullString `json:"Name"`
	Name string `json:"Name"`
	Type string `json:"Type"`
	Director string `json:"Director"`
	Artist string `json:"Artist"`
	Rating int `json:"Rating"`
	NoOfRatedUser int `json:"NoOfRatedUser"`
}


type RatingModel struct{
	UserId string `json:"UserId"`
	Name string `json:"Name"`
	Rating int `json:"Rating"`
}

type CommentModel struct{
	UserId string `json:"UserId"`
	Name string `json:"Name"`
	Comment string `json:"Comment"`
}