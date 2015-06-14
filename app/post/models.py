"""
Post models

"""
from app import db


class Post(db.Model):

    __tablename__ = 'post'

    id = db.Column(db.String, primary_key=True)
    image_id = db.Column(db.String, nullable=False)
    text = db.Column(db.String, nullable=True)
    created_at = db.Column(db.DateTime, nullable=False)
