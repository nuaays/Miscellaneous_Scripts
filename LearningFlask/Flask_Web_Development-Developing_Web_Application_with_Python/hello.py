#!/usr/bin/env python
#coding:utf-8
#
from flask import Flask
from flask import request, url_for, render_template
from flask import make_response
from flask import redirect
from flask_bootstrap import Bootstrap

from flask_script import Manager
app = Flask(__name__)
app.debug = True
manager = Manager(app)

@app.route("/")
def hello():
    user_agent = request.headers.get('User-Agent')
    response = make_response("<h1>This document carries a cookie!</h1><br>"+user_agent)
    response.set_cookie('answer','42')
    return response

@app.errorhandler(404)
def page_not_found(e):
    return render_template('404.html'), 404

@app.errorhandler(500)
def internal_server_error(e):
    return render_template('500.html'), 500

@app.route("/bad")
def redirect_to_error():
    return render_template('500.html'), 500

#@app.route("/redirect")
#def redirect():
#    return redirect('http://www.baidu.com')

@app.route("/user/<name>")
def user(name):
    #return "<h1>Hello, %s!</h1>" % name
    return render_template('hello.html', name=name) 

@app.route("/about", methods=["GET", "POST"])
def about():
    return "The about page"

if __name__ == "__main__":
    #app.run()
    manager.run()

