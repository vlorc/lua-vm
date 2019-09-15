local buffer = require("buffer")
local tcp = require("net.tcp")

local con,err = tcp:connect("www.google.com:80")

if (nil ~= err) then
    print('connect error: ',err:error())
    return
end

local buf = buffer:new(4096)
con:write(buffer:form("GET / HTTP/1.1\n\n"))

repeat
    local n,e = con:read(buf)
    if (n > 0) then
        print(buf:toString("raw"))
    end
until (nil ~= e)

con:close()