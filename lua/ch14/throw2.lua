function lock2seconds(lock)
	lock:lock()
	pcall(sleep, 2000)
	lock:unlock()
end