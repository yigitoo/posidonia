from os import system as run
import os
filepath = os.path.join(os.path.dirname(__file__), input('Give filename for creating executable: ') + '.py')
run('pyinstaller --onefile scripts/desktop_app.py -n ')
