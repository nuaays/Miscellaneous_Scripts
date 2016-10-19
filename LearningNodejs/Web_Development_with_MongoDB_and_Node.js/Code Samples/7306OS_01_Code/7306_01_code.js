$ mongo
> db.contacts.find({email: ‘jason@kroltech.com’}).pretty()

{
    "email" : "jason@kroltech.com",
    "phone" : "123-456-7890",
    "gravatar" : "751e957d48e31841ff15d8fa0f1b0acf",
    "_id" : ObjectId("52fad824392f58ac2452c992"),
    "name" : {
        "first" : "Jason",
        "last" : "Krol"
    },
    "__v" : 0
}