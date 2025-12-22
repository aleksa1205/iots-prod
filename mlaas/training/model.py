import joblib
import pandas as pd
import numpy as np
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_absolute_error
from sklearn.model_selection import train_test_split 

# Load and prepare dataset
df = pd.read_csv("./data.csv")
df = df.drop(columns=["temperature", "humidity", "apparentTemperature", "pressure"])

def create_dataset(df, n_past=20, n_future=200):
    """
    Generate features and targets for time series forecasting.

    Each feature consists of `n_past` consecutive past values of "use [kW]".
    Each target is the average of the next `n_future` values.
        
    Arguments:
        df: DataFrame with "use [kW]" column.
        n_past: Number of past time steps to use as input.
        n_future: Number of future time steps to predict average as the target.
        
    Returns:
        tuple:
            x: 2D array of past values.
            y: 1D array of average future values.
    """
    x, y = [], []
    target_values = df["use [kW]"].values

    for i in range(len(target_values) - n_past - n_future):    
        past_window = target_values[i : i + n_past]
        future_avg = target_values[i + n_past : i + n_past + n_future].mean()
        x.append(past_window)
        y.append(future_avg)
        
    return np.array(x), np.array(y)
    
x, y = create_dataset(df)

# Train/Test split
x_train, x_test, y_train, y_test = train_test_split(
    x, y, test_size=0.3, shuffle=False
)

# Train model
model = LinearRegression()
model.fit(x_train, y_train)

# Evaluate model
y_pred = model.predict(x_test)
mae = mean_absolute_error(y_test, y_pred)
print(f"Mean Absolute Error: {mae:.7f}")

# Save model
joblib.dump(model, "ml_model.pkl")