local regexp = require('regexp')
local re = regexp:mustCompile("\\d+")

local result = re:findAllString("N086-156-1231",-1)
print(#result,result[1],result[2],result[3])
