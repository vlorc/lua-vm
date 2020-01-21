local buffer = require("buffer")

local buf = buffer:form(1,2,3,4,5,6,7,8,9)

print(buf:toHash())
print(buf:toNumber(1,1))
print(buf:toString("BCD"))
print(buf:index(5))