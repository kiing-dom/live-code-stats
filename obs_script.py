import obspython as obs
import subprocess
import os

server_process = None
server_exe = "d:\\projects\\live-code-stats\\server.exe"
source_name = "live code stats"


def script_description():
    return "Starts the live-code-stats server when OBS opens and stops it on close."


def script_properties():
    props = obs.obs_properties_create()
    obs.obs_properties_add_path(
        props, "server_exe", "Server Executable",
        obs.OBS_PATH_FILE, "Executables (*.exe)", ""
    )
    obs.obs_properties_add_text(
        props, "source_name", "Browser Source Name",
        obs.OBS_TEXT_DEFAULT
    )
    return props


def script_defaults(settings):
    obs.obs_data_set_default_string(settings, "server_exe", "d:\\projects\\live-code-stats\\server.exe")
    obs.obs_data_set_default_string(settings, "source_name", "live code stats")


def script_update(settings):
    global server_exe, source_name
    server_exe = obs.obs_data_get_string(settings, "server_exe")
    source_name = obs.obs_data_get_string(settings, "source_name")


def script_load(settings):
    script_update(settings)

    if not os.path.exists(server_exe):
        print(f"[live-code-stats] server executable not found: {server_exe}")
        return

    global server_process
    try:
        server_process = subprocess.Popen(
            [server_exe],
            creationflags=subprocess.CREATE_NEW_PROCESS_GROUP | subprocess.CREATE_NO_WINDOW
        )
        print(f"[live-code-stats] server started (pid {server_process.pid})")
    except Exception as e:
        print(f"[live-code-stats] failed to start server: {e}")

    obs.timer_add(refresh_source, 2000)


def refresh_source():
    obs.timer_remove(refresh_source)
    source = obs.obs_get_source_by_name(source_name)
    if not source:
        print(f"[live-code-stats] source '{source_name}' not found")
        return
    proc_handler = obs.obs_source_get_proc_handler(source)
    cd = obs.calldata_create()
    obs.proc_handler_call(proc_handler, "refresh_browser_source", cd)
    obs.calldata_destroy(cd)
    obs.obs_source_release(source)
    print(f"[live-code-stats] browser source '{source_name}' refreshed")


def script_unload():
    global server_process
    obs.timer_remove(refresh_source)
    if server_process:
        try:
            server_process.terminate()
            server_process.wait(timeout=5)
            print("[live-code-stats] server stopped")
        except subprocess.TimeoutExpired:
            server_process.kill()
            print("[live-code-stats] server force killed")
        server_process = None
