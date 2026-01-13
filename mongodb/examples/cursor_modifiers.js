// Cursor modifiers
db.users.find().sort({ age: -1 })
db.users.find().limit(10)
db.users.find().skip(5)
db.users.find().projection({ name: 1, age: 1 })
db.users.find().project({ name: 1, email: 1 })

// Chained modifiers
db.users.find().sort({ age: -1 }).limit(10)
db.users.find().sort({ createdAt: -1 }).skip(20).limit(10)
db.users.find({ status: "active" }).sort({ name: 1 }).limit(100).skip(0)

// With projection
db.users.find().sort({ age: 1 }).projection({ name: 1, age: 1, _id: 0 })

// Complex query with all modifiers
db.users.find({ age: { $gt: 18 } }).sort({ lastName: 1, firstName: 1 }).skip(10).limit(20).projection({ firstName: 1, lastName: 1, email: 1 })
