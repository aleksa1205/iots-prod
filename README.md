# Internet of Things and Services (IoTS)
As part of the **Internet of things and Services** univeristy course, several projects were developed to explore IoT data acquistion, processing and service-oriented architectures. These projects focused on building scalable systems for managing sensor data, implementing REST and gRPC APIs, and deploying services in containerized environments.

## Project I - IOT Microservices Application
THe project focused on building a **microservice-based system** for managing IoT sensor data. It consisted of two two services: `Gateway` and `DataManager`.

### DataManager
The **Data Manager** service was responsible for all interactions with the **PostgreSQL** database. It exposed a **gRPC API** providing full CRUD operations on sensor data, enabling the **Gateway** service to efficiently access and manipulate stored information. Its primary role was to act as a reliable data access layer, ensuring consistent storage and retrieval of IoT sensor readings, while also supports aggregation operations.

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

### Sensor Generator
The **Sensor Generator** service simulates an IoT sensor by reading data from a `.csv` file and sending it to the Gateway service. The Gateway then forwards the data to the DataManager.

#### Environment variables
| Name                    | Description                                           | Example                  | Required | Additional Notes              |
|:------------------------|:-----------------------------------------------------|:------------------------|:--------|:------------------------------|
| `Gateway:Address`   | Base address of the Gateway service used for communication. | `http://gateway:8080` | Yes      | —                            |
| `Gateway:BatchSize`       | Number of records to include in each batch sent to Gateway. | `10` | Yes       | — |
| `Gateway:Endpoint`   | Relative endpoint on the Gateway for data submission.  | `Endpoint/Post` | Yes      | —                            |
| `Gateway:BatchTimeout`       | Time (in miliseconds) to wait before sending the next batch. | `10000` | Yes       | — |

## Project II - MQTT Message Broker
In this part of the project, we integrated microservice communication using MQTT message broker. I choose **Mosquitto** as our broker. To emphasize the microservice architecture, additional services were introduced: **MQTT Client** and **Event Manager**, and the **Data Manager** was updated to support this communication. An appropriate **AsyncAPI specification** was also added to document the messaging interactions. 

### Data Manager
The **Data Manager** now sends data to the **Event Manager**. To support this, it initializes an MQTT client and publishes messages to the broker.

#### Environement Variablaes
| Name                    | Description                                           | Example                  | Required | Additional Notes              |
|:------------------------|:-----------------------------------------------------|:------------------------|:--------|:------------------------------|
| `BROKER`   | Address of the MQTT broker. | `mqtt:1883` | Yes      | —                            |
| `CLIENT_ID`       | Identifier of the Data Manager in the broker. | `data_manager_id` | Yes       | — |
| `TOPIC`       | Topic to which the Data Manager publishes messages. | `data_manager_topic` | Yes       | — |

### Event Manager
Eventmanager is used to receive data from the Data Manager that are sent on the POST method. If they are above the given used and generated threshold for KW which is appropriate for our data set. It will be sent to a different topic where Mqtt client listents on.
The **Event Manager** receives data from the **Data Manager** via `POST` requests. It evaluates the received data agains the predefined thresholds for usage (`KW`) and generation (`KW`). If the data exceeds these thresholds, it is published to a separate MQTT topic. which the **MQTT Client** subscribes to.

#### Environement Variablaes
| Name                    | Description                                           | Example                  | Required | Additional Notes              |
|:------------------------|:-----------------------------------------------------|:------------------------|:--------|:------------------------------|
| `HOST`   | Address of Event Manager service for communication. | `localhost` | Yes      | —                            |
| `PORT`       | Port used by the Event Manager service. | `7000` | Yes       | — |
| `BROKER`       | Address of the MQTT broker.  | `mqtt:1883` | Yes       | — |
| `CLIENT_ID`   | Identifier of the Event Manager in the broker. | `event_manager_id` | Yes      | —                            |
| `RECV_TOPIC`       | MQTT topic from which the Event Manager receives data. | `event_manager_recv_topic` | Yes       | — |
| `SEND_TOPIC`       | MQTT topic to which Event Manager publishes data. | `event_manager_send_topic` | Yes       | — |
| `USED_THRESHOLD`       | Threshold value for used KW. | `1.5` | Yes       | — |
| `GEN_THRESHOLD`       | Threshold value for generated KW. | `1.2` | Yes       | — |

### MQTT Client
The **MQTT Client** is responsible for subscribing to a specific topic and receiving data published by the **Event Manager**.

#### Environement Variablaes
| Name                    | Description                                           | Example                  | Required | Additional Notes              |
|:------------------------|:-----------------------------------------------------|:------------------------|:--------|:------------------------------|
| `HOST`   | Address of MQTT Client service for communication. | `localhost` | Yes      | —                            |
| `PORT`       | Port used by the MQTT Client service. | `7000` | Yes       | — |
| `BROKER`       | Address of the MQTT broker.  | `mqtt:1883` | Yes       | — |
| `CLIENT_ID`   | Identifier of the MQTT Client in the broker. | `mqtt_client_id` | Yes      | —                            |
| `TOPIC`       | MQTT topic from which the MQTT Client receives data. | `mqtt_client_topic` | Yes       | — |
