import joblib
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_absolute_error
from sklearn.model_selection import train_test_split 
from sklearn.preprocessing import StandardScaler
from sklearn.pipeline import Pipeline

# Load and prepare dataset
df = pd.read_csv("./data.csv")
df = df.drop(columns=["time", "temperature", "humidity", "apparentTemperature", "pressure"])

def create_dataset(df, n_past=20, n_future=50):
    """
    Generate features and targets for time series forecasting.

    Each feature consists of `n_past` consecutive past values of
    "use [kW]" and "gen [kW]". These are flattened into a 1D array.
        
    Arguments:
        df: DataFrame with "use [kW]" and "gen [kW]" columns.
        n_past: Number of past time steps to use as input (default 20).
        n_future: Number of future time steps to average as the target (default 50).
        
    Returns:
        tuple:
            x: 2D array of past values, shape (num_samples, n_past*2)
            y: 2D array of future averages, shape (num_samples, 2)
    """

    x, y = [], []
    use_values = df["use [kW]"].values
    gen_values = df["gen [kW]"].values

    for i in range(len(use_values) - n_past - n_future):    
        past_use = use_values[i : i + n_past]
        past_gen = gen_values[i : i + n_past]
        
        past_window = np.column_stack((past_use, past_gen))
        x.append(past_window.flatten())
        future_use_avg = use_values[i + n_past : i + n_past + n_future].mean()
        future_gen_avg = gen_values[i + n_past : i + n_past + n_future].mean()

        y.append([future_use_avg, future_gen_avg])
        
    return np.array(x), np.array(y)
    
x, y = create_dataset(df)

# Train/Test split
x_train, x_test, y_train, y_test = train_test_split(
    x, y, test_size=0.3, shuffle=False
)

# Train model
model = Pipeline([
    ("scaler", StandardScaler()),
    ("regressor", LinearRegression()),
])

model.fit(x_train, y_train)

# Evaluate model
y_pred = model.predict(x_test)

mae_use = mean_absolute_error(y_test[:, 0], y_pred[:, 0])
mae_gen = mean_absolute_error(y_test[:, 1], y_pred[:, 1])

print(f"MAE use [kW]: {mae_use:.6f}")
print(f"MAE gen [kW]: {mae_gen:.6f}")

results = pd.DataFrame({
    "Actual use [kW]": y_test[:, 0],
    "Predicted use [kW]": y_pred[:, 0],
    "Actual gen [kW]": y_test[:, 1],
    "Predicted gen [kW]": y_pred[:, 1],
})

# ---- GEN ----
plt.figure(figsize=(10, 4))
plt.plot(results["Actual gen [kW]"], label="Actual gen")
plt.plot(results["Predicted gen [kW]"], label="Predicted gen")
plt.legend()
plt.title("Gen [kW] prediction")
plt.show()

# ---- USE ----
plt.figure(figsize=(10, 4))
plt.plot(results["Actual use [kW]"], label="Actual use")
plt.plot(results["Predicted use [kW]"], label="Predicted use")
plt.legend()
plt.title("Use [kW] prediction")
plt.show()

# Save model
joblib.dump(model, "ml_model.pkl")