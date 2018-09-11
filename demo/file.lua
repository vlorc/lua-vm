local buffer = require("buffer")
local fs = require("fs")
local reader = require("io.reader")
local fd,err = fs:open("main/main.go")

if (nil ~= err) then
    print('open error: ',err:error())
    return
end

local rd = reader:new(fd)

repeat
    local s,e = rd:readLine()
    if (nil == e) then
        print(s)
    end
until (nil ~= e)

fd:close()