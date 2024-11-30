import polars as pl

df = pl.read_csv("./yellow_tripdata_2015-03.csv")

x = df.sum()

print(x)
