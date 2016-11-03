#coding:utf-8

from flask import render_template, flash, redirect
from app import app
from forms import LoginForm

@app.route('/login', methods = ["GET", "POST"])
def login():
    form = LoginForm()
    return render_template('login.html',
                           title = 'Sign In',
                           form = form)
