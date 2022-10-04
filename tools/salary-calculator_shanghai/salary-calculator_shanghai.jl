using Printf

# http://salarycalculator.sinaapp.com/

base = 6520.0 # 2022
hf = 7.0

# println(Base.source_path())
# @printf("%s\n", ARGS)

if length(ARGS) > 0
    base = parse(Float64, ARGS[1])
end

if length(ARGS) > 1
    hf = parse(Float64, ARGS[2])
end

perc_names = ("养老保险金", "医疗保险金", "失业保险金", "工伤保险金", "生育保险金", "基本住房公积金")
perc_p = (8.0,  2.0, 0.5, 0.0,  0.0, hf)
perc_c = (16.0, 9.5, 0.5, 0.16, 1.0, hf)

ps = base * sum(perc_p)/100.0
cs = base * sum(perc_c)/100.0

@printf(
  "个人税后收入: %.2f, 雇佣基本成本: %.2f\n",
  round(base - ps, digits=2),
  round(base + cs, digits=2),
)
