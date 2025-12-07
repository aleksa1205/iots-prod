# Internet of Things and Services (IoTS)
As part of the **Internet of things and Services** univeristy course, several projects were developed to explore IoT data acquistion, processing and service-oriented architectures. These projects focused on building scalable systems for managing sensor data, implementing REST and gRPC APIs, and deploying services in containerized environments.

## Project I - IoT Microservices Application
THe project focused on building a **microservice-based system** for managing IoT sensor data. It consisted of two two services: `Gateway` and `DataManager`.

### DataManager
The **DataManager** service was responsible for all interactions with the **PostgreSQL** database. It exposed a **gRPC API** providing full CRUD operations on sensor data, enabling the **Gateway** service to efficiently access and manipulate stored information. Its primary role was to act as a reliable data access layer, ensuring consistent storage and retrieval of IoT sensor readings, while also supports aggregation operations.

#### Environment variables
| Name                    | Description                                           | Example                  | Required | Additional Notes              |
|:------------------------|:-----------------------------------------------------|:------------------------|:--------|:------------------------------|
| `DB_CONNECTION_STRING`   | Connection string to our PostgreSQL database. | ` host=${HOSTNAME} user=${POSTGRES_DB_USER} password=${POSTGRES_DB_PASSWORD} dbname=${POSTGRES_DB_NAME} port=5432 sslmode=disable` | Yes      | —                            |
| `HOST`       | Host for our service. | `localhost` | Yes       | — |
| `PORT`       | Port for our service | `6000` | Yes       | — |

### Gateway
The **Gateway** service exposed **REST API endpoints** for CRUD operations and data aggregation functions such as **min**, **max**, **avg** and **sum** over the specific time intervals. In **development mode**, it **downloads the OpenAPI schema locally** to ensure that latest version is always available and up-to-date. The service also acts as a proxy between clients and **DataManager** service, ensuring seamless communication via gRPC.

#### Environment variables
| Name                    | Description                                           | Example                  | Required | Additional Notes              |
|:------------------------|:-----------------------------------------------------|:------------------------|:--------|:------------------------------|
| `DataManager:Address`   | Address of the DataManager service for communication. | `http://localhost:8080` | Yes      | —                            |
| `Gateway:Address`       | URL of the Gateway service for fetching OpenAPI schema. | `http://localhost:5173` | No       | Used only in `Development` mode. |

### SensorGenerator
The **SensorGenerator** service simulates an IoT sensor by reading data from a `.csv` file and sending it to the Gateway service. The Gateway then forwards the data to the DataManager.

#### Environment variables
| Name                    | Description                                           | Example                  | Required | Additional Notes              |
|:------------------------|:-----------------------------------------------------|:------------------------|:--------|:------------------------------|
| `Gateway:Address`   | Base address of the Gateway service used for communication. | `http://gateway:8080` | Yes      | —                            |
| `Gateway:BatchSize`       | Number of records to include in each batch sent to Gateway. | `10` | Yes       | — |
| `Gateway:Endpoint`   | Relative endpoint on the Gateway for data submission.  | `Endpoint/Post` | Yes      | —                            |
| `Gateway:BatchTimeout`       | Time (in miliseconds) to wait before sending the next batch. | `10000` | Yes       | — |
