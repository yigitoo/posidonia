#!/usr/bin/env python
#author: Yigit GÜMÜŞ
import threading
import sys
import os
import time

class AppExec(threading.Thread):
    def __init__(self, title, command):
        threading.Thread.__init__(self)
        self.cmd = command
        self.title = title

    def run(self):
        os.system(self.cmd)

lstTitle = ["Setup server", "Setup Website", "Run Server", "Run Application"]
lstCmd   = [
    "cd server && go mod tidy && cd ..",
    "cd posidonia-website && sudo bundle install && cd ..",
    "cd server && make > /dev/null 2>&1",
    "cd posidonia-website && make > /dev/null 2>&1"
]


if __name__ == "__main__":
    setup_server = None
    setup_website = None

    if sys.argv[-1] == "--wsetup":
        setup_server = AppExec(lstTitle[0], lstCmd[0])
        setup_website = AppExec(lstTitle[1], lstCmd[1])
        print("======================\nSetup Server\n")
        setup_server.run()
        print("Server Setup Finished Successfully\n======================")
        print("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
        print("======================\nSetup Website\n")
        setup_website.run()
        print("Server Website Finished Successfully\n======================")

    print("\n========================================\nSTARTING PROJECT!!!\n")
    time.sleep(1)
    start_server = AppExec(lstTitle[2], lstCmd[2])
    start_application = AppExec(lstTitle[3], lstCmd[3])

    t_server = threading.Thread(target=start_server.run, daemon=True)
    t_website = threading.Thread(target=start_application.run, daemon=True)

    t_server.start()
    t_website.start()

    print("\n============================\nPROGRAM FINISHED\n================================")
