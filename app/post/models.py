"""
Post models

"""
from app import db


class Post(db.Model):

    __tablename__ = 'post'

    id = db.Column(db.Integer, primary_key=True)
    text = db.Column(db.String, nullable=True)
