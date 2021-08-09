#! python3
# -*- coding: utf-8 -*-

import datetime

t1 = datetime.datetime.utcnow()
t2 = datetime.datetime.now()

print(datetime.datetime.timestamp(t1), datetime.datetime.timestamp(t2))

print(t1, t2)

print(t1.isoformat(), t2.isoformat())

d = datetime.datetime.utcnow()
d_with_timezone = d.replace(tzinfo=pytz.UTC)
d_with_timezone.isoformat()
