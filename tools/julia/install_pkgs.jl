using Pkg


if length(ARGS) > 0
    file = ARGS[1]
else
    file = "pkgs_Julia.txt"
end

# lines = split(read(text, String), "\n")
lines = [strip(line) for line in readlines(file)]
lines = filter(x -> !(occursin("#", x) || x == ""), lines)
# lines = [line for line in lines if !(occursin("#", line) || line == "")]
pkgs = collect(Iterators.flatten([split(line) for line in lines]))
if length(pkgs) == 0
    println("no package to install")
    exit(0)
end


println("### Installing $(length(pkgs)) packages...")

for p in pkgs
#=
    installedList = keys(Pkg.installed())
    if p in installedList continue end
=#
    try
        println(">>> installing $p")
        Pkg.add(p) # Pkg.build(p)
	catch
        filter!(e -> e ≠ p, pkgs)
        println("!!! failed to install $p")
    end
end

println("### Precompiling packages...")

for p in pkgs
    println(">>> precompiling $p")
    eval(Meta.parse(string("using $p")))
end
