local dns = require("net.dns")
local ips = dns:lookup("www.baidu.com")

for key,value in ips() do
   print(key, value)
end