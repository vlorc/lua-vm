local buffer = require("buffer")
local time = require("time")
local udp = require("net.udp")

local server,err = udp:listen("0.0.0.0:0")
if (nil ~= err) then
    print('listen error: ',err)
    return
end

print('listen addr: ',server:localAddr())

local buf = buffer:new(256)

repeat
    local n,e = server:read(buf)
    if (n > 0) then
        print("data: ",time:format("2006-01-02 15:04:05"),"> ",buf:slice(1,n):tostring("raw"))
    end
until (nil ~= e)

server:close()