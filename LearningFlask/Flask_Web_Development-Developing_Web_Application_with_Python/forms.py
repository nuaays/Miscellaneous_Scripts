
from __future__ import print_function

from wtforms import Form, BooleanField, DateTimeField
from wtforms import TextAreaField, TextField
from wtforms.validators import Length, required

class AppointmentForm(Form):
  title = TextField('Title', [Length(max=255)])
  start = DateTimeField('Start', [required()])
  end = DateTimeField('End')
  allday = BooleanField('All Day')
  location = TextField('Location', [Length(max=255)])
  description = TextAreaField('Description')


# Demonstration of a WTForms form by itself.
form = AppointmentForm()
print('Here is how a form field displays:')
print(form.title.label)
print(form.title)