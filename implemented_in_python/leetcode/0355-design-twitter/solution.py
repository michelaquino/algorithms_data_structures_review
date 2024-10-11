# https://leetcode.com/problems/design-twitter/description/


class User:
    def __init__(self, id):
        self.id = id
        self.followees = {}
        self.feed = []  # (timestamp, id)

    def addFollowee(self, user):
        self.followees[user.id] = user

    def removeFollowee(self, userId):
        if userId not in self.followees:
            return

        del self.followees[userId]

    def addPost(self, time, tweetId):
        heapq.heappush(self.feed, (time * -1, tweetId))

    def getTenPosts(self):
        return heapq.nsmallest(10, self.feed)

    def getFeed(self):
        posts = self.getTenPosts()

        for followee in self.followees.values():
            posts += followee.getTenPosts()

        heapq.heapify(posts)
        feed = []

        for post in heapq.nsmallest(10, posts):
            feed.append(post[1])

        return feed


class Twitter:
    def __init__(self):
        self.users = {}
        self.postCount = 0

    def postTweet(self, userId: int, tweetId: int) -> None:
        self.postCount += 1
        user = self.users.get(userId, User(userId))
        user.addPost(self.postCount, tweetId)
        self.users[userId] = user

    def getNewsFeed(self, userId: int) -> List[int]:
        if userId not in self.users:
            return []

        return self.users[userId].getFeed()

    def follow(self, followerId: int, followeeId: int) -> None:
        follower = self.users.get(followerId, User(followerId))
        followee = self.users.get(followeeId, User(followeeId))
        follower.addFollowee(followee)

        self.users[followerId] = follower
        self.users[followeeId] = followee

    def unfollow(self, followerId: int, followeeId: int) -> None:
        if followerId not in self.users:
            return

        follower = self.users[followerId]
        follower.removeFollowee(followeeId)


# Your Twitter object will be instantiated and called as such:
# obj = Twitter()
# obj.postTweet(userId,tweetId)
# param_2 = obj.getNewsFeed(userId)
# obj.follow(followerId,followeeId)
# obj.unfollow(followerId,followeeId)
