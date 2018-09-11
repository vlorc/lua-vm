local time = require("time")

local tick = time:ticker(1000)
local count = 0

time:after(5000,function()
    print("after 5s")
end)

repeat
    tick:wait()
    count = count + 1
    print(count)
until(count > 10)

tick:close()