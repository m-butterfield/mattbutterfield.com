"""
Helpers for tests

"""
from unittest import TestCase

from app import app
from app import db


class BaseTest(TestCase):

    @classmethod
    def setUpClass(cls):
        if not app.config['TESTING']:
            raise EnvironmentError("Only run tests in a testing environment!")
        db.drop_all()
        db.create_all()
