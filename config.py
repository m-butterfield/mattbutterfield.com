import os

DEBUG = True

BASE_DIR = os.path.abspath(os.path.dirname(__file__))

SQLALCHEMY_DATABASE_URI = 'sqlite:///' + os.path.join(BASE_DIR, 'app.db')

SERVER_NAME = 'localhost:8080'

S3_IMAGE_BUCKET = 'images.mattbutterfield.com'
S3_IMAGE_FOLDER = 'post_images'
