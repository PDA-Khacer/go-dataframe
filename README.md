# go-dataframe
Project clone pandas python

----------------------
Using idea for Json handler from: https://github.com/marhaupe/json2struct
check it in vendor_lib
----------------------
Current method support:

`IO`:
- `SeriesJsonNormalize` - convert SeriesData -> to `Dataframe<string>` (JSON flat)

`Dataframe`:
- `Agg` - support Max, Min, Sum, Avg (Noted: Dataframe must number)
- `Apply` - apply some function to all cell
- `Drop` - drop col
- `AsType` - Cast to other type (if can)
- `DataframeValuesCount` - like `Series.values_count`, counter all values
