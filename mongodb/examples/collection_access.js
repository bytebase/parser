// Collection access patterns
db.users.find()
db["users"].find()
db['users'].find()
db.getCollection("users").find()
db.getCollection('users').find()

// Collection names with special characters
db["user-logs"].find()
db["my.collection"].find()
db.getCollection("user-events").find()

// Method on getCollectionNames
db.getCollectionNames()
