
# EFishery Backend Test
## Authors

- [Luthfi Yufajjiru](https://www.github.com/luthfiyufajjiru)


## How to get started

- Make sure golang and docker is on your computer.
- Use Visual Studio Code to open both app in one workspace.
- Install Golang official extension to debug.
- After clone this repository, open `services.code-workspace` with VsCode.
- There is docker-compose.yml and docker-compose-prod.yml.
- Run this command in the terminal: `docker compose -f docker-compose-prod.yml up authapp -d ; docker compose -f docker-compose.yml up fetchapp -d`
- That command means you are starting auth app as a compiled binary, and the fetchapp is able to debug.
- You can swap between fetchapp or authapp for that command.
- The api specification is served by postman and you can import with `api-spec.json`.
## API Reference Overview

### Auth App

#### Sign Up

```http
  POST /api/v1/account/sign-up
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `phone_number` | `string` | **Required**. Starts with international code, example: +62 |
| `name` | `string` | **Required**. Not null. |
| `user_role` | `string` | User role |

Then you would get a four chars random generated password.

#### Sign In

```http
  POST /api/v1/account/sign-in
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `phone_number` | `string` | **Required**. Starts with international code, example: +62 |
| `password` | `string` | **Required**. Not null. |

Then you would get a string of JWT token.

#### Claim

```http
  GET /api/v1/account/claim
```
| Header | Description                |
| :------| :------------------------- |
| `Token` | **Required**. Place your token after sign in here. |

Then you would get private claim if valid

| Field | Type     |
| :-------- | :------- |
| `phone_number` | `string`|
| `name` | `string` |
| `user_role` | `string` |
| `timestamp` | `int64` |

### Fetch App

#### Get all items

```http
  GET /api/v1/items
```

| Header | Description                |
| :------| :------------------------- |
| `Token` | **Required**. Place your token after sign in here. |

#### Get the agregate

```http
  GET /api/items/agregate
```

| Header | Description                |
| :------| :------------------------- |
| `Token` | **Required**. Place your token after sign in here. |