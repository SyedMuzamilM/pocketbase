migrate((db) => {
  const collection = new Collection({
    "id": "sksyfpchjdnjuu3",
    "created": "2022-12-22 08:38:22.752Z",
    "updated": "2022-12-22 08:38:22.752Z",
    "projectName": "test",
    "name": "books_clone_from_url",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "7nmusoas",
        "name": "author",
        "type": "text",
        "required": true,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "sk4jfivv",
        "name": "name",
        "type": "text",
        "required": true,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      }
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("sksyfpchjdnjuu3");

  return dao.deleteCollection(collection);
})
