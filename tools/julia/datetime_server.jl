import Dates, TimeZones

import Genie


function currentTime()
  now = Dates.now()

  TimeZones.ZonedDateTime(
    Dates.year(now), Dates.month(now), Dates.day(now),
    Dates.hour(now), Dates.minute(now), Dates.second(now),
    Dates.millisecond(now), TimeZones.localzone(),
  )
end


port = length(ARGS) > 0 ? parse(Int64, ARGS[1]) : 8000
Genie.config.run_as_server = true

Genie.route("/") do
  Dates.format(currentTime(), "yyyy-mm-ddTHH:MM:SS.sssz")
end

println("starting service $(port)")
Genie.startup(port=port)
