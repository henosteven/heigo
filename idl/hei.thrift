namespace go heiThrift

struct User {
    1: i64  userId
    2: string UserName
}

service UserHandler {
    User getUser(1: i64 userId)
    i64  addUser(1: string username)
}