# Project Title

MetaGym Web Forum REST API (BackEnd)

## Demo link:

To be deployed soon.

## Table of Content:

- [About The App](#about-the-app)
- [Screenshots](#screenshots)
- [Technologies](#technologies)
- [Setup](#setup)
- [Status / Features](#status)
- [Credits](#credits)
- [License](#license)

## About The App

**[MetaGym Web Forum API]** is a backend RESTful API that is built using Golang. This api is to mainly serve the MetaGym Web Forum (Frontend). For more information about the app check the [MetaGym Web Forum (FrontEnd) Repository](https://github.com/jmsandiegoo/metagym_web_forum_frontend).

## Technologies

**Language:**\
Golang

**Server Framework & Libraries:**\
Gin Web Framework\
GORM

**Database:**\
PostgreSQL

**Testing Tool:**\
Postman

## Setup

- First setup and run PostgreSQL database (this project uses version 15)
- Create a database name `'metagym-forum-db'` under the user `postgres` (feel free to change it as long it is the same as the one in `.env.local` file)
- Run this insert sql script to pre populate the database with the interest data that would act as 'category' for the threads and 'interests' for the users.

```
INSERT INTO public.interests(
	id, name, created_at, updated_at)
	VALUES ('990145d4-6bef-4e2b-a1a3-232b3de3efbf', 'Gym', '2023-01-03 23:51:59.872447+08', '2023-01-03 23:51:59.872447+08');

INSERT INTO public.interests(
	id, name, created_at, updated_at)
	VALUES ('8eed2e94-7d65-49c2-ac6e-2bece605dc0b', 'Calisthenics', '2023-01-03 23:51:59.872447+08', '2023-01-03 23:51:59.872447+08');

INSERT INTO public.interests(
	id, name, created_at, updated_at)
	VALUES ('739b250c-2c84-4f05-8860-db9d1c7158d5', 'Powerlifting', '2023-01-03 23:51:59.872447+08', '2023-01-03 23:51:59.872447+08');

INSERT INTO public.interests(
	id, name, created_at, updated_at)
	VALUES ('bf1173b9-0105-4837-91e8-06a1abc1d90d', 'Nutrition', '2023-01-03 23:51:59.872447+08', '2023-01-03 23:51:59.872447+08');

INSERT INTO public.interests(
	id, name, created_at, updated_at)
	VALUES ('c2012c99-26c9-4f16-8a4f-b3b1b1fbfc43', 'Education', '2023-01-03 23:51:59.872447+08', '2023-01-03 23:51:59.872447+08');

INSERT INTO public.interests(
	id, name, created_at, updated_at)
	VALUES ('a7792f04-7e1e-4b74-a3fe-d4af92219519', 'Running', '2023-01-03 23:51:59.872447+08', '2023-01-03 23:51:59.872447+08');

INSERT INTO public.interests(
	id, name, created_at, updated_at)
	VALUES ('c55c49e6-f9f9-4f09-a469-20b7678b872d', 'Tips & Tricks', '2023-01-03 23:51:59.872447+08', '2023-01-03 23:51:59.872447+08');
```

- After setting up and running the database download or clone this repository
- create a `.env` and a `.env.local` file in the root project directory.
  In the `.env` file add the following

```
# Database credentials
DB_HOST="<<DB_HOST>>"
DB_USER="<<DB_USER>>"
DB_PASSWORD="<<DB_PASSWORD>>"
DB_NAME="diary_app"
DB_PORT="<<DB_PORT>>"

# Authentication credentials
TOKEN_TTL="2000"
JWT_PRIVATE_KEY="THIS_IS_NOT_SO_SECRET+YOU_SHOULD_DEFINITELY_CHANGE_IT"

# CORS
FRONTEND_URL="<<URL_OF_FRONTEND>>"
```

In the `.env.local` file add the values this time:

```
# Database credentials (change according to your db setup)
DB_HOST="127.0.0.1"
DB_USER="postgres"
DB_PASSWORD="myawesomepassword"
DB_NAME="metagym-forum-db"
DB_PORT="5432"

# Authentication credentials
TOKEN_TTL="2000"
JWT_PRIVATE_KEY="THIS_IS_NOT_SO_SECRET+YOU_SHOULD_DEFINITELY_CHANGE_IT"
DOMAIN="localhost"

# CORS
FRONTEND_URL="http://localhost:3000"
```

- run `go run cmd/server/main.go` to start the api server, assuming the database server is already running.

## Status / Features

**[MetaGym Web Forum API]** currently has the following features:

Note: Attempted to implement transaction and locking to prevent race condition and to keep things ATOMIC. Validation of request data is still in development as well.

**Authentication Endpoints (JWT)**\
&emsp;- Login / Signup\
&emsp;- Get Current Authenticated User\
&emsp;- Forgot Password (In development...)

**Thread Endpoints**\
&emsp;- Get Threads\
&emsp;- Create Thread\
&emsp;- Update Thread\
&emsp;- Delete Thread

**Comment Endpoints**\
&emsp;- Get Thread Comments\
&emsp;- Create Thread Comment\
&emsp;- Update Thread Comment\
&emsp;- Delete Thread Comment

**Interests Endpoint**\
&emsp;- Get All Interests

**Search Endpoint**\
&emsp;- Search a Thread w/ Interests & Title Keyword

**Upvote / Downvote Endpoints**\
&emsp;- Upvote Thread\
&emsp;- Downvote Thread\
&emsp;- Upvote Comment\
&emsp;- Downvote Comment

**User Endpoints**\
&emsp;- Get User Details\
&emsp;- Onboard New User\
&emsp;- Update Profile and Account Endpoint (In development...)

`More features` will be out in the future like:\
&emsp;- Image support\
&emsp;- Thread Body Formatting\
&emsp;- Pagination & Sorting of Thread and Comment List\
&emsp;- Search Users
and more...

## Credits

List of contriubutors:

- [Jm San Diego](https://github.com/jmsandiegoo)

## License

MIT license
