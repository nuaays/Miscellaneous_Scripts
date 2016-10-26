from flask_wtf import FlaskForm
from wtforms import TextField, BooleanField
from wtforms.validators import Required

class LoginForm(FlaskForm):
	#<input id="openid" name="openid" size="80" type="text" value="">
    openid = TextField('openid', validators = [Required()])
    #<input id="remember_me" name="remember_me" type="checkbox" value="y"> Remember Me
    remember_me = BooleanField('remember_me', default = False)
    
