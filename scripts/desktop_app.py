import webview, os

os.environ['APP_URL'] = 'http://localhost:1234/'

link_of_webview = input("Give the URL link of app: ")
if link_of_webview not in ['', None]:
    os.environ['APP_URL'] = link_of_webview

window = webview.create_window(
    'POSIDONIA OCEANICA KORUYUCU',
    f'{os.environ["APP_URL"]}',
    fullscreen=False)

webview.start()
