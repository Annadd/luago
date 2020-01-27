co = coroutine.create(function() print("hello") end)
print(type(co)) --> thread

main = coroutine.running()
print(type(main))  --> thread
print(coroutine.status(main)) --> running