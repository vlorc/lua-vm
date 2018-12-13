local M = {}

function M:hello(s)
    print("hello:",s)
end

function M:add(a,b)
    return a + b
end

return M