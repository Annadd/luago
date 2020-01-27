co = coroutine.create(function()
	for k, v in pairs({"a", "b", "c"}) do
		coroutine.yield(k, v)
	end
	return "d", 4
end)

print(coroutine.resume(co)) --> true 1 a
print(coroutine.resume(co)) --> true 2 b
print(coroutine.resume(co)) --> true 3 c
print(coroutine.resume(co)) --> true 4 d
print(coroutine.resume(co)) --> false cannot resume dead coroutine