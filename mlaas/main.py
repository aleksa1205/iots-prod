from fastapi import FastAPI
import joblib
import numpy as np
from pydantic import BaseModel

model = joblib.load("./training/ml_model.pkl")
app = FastAPI()

class PredictRequest(BaseModel):
    past_values: list[float]
    
class PredictResponse(BaseModel):
    prediction: float

@app.get("/")
def read_root():
    return {"message": "Welcome to the MLaaS API"}

@app.post("/predict/")
def predict(request: PredictRequest) -> PredictResponse:
    past_values = request.past_values
    predicted_avg = model.predict([past_values])[0]
    return {"prediction": predicted_avg}