using Printf

println("Hello, Julia!")

@printf("1/42 = %f, 1/24 = %f\n", 1/42, 1/24)

@printf(
  "Life, the Universe and Everything: %d, %d, %d, %d\n",
  Int(0b101010), Int(0x2a), Int('*'),
  (-80538738812075974)^3 + 80435758145817515^3 + 12602123297335631^3,
)
# https://news.mit.edu/2019/answer-life-universe-and-everything-sum-three-cubes-mathematics-0910
