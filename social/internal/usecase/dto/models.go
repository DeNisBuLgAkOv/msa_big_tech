package dto

type FriendRequestDto struct {
	FromUserID string
	ToUserID   string
}

type ChangeFriendRequestDto struct {
	UserID    string
	RequestID string
}

type ListFriendsDto struct {
	UsersID string
	Limit   int64
	Cursor  *string
}

type ListFriendsResponse struct {
	Friends    []string
	NextCursor *string
}
