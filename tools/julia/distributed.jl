using Distributed

using TOML

#= example
addprocs(
    [("root@172.17.0.2:22", 2)];
    dir="/root/wk_julia",
    exename="/opt/julia-1.5.3/bin/julia",
    tunnel=true,
)
=#

function addmachine(c)
    println(">>> connect to machine $(c["name"]) => $(c["addr"])")

    addprocs(
        [(c["addr"], c["n"])];
        dir=c["dir"], exename=c["exename"], tunnel=true,
    )
end

config = TOML.parsefile("config.toml")
machines = config["machines"]

println("### found $(length(machines)) machines")

for i in 1:length(machines)
    addmachine(machines[i])
    machines[i]["id"] = i+1
end

println(">>> nworkers: $(nworkers())")
