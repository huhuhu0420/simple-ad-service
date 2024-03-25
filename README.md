# Simple Ad Service (SADs 😢)

This is a simple backend service which can create and get ads.

### Quick Start

```sh
# Clone the repo
git clone https://github.com/huhuhu0420/simple-ad-service
# Change directory
cd simple-ad-service
# Copy env file
cp example.env .env
# Start the service
docker-compose up
```

There're two containers will be created, one is Golang for backend server, another is postgreSQL for DBMS.

Database will be initialized with some data, you can check the data in `sql/*.sql`.

Go to http://localhost:5000/swagger/index.html can see swagger api document.

### Project Explanation

#### Golang Server

The server is built with Golang, and it uses `gin` as the web framework. The server provides two APIs:
`/ad POST`: create an ad
- the request body should be a JSON object, and the JSON object should contain the following fields:
  - title: string
  - startAt: string (in the format of "YYYY-MM-DD")
  - endAr: string (in the format of "YYYY-MM-DD")
  - conditions: object (optional)
    - ageStart: int
    - ageEnd: int
    - country: array of string
    - gender: array of string
    - platform: array of string

`/ad GET`: get conditions matched ads
- the request query supports the following fields:
    - Age: int
    - Country: string
    - Gender: string
    - Platform: string

The project stucture follows the clean architecture, and the main components are:
- `ads/handler`: contains the API handlers
- `ads/service`: contains the business logic
- `ads/repository`: contains the data access logic
- `domain`: contains the data models
- `db`: contains the database connection

Using `pgx` as the database driver, and `pgpool` as the connection pool.

#### Database

The database is built with postgreSQL, and it contains six tables:
- `ads`: contains the ad information, including the ad id, ad title, start date and end date
- `ad_ages`: contains the ad age information
- `ad_genders`: contains the ad gender information
- `ad_platforms`: contains the ad platform information
- `ad_countries`: contains the ad country information
- `countries`: contains country code in ISO 3166-1 alpha-2
