function tbp(n)
   if n >= 365
     return 1.0
   end

   return 1 - reduce(*, [i/365 for i in (365-n+1):365])
   # return 1 - prod([i/365 for i in (365-n+1):365])
end
