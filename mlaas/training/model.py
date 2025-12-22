import joblib
import pandas as pd
import numpy as np
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_absolute_error
from sklearn.model_selection import train_test_split 

# Dataset adapted
df = pd.read_csv("./data.csv")
df = df.drop(columns=["temperature", "humidity", "apparentTemperature", "pressure"])

def create_dataset(df, n_past=20, n_future=150):
    x, y = [], []
    target_values = df["use [kW]"].values

    for i in range(len(target_values) - n_past - n_future):    
        past_window = target_values[i : i + n_past]
        future_avg = target_values[i + n_past : i + n_past + n_future].mean()
        x.append(past_window)
        y.append(future_avg)
    
    return np.array(x), np.array(y)
    
x, y = create_dataset(df)

x_train, x_test, y_train, y_test = train_test_split(
    x, y, test_size=0.2, shuffle=False
)

model = LinearRegression()
model.fit(x_train, y_train)

y_pred = model.predict(x_test)

mae = mean_absolute_error(y_test, y_pred)
print(f"Mean Absolute Error: {mae}")

joblib.dump(model, "ml_model.pkl")