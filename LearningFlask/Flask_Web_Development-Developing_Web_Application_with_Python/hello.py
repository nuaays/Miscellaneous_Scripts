from flask import Flask
from flask import request
from flask import make_response
from flask import redirect

from flask_script import Manager
app = Flask(__name__)
manager = Manager(app)

@app.route("/")
def hello():
    user_agent = request.headers.get('User-Agent')
    response = make_response("<h1>This document carries a cookie!</h1><br>"+user_agent)
    response.set_cookie('answer','42')
    return response

@app.route("/bad")
def index():
    return "<h1>Bad Request</h1>", 404

#@app.route("/redirect")
#def redirect():
#    return redirect('http://www.baidu.com')

@app.route("/user/<name>")
def user(name):
    return "<h1>Hello, %s!</h1>" % name



if __name__ == "__main__":
    #app.run()
    manager.run()

