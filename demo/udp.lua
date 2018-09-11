local buffer = require("buffer")
local time = require("time")
local udp = require("net.udp")

local server,err = udp:listen("127.0.0.1",1024)

if (nil ~= err) then
    print('listen error: ',err)
    return
end

local buf = buffer:new(256)
local tkn = buffer:form('"st":"1"')

repeat
    local n,e = server:read(buf)
    if (buf:slice(1,n):index(tkn) > 0) then
        print("alarm: ",time:format("2006-01-02 15:04:05"))
    end
until (nil ~= e)

server:close()