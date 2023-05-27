import os

def get_path(rel_path: str) -> str:
    return os.path.abspath(os.path.dirname(__file__)) + rel_path

def get_env_data_as_dict(path: str) -> dict:
    with open(path, 'r') as f:
       return dict(tuple(line.replace('\n', '').split('=')) for line
                in f.readlines() if not line.startswith('#'))


go_backend = get_path('/server/.env')
ruby_app = get_path('/posidonia-website/.env')

go_backend_env_content = get_env_data_as_dict(go_backend)
ruby_app_env_content = get_env_data_as_dict(ruby_app)
try:
    new_port_for_go = int(input("Give a port name for syncronize both side\n[this port num will apply for go backend]\nInput: "))
    if new_port_for_go < 1000 or new_port_for_go > 2**16:
        raise SystemExit(f"Give a port num in range of 1000-{str(2**16)}...")
except (ValueError, TypeError, NameError):
    print(f"Please give a integer for port number range in 1000-{str(2**16)}")
    raise SystemExit()
with open(go_backend, 'w+') as f:
    del go_backend_env_content['GO_PORT']
    for key, val in go_backend_env_content.items():
        f.write(f"{key}={val}\n")
    f.write(f'GO_PORT={new_port_for_go}')

with open(ruby_app, 'w+') as f:
    f.write(f'GO_PORT={new_port_for_go}')

