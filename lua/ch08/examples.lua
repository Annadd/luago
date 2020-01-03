function f(a, b, c)
	print(a, b, c)
end

f()
f(1, 2)
f(1, 2, 3, 4, 5)

function f(a, ...)
	local b, c = ...
	local t = {a, ...}
	print(a, b, c, #t, ...)
end

f()
f(1, 2)
f(1, 2, 3, 4)

function f()
	return 1, 2, 3
end

f()
a, b = f()
assert(a == 1 and b == 2)
a, b, c = f()
assert(a == 1 and b == 2 and c == 3)
a, b, c, d = f()
assert(a == 1 and b == 2 and c == 3 and d == nil)

function f()
	return 1, 2, 3
end

a, b = (f())
assert(a == 1 and b == nil)
a, b, c = 5, f(), 5 
assert(a == 5 and b == 1 and c == 5)

function f() return 3, 2, 1 end 
function g() return 4, f() end 
function h(a, b, c, d) print(a, b, c, d) end

h(4, f())
h(g())
print(table.unpack({4, f()}))