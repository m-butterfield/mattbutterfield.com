"""
Post models

"""
import os

from urlparse import urlunsplit

from app import app
from app import db


class Post(db.Model):

    __tablename__ = 'post'

    id = db.Column(db.String, primary_key=True)
    image_uri = db.Column(db.String, nullable=False)
    text = db.Column(db.String, nullable=True)
    created_at = db.Column(db.DateTime, nullable=False)
    previous_post_id = db.Column(
        db.String, db.ForeignKey('post.id'), nullable=True)
    next_post_id = db.Column(
        db.String, db.ForeignKey('post.id'), nullable=True)

    @property
    def image_url(self):
        """
        Build a url for this post's image on S3 based on its image_uri

        """
        return urlunsplit((
            app.config['S3_URL_SCHEME'],
            app.config['S3_IMAGE_BUCKET'],
            os.path.join(app.config['S3_IMAGE_FOLDER'], self.image_uri),
            None,
            None))
