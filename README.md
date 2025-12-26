## Project III - MLaaS and NATS

### MLaaS
**MLaaS** is a microservice used to train model on incoming data and predict future values based on historical measurments.

#### Tech Stack
⚙️ **Language:** Python  

🛠 **Web Framework:** FastAPI  

📊 **ML / Data:** scikit-learn, pandas, numpy  

💾 **Model Serialization:** joblib  

#### Endpoints
- **GET /** – Root endpoint. Can be used to check if the service is running.  
- **POST /predict/** – Accepts recent historical data and returns predicted future values from the trained ML model.

#### Development
1. Make sure all dependencies are installed 
```bash
pip install -r requirements.txt
```

2. Run the service in development mode:
```
uvicorn main:app --reload
```

3. Open Swagger UI to explore and test endpoints easily:
```
http://{address}:{port}/docs
```

# Other
- [Kaggle - Smart Home Dataset With Weather Information](https://www.kaggle.com/datasets/taranvee/smart-home-dataset-with-weather-information)
