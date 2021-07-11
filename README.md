# Clere Coding Challenge

The aim of this coding challenge is to evaluate whether you are able to build front-end software to a given set of requirements, therefore the technologies you use to accomplish this are not important. If you are already comfortable with [Vue.js](https://vuejs.org/) or similar frameworks, lean towards using these as our development team also use these technologies, we would like to emphasize that this is **not** a requirement, and you should use whatever you are most comfortable with to produce the highest quality outcome.

The use of external libraries including component, or helper libraries such as [Axios](https://www.npmjs.com/package/axios) is allowed.

You will be building a simple table view to show a list of products, along with an editor to create new products as well as update existing ones, and also the ability to delete products. Each product has an "id", "name", "price", and "currency".

We have built a RESTful API which is documented below for you to integrate into your front-end, this API implements a CRUD system to handle the state of your products. You should use the endpoints specified to populate your table, as well as modify the list of products.

# Front-end Requirements & Wireframes

To visualise the requirements of the challenge, we have provided wireframe illustrations of what the final product should look like. These are simply mockups so feel free to add as much creative flair as you like.

## Table View

![coding-challenge-table](https://user-images.githubusercontent.com/28734598/124751825-fce7e180-df1e-11eb-90a9-479c24c58aeb.png)

This page should show a button to add new products, and a table of existing products.

Clicking the "Add product" button should either take the user to a new page, or present a pop up dialog window showing the "Product Editor" form.

The table itself should have columns to show the Id, Name, Price, and Actions (buttons to edit and delete existing products), the price column values should show a concatenation of a product's "price" and "currency" (e.g. "9.99 USD").

The Actions column should contain two buttons on each row, one to edit the product, and one to delete it. Clicking the edit product button should show similar behaviour to adding a new product, however the form should already be populated with the data from the selected product, and upon saving will update the product. Clicking the delete button should delete the product and also refresh the table to reflect this.

Inclusion of pagination on the table is optional.

## Editor View

![coding-challenge-editor](https://user-images.githubusercontent.com/28734598/124751836-007b6880-df1f-11eb-9735-e98707dadc8a.png)

This page/pop up dialog should show a form with three input fields, one for "Name", "Price", and "Currency", as well as buttons to either cancel or save the state of the form.

When creating a new product, all fields should be blank. When updating an existing product, the fields should be populated with the data of the product being modified. Upon pressing save, a product will be either created or updated, then the user should be taken back to the table view. When pressing cancel, the user will simply be taken back to the table view.

# API Documentation

This API is hosted publicly on:
```
https://test.clerenet.com
```

Alternatively, if you would prefer to have an isolated environment you may clone this repo and run the API locally. You can do this by either installing golang and building a binary using the `run.sh` script:
```bash
./run.sh
```

Or if you prefer to use Docker, you can use the public docker image of this API to run a container locally:
```bash
docker run -it deltabrot/clere-coding-challenge-api:latest
```

**Get Product**
----
  Retrieves a single product by a specified id.

* **URL:**
  `/product/{id}`

* **Method:**
  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
```json
{
    "id": 1,
    "name": "Example name",
    "price": 9.99,
    "currency": "GBP"
}
```

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to retrieve product" }`

* **Sample Call:**

```bash
curl -X GET https://test.clerenet.com/product/1
```

**Get All Products**
----
  Retrieves all stored products.

* **URL:**
  `/product`

* **Method:**
  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
```json
[
    { "id": 1, "name": "Example name 1", "price": 2.99, "currency": "GBP"},
    { "id": 2, "name": "Example name 2", "price": 4.99, "currency": "GBP"},
    { "id": 3, "name": "Example name 3", "price": 9.99, "currency": "GBP"}
]
```

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to retrieve products" }`

* **Sample Call:**

```bash
curl -X GET https://test.clerenet.com/product
```

**Create Product**
----
  Creates a product.

* **URL:**
  `/product`

* **Method:**
  `POST`

* **Data Params:**
```json
{
    "name": "string",
    "price": 0,
    "currency": "string"
}
```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
```json
{
    "id": 1,
    "name": "Example name 1",
    "price": 2.99,
    "currency": "GBP"
}
```

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to create product" }`

* **Sample Call:**

```bash
curl -X POST -d '{"name":"Test","price":10.99,"currency":"USD"}' https://test.clerenet.com/product
```

**Update Product**
----
  Updates a product, it uses the provided id in the data params to determine which product to update.

* **URL:**
  `/product`

* **Method:**
  `PUT`

* **Data Params:**
```json
{
    "id": 0,
    "name": "string",
    "price": 0,
    "currency": "string"
}
```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
```json
{
    "id": 1,
    "name": "Updated name 1",
    "price": 2.99,
    "currency": "GBP"
}
```

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to update product" }`

* **Sample Call:**

```bash
curl -X PUT -d '{"id": 1,"name":"Test","price":10.99,"currency":"USD"}' https://test.clerenet.com/product
```

**Delete Product**
----
  Deletes a product by the specified id.

* **URL:**
  `/product/{id}`

* **Method:**
  `DELETE`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `empty`

* **Error Resonse:**

  * **Code:** 500 <br />
    **Content:** `{ "message": "Error: Failed to delete product" }`

* **Sample Call:**

```bash
curl -X DELETE https://test.clerenet.com/product/1
```
