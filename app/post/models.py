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

    @property
    def next_post(self):
        """
        Return the next most recently created post or None

        """
        return (db.session.query(Post)
                .filter(Post.created_at > self.created_at)
                .order_by(Post.created_at).first())

    @property
    def previous_post(self):
        """
        Return the next oldest created post or None

        """
        return (db.session.query(Post)
                .filter(Post.created_at < self.created_at)
                .order_by(Post.created_at.desc()).first())
