"""
Store service for functional testing
"""

import subprocess
import atexit
import os
import tempfile
import shutil
import time
import requests


def start_process(command, cwd, logfp):
    """
    Start a process
    """
    environ = os.environ.copy()
    environ["LC_ALL"] = "C"

    return subprocess.Popen(command, cwd=cwd, stderr=subprocess.STDOUT,
                            stdout=logfp, env=environ)


class StoreService(object):
    """
    Store service startup/shutdown.
    """
    atexit_handler = False

    process = None
    tmpdir = None
    http_address = "127.0.0.1:8001"
    logfp = None

    def start(self):
        """
        Start store-service
        """
        # Regiester exit handler
        if not self.atexit_handler:
            atexit.register(self.stop)
            self.atexit_handler = True

        # Make a temporary directory to run store service
        if self.tmpdir is None:
            self.tmpdir = tempfile.mkdtemp(prefix='store-service')
            shutil.copy("../config.base.yaml", os.path.join(self.tmpdir, "config.base.yaml"))
            shutil.copy("../bin/store-service", self.tmpdir)

        if self.logfp is None:
            self.logfp = open("store-service.log", "w")

        if self.process is None:
            self.process = start_process(
                ['./store-service'], cwd=self.tmpdir, logfp=self.logfp)

            for _ in xrange(120):
                try:
                    res = requests.get("http://" + self.http_address + "/_version", timeout=0.5)
                    if res.status_code == 200:
                        break
                except requests.ConnectionError:
                    pass

                time.sleep(0.3)

    def stop(self):
        """
        Stop the process
        """
        if self.process is not None:
            self.process.terminate()
            self.process.wait()
            self.process = None

        if self.tmpdir is not None:
            shutil.rmtree(self.tmpdir)
            self.tmpdir = None

    def restart(self):
        """
        Restart the process
        """
        self.stop()
        self.start()


# Singleton instance
service = StoreService()
