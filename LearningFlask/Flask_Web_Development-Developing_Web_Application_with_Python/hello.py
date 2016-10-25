#!/usr/bin/env python
# -*- coding=UTF-8 -*-
#
import os, sys
from flask import Flask
from flask import request, url_for, render_template
from flask import make_response
from flask import redirect
from flask_bootstrap import Bootstrap
from flask_wtf import FlaskForm, Form
from wtforms import StringField, SubmitField
from wtforms.validators import Required

from flask_sqlalchemy import SQLAlchemy
from flask_script import Manager


app = Flask(__name__)
app.config.from_object('config')
db = SQLAlchemy(app)
bootstrap = Bootstrap(app)
manager = Manager(app)




class Role(db.Model):
    __tablename__ = 'roles'
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(64), unique=True)

class NameForm(Form):
    name = StringField('What is your name', validators=[Required()])
    submit = SubmitField('Submit')

import random

from flask import g

@app.before_request
def set_on_g_object():
  x = random.randint(0, 9)
  app.logger.debug('before request: g.x is {x}'.format(x=x))
  g.x = x

@app.after_request
def get_on_g_object(response):
  app.logger.debug('after request: g.x is {g.x}'.format(g=g))
  return response



@app.route("/")
def index():
    user_agent = request.headers.get('User-Agent')
    # response = make_response("<h1>This document carries a cookie!</h1><br>"+user_agent)
    # response.set_cookie('answer','42')
    name = None
    form = FlaskForm()
    if form.validate_on_submit():
        name = form.name.data
        form.name.data = ''
    return render_template('index.html', form=form, name=name)

@app.route('/object/')
def return_object():
    headers = {'Content-Type': 'text/plain'}
    return make_response('Hello, world!', 200, headers)

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

