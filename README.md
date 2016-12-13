The goal of this project is to develop the DevOps best practices.

It starts from a store.

# Add a new item
curl -XPOST http://localhost :8094/store/items -H "Content-Type: application/json" -d '{"brand":"apple", "name":"iPhone7", "description":"The latest iphone"}' —verbose
# Get an existing item
curl http://localhost:8094/store/items/d9494812-6154-4ac6-9177-5e6ee3eb648b
