/*
 * @lc app=leetcode id=355 lang=golang
 *
 * [355] Design Twitter
 */

// @lc code=start
type Twitter struct {
	tweetId []int
	userId []int
	follow [][2]int
}


/** Initialize your data structure here. */
func Constructor() Twitter {

}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	x
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
    
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	*Twitter.follow=append(*Twitter.follow,[followerId,followeeId])
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if val,ok:=m[[followerId,follweeId]];ok{

	}
	*Twitter.follow=append(*Twitter.follow,[followerId,followeeId])
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end

