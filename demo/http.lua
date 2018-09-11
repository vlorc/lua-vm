local http = require("net.http")

local str,err = http:getString("http://www.baidu.com")

if (nil ~= err) then
    print('getString error: ',err:error())
else
    print(str)
end