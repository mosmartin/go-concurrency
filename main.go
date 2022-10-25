package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	userID := 5

	go fetchUserData(userID)
	go fetchUserRecommendations(userID)
	go fetchUserFriends(userID)
	go fetchUserPosts(userID)
	go fetchUserComments(userID)
	go fetchUserLikes(userID)

	fmt.Println(time.Since(now))
}

func fetchUserData(userID int) {
	time.Sleep(50 * time.Millisecond)
}

func fetchUserRecommendations(userID int) {
	time.Sleep(100 * time.Millisecond)
}

func fetchUserFriends(userID int) {
	time.Sleep(150 * time.Millisecond)
}

func fetchUserPosts(userID int) {
	time.Sleep(100 * time.Millisecond)
}

func fetchUserComments(userID int) {
	time.Sleep(200 * time.Millisecond)
}

func fetchUserLikes(userID int) {
	time.Sleep(50 * time.Millisecond)
}
