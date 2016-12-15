from .base import BaseTestCase


class HealthTestCase(BaseTestCase):
    def test_health(self):
        """
        Service should report health success
        """
        self.assert_request('get', '/_health')
