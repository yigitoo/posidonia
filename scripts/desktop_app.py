import webview, os

os.environ['APP_URL'] = input("Give the link of app:")
window = webview.create_window(
    'POSIDONIA OCEANICA KORUYUCU',
    f'{os.environ["APP_URL"]}',
    fullscreen=False)

webview.start()
