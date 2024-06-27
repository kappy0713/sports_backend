package domain

type Relationships struct {
	ID         string `json:"id" gorm:"primaryKey"`
	FollowedID string `json:"followed_id"`
	FollowerID string `json:"follower_id"`
	Follower   User   `gorm:"foreignKey:FollowerID"`
	Followed   User   `gorm:"foreignKey:FollowedID"`
}

type RelationshipsResponse struct {
	ID         string `json:"id" gorm:"primaryKey"`
	FollowedID string `json:"followed_id"`
	FollowerID string `json:"follower_id"`
	Follower   User
	Followed   User
}
