package usersclient

type User struct {
	ID       string `json:"_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Alias    string `json:"alias,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type UpdateResult struct {
	MatchedCount  int `json:"matchedCount,omitempty"`
	ModifiedCount int `json:"modifiedCount,omitempty"`
	UpsertedCount int `json:"upsertedCount,omitempty"`
}

type DeleteResult struct {
	DeletedCount int `json:"deletedCount,omitempty"`
}

type InsertedResult struct {
	InsertedID string `json:"insertedID,omitempty"` //example "61dd1635b9fd2fb647c16b09"
}
