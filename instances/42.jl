using Printf

println("1/42 = ", 1/42)
println("1/24 = ", 1/24)

@printf(
  "%d, %d, %d, %d\n",
  Int(0b101010), Int(0x2a), Int('*'),
  (-80538738812075974)^3 + 80435758145817515^3 + 12602123297335631^3,
)
# https://news.mit.edu/2019/answer-life-universe-and-everything-sum-three-cubes-mathematics-0910

println("Life, the Universe and Everything: ", 42)
