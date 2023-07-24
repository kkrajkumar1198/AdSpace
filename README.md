# Ad Space

The Ad Space project is a simple web application built using Go and MySQL to manage ad space and facilitate bidding by different bidders.

## Overview

The project consists of two main services: the **Supply Side Service** and the **Demand Side Service**. The Supply Side Service is responsible for listing the ad spaces, while the Demand Side Service lists all the bidders interested in bidding for the ad spaces and helps to create a bids.

## Prerequisites

Before running the Ad Space project, ensure you have the following installed on your system:

- Go (1.16 or higher)
- Docker
- Docker Compose

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/kkrajkumar1198/AdSpace.git
cd AdSpace
```

2. Build and run the Docker containers:

```bash
docker-compose up --build
```

3. The application will be accessible at the following URLs:

- Supply Side Service: [http://localhost:8080](http://localhost:8080)
- Demand Side Service: [http://localhost:8081](http://localhost:8081)

## API Endpoints

### Supply Side Service

- **GET /ad_spaces**: Get a list of all available ad spaces.

### Demand Side Service

- **GET /list_bidders**: Get a list of all bidders interested in bidding for ad spaces.

- **POST /createnewbids**: Create new bids on the adSpace and store it in the database.

```json
{
  "bidder_id": 1,
  "ad_space_id": 1,
  "amount": 100.50,
  "bid_time": "2023-07-25 15:00:00"
}
```

- **POST /get_ad_space_bids**: Get bids details on the available adSpace.

```json
{
    "ad_space_id":1
}
```

## Database

The application uses a MySQL database to store ad spaces, bidders, and bids. The database configuration is specified in the `docker-compose.yml` file.
