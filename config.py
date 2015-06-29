import os

DEBUG = True

BASE_DIR = os.path.abspath(os.path.dirname(__file__))

TESTING = os.environ.get('TESTING')

if TESTING:
    SQLALCHEMY_DATABASE_URI = 'sqlite:///' + os.path.join(BASE_DIR, 'test.db')
else:
    SQLALCHEMY_DATABASE_URI = 'sqlite:///' + os.path.join(BASE_DIR, 'app.db')

SERVER_NAME = 'localhost:8080'

S3_URL_SCHEME = 'http'
S3_IMAGE_BUCKET = 'images.mattbutterfield.com'
S3_IMAGE_FOLDER = 'post_images'