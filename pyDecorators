import os, sys



def fibon(n):
	a = b = 1
	for i in range(n):
		yield a
		a , b = b, a+b

def generator_function():
	for i in range(10):
		yield i

for item in generator_function():
	print item


for x in fibon(1000):
	print x


### Map
items = [1,2,3,4,5]
squared = []
for i in items:
	squared.append(i**2)

squared_map = list(map(lambda x:x**2, items))
print squared_map

def multiply(x):
	return (x*x)
def add(x):
	return (x+x)
funcs = [multiply, add]
for i in range(5):
	value = list(map(lambda x:x(i), funcs))
	print value

###Filter
number_list = range(-5, 5)
less_than_zero = list(filter(lambda x: x < 0, number_list))
print less_than_zero

###Decorators
def hi(name="yasoob"):
	print "hi " + name

	def greet():
		return "now you are in the greet() function"
	def welcome():
		return "now you are in the welcome() function"

	if name == "yasoob":
		return greet
	else:
		return welcome

a = hi()
print(a)
print(a())


###The first Decorator
print "========================================================="
from functools import wraps

def a_new_decorator(a_func):
	@wraps(a_func)
	def wrapTheFunction():
		print "I am doing some boring work before executing a_func"
		a_func()
		print "I am doing some boring work after executing a_func"
	return wrapTheFunction

#@a_new_decorator
def test():
	print "hello world"

test()
test = a_new_decorator(test)
test()
print test.__name__




### Decorator used for logging
#with parameter 
# from functools import wraps
# def logit(logfile="out.log"):
# 	def logging_decorator(func):
# 		@wraps(func)
# 		def wrap_function(*args, **kwargs):
# 			log_string = func.__name__ + " was called"
# 			print log_string
# 			with open(logfile, 'a') as f:
# 				f.write(log_string)
# 			#return func(*args, **kwargs)
# 		return wrap_function
# 	return logging_decorator


# @logit
# def func1():
# 	pass

# func1()

# @logit(logfile="func2.log")
# def func2():
# 	pass
# func2()


### decorator class
class logit(object):
	def __init__(self, logfile='out.log'):
		self.logfile  = logfile
	def __call__(self, func):
		log_string = func.__name__ + "was called"
		print log_string
		with open(self.logfile, 'a') as f:
			f.write(log_string + "\n")
		#
		self.notify()
	def notify(self):
		pass

@logit
def myfunc1():
	pass
