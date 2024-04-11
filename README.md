
# Sup2Recette

A brief description of what this project does and who it's for

This project is built in [Golang](https://go.dev/learn/)

## Run Locally

Clone the project

```bash
  git clone https://github.com/Vova-code/Sup2Recette
```

Go to the project directory

```bash
  cd Sup2Recette/GoApi
```

Start the server

```bash
  go run .main.go
```


## API Reference

#### Get all items

```http
  GET /recipes
```

#### Get item

```http
  GET /recipes/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of recipe |

#### Add item

```http
  POST /recipes
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Required**. Title of the new recipe |
| `steps`      | `[]string` | **Required**. Steps to achieve recipe |

#### Add a like to a recipe

```http
  POST /recipes/thums-up/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to like |

#### Add a dislike to a recipe

```http
  POST /recipes/thums-down/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to dislike |

#### Update a recipe

```http
  PUT /recipes
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of the recipe |
| `title`      | `string` | **Optionnal**. Id of item to dislike |
| `steps`      | `[]string` | **Optionnal**. Steps to achieve recipe |

At least `title` or `steps` should be present to update correctly a recipe.

#### Delete a recipe

```http
  DELETE /recipes/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of recipe |


## Database

This project stores his data on [Neon](https://neon.tech/) which is a serverless Postgres database
So there is no need to run a docker to use this app.
But you can if you prefer, just replace the connection string present in `database/postgres.go 26:13` file.

