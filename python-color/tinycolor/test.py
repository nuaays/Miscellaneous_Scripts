#coding:utf-8
import sys  

reload(sys)  
sys.setdefaultencoding('utf8')

from tinycolor import color as c

print("Hey, this is {}!".format(c.green("cool")))
# Will print a green "cool"

print("It also works with {} colors".format(c.green_on_blue("background")))
# Will print "background" in green text on blue background

print("And with {} colors".format(c.bright_green("bright")))
# Will print a bright green "bright"

print "this is {}".format(c.red("red"))
print "this is {}".format(c.green("green"))
print "this is {}".format(c.yellow("yellow"))
print "this is {}".format(c.magenta("magenta"))
print "this is {}".format(c.cyan("cyan"))
print "this is {}".format(c.white_on_red("红底白字"))
print "this is {}".format(c.white_on_green("绿底白字"))
print "this is {}".format(c.white_on_magenta("紫底白字"))
