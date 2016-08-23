docker run -t -i -v /data:/data --name mongodb1 -d mongo
docker run -t -i -v /data:/data -p 27017:27017 --name mongodb2 -d mongo


docker exec -it mongodb2 bash
mongo –host localhost –port 27017



show dbs;
use mongodb;
show dbs;
db.createCollection("catalog")
show collections;

db.createCollection("catalog_capped", {capped: true, autoIndexId: true, size: 64 * 1024, max: 1000} )
db.runCommand( { create: "catalog_capped_2", capped: true, size: 64 * 1024, max: 1000 } )

db.catalog.count()


doc1 = {"catalogId" : "catalog1", "journal" : 'Oracle Magazine', "publisher" : 'Oracle Publishing', "edition" : 'November December 2013',"title" : 'Engineering as a Service',"author" : 'David A. Kelly'}
db.catalog.insert(doc1)
db.catalog.count()
db.catalog.find()


doc2 = {"_id": ObjectId("507f191e810c19729de860ea"), "catalogId" : "catalog1", "journal" : 'Oracle Magazine', "publisher" : 'Oracle Publishing', "edition" : 'November December 2013',"title" : 'Engineering as a Service',"author" : 'David A. Kelly'};

db.catalog.findOne()
db.catalog.findOne({},{edition:1, title:1, author:1})
db.catalog.drop()


#updating a document
db.catalog.save()
#outputting documents as JSON
db.catlog.find().forEach(printjson)


#backupo database
mongodump --db test --out /data/backup
#restor/revovery database from backup
mongoretore --db testrestore /data/backup/test


#connect by ip:port
mongo localhost:27017/testrestore
show dbs;
use mongodb
show collections
db.catalog.find()

#remove documents in mongo
db.catalog.remove({ _id: ObjectId("561ff033380a18f6587b0aa5") })
#remove all documents in mongo
db.catalog.remove({})
