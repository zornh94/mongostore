var database = "poc";
db = db.getSiblingDB(database);
db.dropDatabase();

entityA =[
    { _id:"a0", item:"item0" },
    { _id:"a1", item:"item1" },
    { _id:"a2", item:"item2" },
    { _id:"a3", item:"item3" }
];

entityB = [
    {_id:"b0", name:"name0"},
    {_id:"b1", name:"name1"},
    {_id:"b2", name:"name2"},
    {_id:"b3", name:"name3"}
];

collection1 = "entityA"
collection2 = "entityB"

entityA.forEach((o,i)=> {
    print(o.item +" ("+db[collection1].insertOne(o).insertedId+") test1 have created. ");
});

print("\n")
entityB.forEach((o,i)=> {
    print(o.name +" ("+db[collection2].insertOne(o).insertedId+") test1 have created. ");
});
