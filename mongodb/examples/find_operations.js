// Basic find operations
db.users.find()
db.users.find({})
db.users.findOne()
db.users.findOne({})

// Find with filter
db.users.find({ name: "alice" })
db.users.find({ age: 25 })
db.users.findOne({ name: "bob" })

// Find with query operators
db.users.find({ age: { $gt: 25 } })
db.users.find({ age: { $gte: 18, $lt: 65 } })
db.users.find({ status: { $in: ["active", "pending"] } })
db.users.find({ $or: [{ name: "alice" }, { name: "bob" }] })
db.users.find({ tags: { $all: ["mongodb", "database"] } })

// Find with nested documents
db.users.find({ "address.city": "New York" })
db.users.find({ profile: { name: "test", active: true } })

// Find with array fields
db.users.find({ tags: "mongodb" })
db.users.find({ scores: { $elemMatch: { $gt: 80, $lt: 90 } } })
