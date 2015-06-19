"""
Views for the application

"""
import json

from flask import render_template

from app import app
from app.post import api as post_api
from app.post.lib import get_post_or_404
from app.post.resources import serialize_post


@app.route('/')
@app.route('/post/<post_id>')
def index(post_id=None):
    if post_id:
        post = get_post_or_404(post_id)
    else:
        post = post_api.get_most_recent_post()
    return render_template('index.html', post=json.dumps(serialize_post(post)))


@app.errorhandler(404)
def page_not_found(error):
    return 'Not found :(', 404
