from fastapi import FastAPI
import joblib
import numpy as np

app = FastAPI()

model = joblib.load("")

@app.get("/")
def read_root():
    return {"message": "Welcome to the MLaaS API"}

@app.post("/predict/")
def predict(data: dict):
    features = np.array(data["features"]).reshape(1, -1)
    prediction = model.predict(features)
    class_name = iris.target_names[prediction][0]
    return {"class": class_name}