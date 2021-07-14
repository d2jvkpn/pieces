import XLSX
import DataFrames as DF

# xlsx = XLSX.readdata("2021-07-14.xlsx", "Sheet1", "A1:H23")
# sheet = XLSX.readxlsx("2021-07-14.xlsx")[1]

input = ARGS[1]
output = ARGS[2]
target = ARGS[3]
println("Input xlsx: $input, Output sql: $output, Target table: $target")

df = DF.DataFrame(XLSX.readtable(input, 1)...)

# select!(df, DF.Not(r"icon")) # :icon
replace!(df.icon, missing => "")

# mapcols(col -> parse.(Int, col), df)

DF.mapcols(col -> typeof(col), df)

DF.mapcols!(col -> string.(col), df)

strRow = row -> "('$(join(row, "', '"))')"

strs = [
  "INSERT INTO $target", "(" * join(names(df), ", ") * ")",
  "VALUES", join(strRow.(eachrow(df)), ",\n  "),
]

sql = join(strs, "\n  ") * ";\n"

open(output, "w") do io
  write(io, sql)
end
