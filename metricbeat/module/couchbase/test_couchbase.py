import os
import sys
import unittest
from parameterized import parameterized

sys.path.append(os.path.join(os.path.dirname(__file__), '../../tests/system'))
import metricbeat


@unittest.skip("See https://github.com/elastic/beats/issues/14660")
class Test(metricbeat.BaseTest):

    # Commented out as part of skipping test. See https://github.com/elastic/beats/issues/14660.
    # Otherwise, the tests are skipped but Docker Compose still tries to bring up
    # the Couchbase service container and fails.
    # COMPOSE_SERVICES = ['couchbase']
    FIELDS = ['couchbase']

    @parameterized.expand([
        ("bucket"),
        ("cluster"),
        ("node"),
    ])
    @unittest.skipUnless(metricbeat.INTEGRATION_TESTS, "integration test")
    def test_couchbase(self, metricset):
        """
        couchbase metricsets tests
        """
        self.check_metricset("couchbase", metricset, self.get_hosts(), self.FIELDS)

    def get_hosts(self):
        return ["http://Administrator:password@" + self.compose_host()]
