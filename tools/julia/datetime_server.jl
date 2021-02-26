using Genie, Genie.Router

using Dates, TimeZones


function currentTime()
  a = now()
  ZonedDateTime(year(a), month(a), day(a), hour(a), minute(a), second(a), millisecond(a), localzone())
end


Genie.config.run_as_server = true
port = length(ARGS) > 0 ? parse(Int64, ARGS[1]) : 8000


route("/") do
  currentTime()
end

println("starting service $(port)")
Genie.startup(port=port)
