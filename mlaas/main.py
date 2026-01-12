from fastapi import FastAPI
import joblib
import numpy as np
from pydantic import BaseModel

model = joblib.load("./training/ml_model.pkl")
scaler = joblib.load("./training/scaler.pkl")
app = FastAPI()

class PastStep(BaseModel):
    use_kw: float
    gen_kw: float

class PredictRequest(BaseModel):
    past_values: list[PastStep]
    
class PredictResponse(BaseModel):
    use_kw: float
    gen_kw: float
    net_kw: float = None

N_PAST = 20
    
def normalize_steps(steps: list[PastStep]) -> list[PastStep]:
    if len(steps) >= N_PAST:
        return steps[-N_PAST:]

    missing = N_PAST - len(steps)
    return [steps[0]] * missing + steps

@app.get("/")
def read_root():
    return {"message": "Welcome to the MLaaS API"}

@app.post("/predict/")
def predict(request: PredictRequest) -> PredictResponse:
    steps = normalize_steps(request.past_values)

    use_values = [s.use_kw for s in steps]
    gen_values = [s.gen_kw for s in steps]

    past_window = np.column_stack((use_values, gen_values))
    model_input = past_window.flatten().reshape(1, -1)
    model_input_scaled = scaler.transform(model_input)

    predicted_next = model.predict(model_input_scaled)[0]
    
    return PredictResponse(
        use_kw=float(predicted_next[0]),
        gen_kw=float(predicted_next[1]),
        net_kw=float(predicted_next[1] - predicted_next[0])
    )